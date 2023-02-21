package base

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	dbmodel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/pkg/rabbitmq"
	"context"
	"strconv"
)

func (s *Service) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (resp *base.DouyinPublishListResponse, err error) {
	resp = base.NewDouyinPublishListResponse()
	var upUid, uid = req.GetUserId(), req.GetBaseReq().GetUserId()

	// 查找用户
	var userInstance dbmodel.User
	if s.cache.User.IsExists(ctx, req.GetUserId()) == 0 {
		userInstance, err = s.dao.User.FindUserById(ctx, upUid)
		if err != nil {
			if err == constants.ErrUserNotExist {
				// TODO 用户上传视频后被软删除导致不存在 默认用户？
				Log.Warnf("get user uid: %d err: %v", upUid, err)
			} else {
				return resp, err
			}
		}
		// TODO DBUser2UserModel Done
		if !s.cache.User.SetUserMessage(ctx, cache.DBUser2UserModel(&userInstance)) {
			Log.Errorf("set user message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	} else {
		userModel, err := s.cache.User.GetUserMessage(ctx, upUid)
		if err != nil {
			return resp, constants.ErrReadCache
		}
		// TODO UserModel2DBUser Done
		userInstance = *cache.UserModel2DBUser(userModel)
	}

	// 查找用户视频
	// TODO: redis 缓存优化 user_publish_list_{uid}
	videoList, err := s.dao.Video.GetPublishVideoListByUserId(ctx, upUid)
	if err != nil {
		Log.Errorf("get publish list err : %v", err)
		return
	}
	// 是否关注
	var isFollow bool
	if s.cache.Relation.FollowIsExists(ctx, uid, upUid) == 0 {
		// TODO: 可以先把所有Follow都加入缓存，再差是否follow，这样读缓存比读内存快
		isFollow, err = s.dao.Relation.IsUserFollowed(ctx, uid, upUid)
		if err != nil {
			// 查找关注关系时数据库出错，跳过该视频，不影响输出结果
			Log.Warnf("get follow err: %v", err)
			isFollow = false
		}
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
				Log.Errorf("set followlist to redis err: %v", err)
				return resp, constants.ErrWriteCache
			}
		}
	} else {
		isFollow = s.cache.Relation.IsFollow(ctx, uid, upUid)
	}

	//user 类型转换
	user := &model.User{
		Id:            userInstance.ID,
		Name:          userInstance.Username,
		FollowCount:   &userInstance.FollowCount,
		FollowerCount: &userInstance.FollowerCount,
		IsFollow:      isFollow,
	}

	//video 类型转换
	var videos []*model.Video
	for _, videoInstance := range videoList {
		var isLike bool

		if s.cache.Like.IsExists(ctx, uid) == 0 {
			isLike, err = s.dao.Like.IsLike(ctx, uid, videoInstance.ID)
			if err != nil {
				Log.Infof("Query IsLike failed: %v", err)
				isLike = false
			}
			userLikes, err := s.dao.Like.QueryUserLikeRecords(ctx, uid)
			if err != nil {
				// TODO: 小心缓存击穿
				Log.Infof("Query QueryUserLikeRecords failed: %v", err)
				isLike = false
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
		} else {
			isLike = s.cache.Like.IsLike(ctx, uid, videoInstance.ID)
		}

		video := &model.Video{
			Id:            videoInstance.ID,
			PlayUrl:       videoInstance.PlayURL,
			CoverUrl:      videoInstance.CoverURL,
			FavoriteCount: videoInstance.FavoriteCount,
			CommentCount:  videoInstance.CommentCount,
			Title:         videoInstance.Title,
			IsFavorite:    isLike,
			Author:        user,
		}
		videos = append(videos, video)
	}
	resp.SetVideoList(videos)
	return
}

func (s *Service) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (r *base.DouyinPublishActionResponse, err error) {
	r = base.NewDouyinPublishActionResponse()
	user_id := *req.BaseReq.UserId
	err = rabbitmq.Produce(user_id, req.Title, *req.PlayUrl)
	go rabbitmq.Consume(ctx)
	if err != nil {
		Log.Errorf("publish action err : %v", err)
		return
	}
	return
}
