package interaction

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

func (s *Service) FavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	resp = interaction.NewDouyinFavoriteListResponse()
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
			Log.Infof("get user err: %v", err)
			continue
		}
		isFollow, err := s.dao.Relation.IsUserFollowed(ctx, req.GetUserId(), videoInstance.UID)
		if err != nil {
			// 查找关注关系时数据库出错，跳过该视频，不影响输出结果
			Log.Infof("get follow err: %v", err)
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
