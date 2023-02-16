package interaction

import (
	dbmodel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"errors"

	"gorm.io/gorm"
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
	//var videos []*model.Video
	videos := make([]*model.Video, len(videoList))
	for _, videoInstance := range videoList {
		userInstance, err := s.dao.User.FindUserById(ctx, videoInstance.UID)
		if err != nil {
			// 某一个视频没有找到作者，跳过该视频，不影响输出结果
			Log.Warnf("get user err: %v", err)
			continue
		}
		isFollow, err := s.dao.Relation.IsUserFollowed(ctx, uid, videoInstance.UID)
		if err != nil {
			// 查找关注关系时数据库出错，跳过该视频，不影响输出结果
			Log.Warnf("get follow err: %v", err)
			continue
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
	var record *dbmodel.Like
	var uid, vid = req.GetBaseReq().GetUserId(), req.GetVideoId()
	record, err = s.dao.Like.QueryRecord(ctx, uid, vid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.dao.Like.CreateRecord(ctx, &dbmodel.Like{UID: uid, Vid: vid, Action: req.GetActionType() == 1})
		if err != nil {
			Log.Errorf("create like record err: %v, uid: %v, vid: %v", err, uid, vid)
			return resp, constants.ErrCreateRecord
		}
		return resp, nil
	}
	if err != nil {
		Log.Errorf("query like record err: %v, uid: %v, vid: %v", err, uid, vid)
		return resp, constants.ErrQueryRecord
	}
	if (record.Action && req.GetActionType() == 1) || (!record.Action && req.GetActionType() == 2) {
		return resp, nil
	}
	record.Action = req.GetActionType() == 1
	err = s.dao.Like.UpdateRecord(ctx, record)
	if err != nil {
		Log.Errorf("update record err: %v, uid: %v, vid: %v", err, uid, vid)
		return resp, constants.ErrUpdateRecord
	}

	return
}
