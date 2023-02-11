package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

func (s *Service) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (resp *base.DouyinPublishListResponse, err error) {
	resp = base.NewDouyinPublishListResponse()

	userInstance, err := s.dao.User.FindUserById(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get user err: %v", err)
		return
	}

	videoList, err := s.dao.Video.GetPublishVideoListByUserId(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get publish list err : %v", err)
		return
	}

	//user 类型转换
	user := &model.User{
		Id:            userInstance.ID,
		Name:          userInstance.Username,
		FollowCount:   &userInstance.FollowCount,
		FollowerCount: &userInstance.FollowerCount,
	}

	//video 类型转换
	var videos []*model.Video
	for _, videoInstance := range videoList {
		video := &model.Video{
			Id:            videoInstance.ID,
			PlayUrl:       videoInstance.PlayURL,
			CoverUrl:      videoInstance.CoverURL,
			FavoriteCount: videoInstance.FavoriteCount,
			CommentCount:  videoInstance.CommentCount,
			Title:         videoInstance.Title,
			Author:        user,
		}
		videos = append(videos, video)
	}
	resp.SetVideoList(videos)
	return
}
