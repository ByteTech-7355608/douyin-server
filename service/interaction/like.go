package interaction

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	dbmodel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func (s *Service) FavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	resp = interaction.NewDouyinFavoriteListResponse()
	var uid = req.GetBaseReq().GetUserId()

	var userLikes []dbmodel.Like
	// 是否已经读取数据
	hasSearch := false
	// 根据 uid 从 like 表中查找喜欢的视频列表 vid list
	if s.cache.Like.IsExists(ctx, uid) == 0 {
		// 根据uid查询喜欢视频列表的视频vid列表未命中缓存,查询数据库
		// TODO: 小心缓存穿透
		userLikes, err = s.dao.Like.QueryUserLikeRecords(ctx, uid)
		hasSearch = true
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return resp, nil
			}
			return resp, constants.ErrQueryRecord
		}
		// 将查询到数据的加入缓存
		if len(userLikes) > 0 {
			kv := make([]string, 0)
			for _, userLike := range userLikes {
				kv = append(kv, strconv.FormatInt(userLike.Vid, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Like.SetFavoriteList(ctx, uid, kv...) {
				Log.Errorf("set favorite like to redis err: %v", err)
				return resp, constants.ErrWriteCache
			}
		}
	}
	// 查询缓存获得用户喜欢视频的vid列表
	if !hasSearch {
		userLikes = s.cache.Like.GetAllUserLikes(ctx, uid)
	}

	//  根据 vid 列表查询 videoList
	videoList := make([]*dbmodel.Video, 0)
	for _, userLike := range userLikes {
		hasSearch = false
		var video *dbmodel.Video
		if s.cache.Video.IsExists(ctx, userLike.Vid) == 0 {
			// 根据vid查询视频未命中缓存,查询数据库
			video, err = s.dao.Video.QueryRecord(ctx, userLike.Vid)
			hasSearch = true
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}
				return resp, constants.ErrQueryRecord
			}
			// TODO: Video2VideoModel Done
			// dao/model/video -> cache.VideoModel
			videoModel := cache.DBVideo2VideoModel(video)
			if !s.cache.Video.SetVideoMessage(ctx, videoModel) {
				Log.Errorf("set video message to redis err: %v", err)
				return resp, constants.ErrWriteCache
			}
		}
		if !hasSearch {
			videoModel, err := s.cache.Video.GetVideoMessage(ctx, userLike.Vid)
			if err != nil {
				return resp, constants.ErrReadCache
			}
			// TODO: VideoModel2DBVideo Done
			// cache.VideoModel -> dao/model/video
			video = cache.VideoModel2DBVideo(videoModel)
		}
		videoList = append(videoList, video)
	}

	// 将video中的视频根据uid找到User拼接得到 kitex_gan/model/video
	var videos []*model.Video
	for _, videoInstance := range videoList {
		hasSearch = false
		var userInstance dbmodel.User
		if s.cache.User.IsExists(ctx, videoInstance.UID) == 0 {
			userInstance, err = s.dao.User.FindUserById(ctx, videoInstance.UID)
			hasSearch = true
			// TODO 小心缓存穿透
			if err != nil {
				if err == constants.ErrUserNotExist {
					// TODO 用户上传视频后被软删除导致不存在 默认用户？
					Log.Warnf("get user uid: %d err: %v", videoInstance.UID, err)
				} else {
					return resp, err
				}
			}
			// TODO DBUser2UserModel Done
			if !s.cache.User.SetUserMessage(ctx, cache.DBUser2UserModel(&userInstance)) {
				Log.Errorf("set user message to redis err: %v", err)
				return resp, constants.ErrWriteCache
			}
		}
		if !hasSearch {
			userModel, err := s.cache.User.GetUserMessage(ctx, videoInstance.UID)
			if err != nil {
				return resp, constants.ErrReadCache
			}
			// TODO UserModel2DBUser Done
			userInstance = *cache.UserModel2DBUser(userModel)
		}

		var isFollow bool
		hasSearch = false
		if s.cache.Relation.FollowIsExists(ctx, uid) == 0 {
			isFollow, err = s.dao.Relation.IsUserFollowed(ctx, uid, videoInstance.UID)
			if err != nil {
				// 查找关注关系时数据库出错，跳过该视频，不影响输出结果
				Log.Warnf("get follow err: %v", err)
				isFollow = false
			}
			hasSearch = true
			// 这里应该将所有的uid对应等于1的加入缓存
			followList, err := s.dao.Relation.FollowidList(ctx, uid)
			if err != nil {
				Log.Errorf("get followlist err: %v", err)
			}
			// 将查询到数据的加入缓存
			if len(followList) > 0 {
				kv := make([]string, 0)
				for _, follow := range followList {
					kv = append(kv, strconv.FormatInt(follow, 10))
					kv = append(kv, "1")
				}
				if !s.cache.Relation.SetFollowList(ctx, uid, kv...) {
					Log.Errorf("set followlist to redis err!")
					return resp, constants.ErrWriteCache
				}
			}

		}
		if !hasSearch {
			isFollow = s.cache.Relation.IsFollow(ctx, uid, userInstance.ID)
		}
		user := &model.User{
			Id:            userInstance.ID,
			Name:          userInstance.Username,
			FollowCount:   &userInstance.FollowCount,
			FollowerCount: &userInstance.FollowerCount,
			IsFollow:      isFollow,
		}
		video := &model.Video{
			PlayUrl:       videoInstance.PlayURL,
			CoverUrl:      videoInstance.CoverURL,
			FavoriteCount: videoInstance.FavoriteCount,
			CommentCount:  videoInstance.CommentCount,
			IsFavorite:    true,
			Title:         videoInstance.Title,
			Author:        user,
		}
		videos = append(videos, video)
	}
	resp.SetVideoList(videos)
	return resp, nil
}

func (s *Service) FavoriteAction(ctx context.Context, req *interaction.DouyinFavoriteActionRequest) (resp *interaction.DouyinFavoriteActionResponse, err error) {
	resp = interaction.NewDouyinFavoriteActionResponse()
	var uid, vid = req.GetBaseReq().GetUserId(), req.GetVideoId()

	if s.cache.Like.IsExists(ctx, uid) == 0 {
		// 当前用户点赞列表缓存不存在
		videoList, err := s.dao.Like.GetFavoriteVideoListByUserId(ctx, uid)
		if err != nil {
			Log.Errorf("get favorite video list err: %v, uid: %v", err, uid)
		}
		if len(videoList) > 0 {
			kv := make([]string, 0)
			for _, video := range videoList {
				kv = append(kv, strconv.FormatInt(video.ID, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Like.SetFavoriteList(ctx, uid, kv...) {
				Log.Errorf("set favorite list to redis err")
				return resp, constants.ErrWriteCache
			}
		}
	}
	if s.cache.User.IsExists(ctx, uid) == 0 {
		// 当前用户缓存不存在
		user, err := s.dao.User.QueryUser(ctx, uid)
		if err != nil {
			Log.Errorf("query user %v err: %v", uid, err)
			return resp, err
		}
		userModel := &cache.UserModel{
			Id:              user.ID,
			Name:            user.Username,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		if !s.cache.User.SetUserMessage(ctx, userModel) {
			Log.Errorf("set user message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	}
	if s.cache.Video.IsExists(ctx, vid) == 0 {
		// 点赞视频缓存不存在
		video, err := s.dao.Video.QueryVideoByID(ctx, vid)
		if err != nil {
			Log.Errorf("query video %v err: %v", vid, err)
			return resp, err
		}
		videoModel := &cache.VideoModel{
			Id:            video.ID,
			AuthorID:      video.UID,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
		}
		if !s.cache.Video.SetVideoMessage(ctx, videoModel) {
			Log.Errorf("set video message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	}
	res := s.cache.Video.GetVideoFields(ctx, vid, "author_id")
	authorID, _ := strconv.ParseInt(res[0].(string), 10, 64)
	if s.cache.User.IsExists(ctx, authorID) == 0 {
		// 视频作者缓存不存在
		user, err := s.dao.User.QueryUser(ctx, authorID)
		if err != nil {
			Log.Errorf("query user %v err: %v", authorID, err)
			return resp, err
		}
		userModel := &cache.UserModel{
			Id:              user.ID,
			Name:            user.Username,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		if !s.cache.User.SetUserMessage(ctx, userModel) {
			Log.Errorf("set user message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	}

	var action int64
	like := s.cache.Like.IsLike(ctx, uid, vid)
	if like && req.GetActionType() == 2 {
		action = -1
	} else if !like && req.GetActionType() == 1 {
		action = 1
	} else {
		return
	}

	// TODO 原子操作
	// 更新点赞列表
	if !s.cache.Like.FavoriteAction(ctx, uid, vid, action) {
		return resp, constants.ErrWriteCache
	}
	// 更新用户点赞数
	if !s.cache.User.IncrUserField(ctx, uid, "favorite_count", action) {
		return resp, constants.ErrWriteCache
	}
	// 更新作者获赞数
	if !s.cache.User.IncrUserField(ctx, authorID, "total_favorited", action) {
		return resp, constants.ErrWriteCache
	}
	// 更新视频获赞数
	if !s.cache.Video.IncrVideoField(ctx, vid, "favorite_count", action) {
		return resp, constants.ErrWriteCache
	}

	return
}
