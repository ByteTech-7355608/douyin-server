package interaction

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"strconv"
)

func (s *Service) FavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	resp = interaction.NewDouyinFavoriteListResponse()
	var uid = req.GetBaseReq().GetUserId()
	// 根据 uid 从 like 表中查找喜欢的视频列表 vid list 然后根据 vid 查询 videoList
	videoList, err := s.dao.Like.GetFavoriteVideoListByUserId(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get favorite video list err: %v", err)
		return
	}

	var videos []*model.Video
	for _, videoInstance := range videoList {
		userInstance, err := s.dao.User.FindUserById(ctx, videoInstance.UID)
		if err != nil {
			// 某一个视频没有找到作者，跳过该视频，不影响输出结果
			Log.Warnf("get user err: %v", err)
		}
		isFollow, err := s.dao.Relation.IsUserFollowed(ctx, uid, videoInstance.UID)
		if err != nil {
			// 查找关注关系时数据库出错，跳过该视频，不影响输出结果
			Log.Warnf("get follow err: %v", err)
			isFollow = false
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
		likeMap := make(map[string]int64, 0)
		for _, video := range videoList {
			likeMap[strconv.FormatInt(video.UID, 10)] = 1
		}
		if !s.cache.Like.SetFavoriteList(ctx, uid, likeMap) {
			Log.Errorf("set favorite like to redis err: %v", err)
			return resp, constants.ErrWriteCache
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
	var authorID int64
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
		authorID = video.UID
	}
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
