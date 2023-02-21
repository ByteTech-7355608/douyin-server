package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"time"
)

func (s *Service) Feed(ctx context.Context, req *base.DouyinFeedRequest) (resp *base.DouyinFeedResponse, err error) {
	resp = base.NewDouyinFeedResponse()
	latestTime := req.GetLatestTime()
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	videos, err := s.dao.Video.QueryVideoByTime(ctx, latestTime)
	if err != nil {
		err = constants.ErrQueryRecord
		Log.Errorf("query video err: %v", err)
		return
	}
	videoList := make([]*model.Video, len(videos))
	for i, v := range videoList {
		v = &model.Video{}
		video := videos[i]
		v.Id = video.ID
		v.Title = video.Title
		v.PlayUrl = video.PlayURL
		v.CoverUrl = video.CoverURL
		v.CommentCount = video.CommentCount
		v.FavoriteCount = video.FavoriteCount
		v.IsFavorite = false
		userID := req.GetBaseReq().GetUserId()
		if userID != 0 {
			v.IsFavorite, err = s.dao.Like.IsLike(ctx, userID, video.ID)
			if err != nil {
				err = constants.ErrQueryRecord
				Log.Warnf("query relation between user %v and video %v err: %v", userID, video.UID, err)
			}
		}
		author, err := s.dao.User.FindUserById(ctx, video.UID)
		if err != nil {
			Log.Warnf("query user %v err: %v", video.UID, err)
			continue
		}
		v.Author = &model.User{
			Id:            author.ID,
			Name:          author.Username,
			FollowCount:   &author.FollowCount,
			FollowerCount: &author.FollowerCount,
			Avatar:        &author.Avatar,
			IsFollow:      false,
		}
		if userID != 0 {
			v.Author.IsFollow, _ = s.dao.Relation.IsFollower(ctx, userID, video.UID)
		}
		videoList[i] = v
	}
	resp.VideoList = videoList
	if len(videos) > 0 {
		nextTime := videos[len(videos)-1].CreatedAt.Unix()
		resp.NextTime = &nextTime
	}
	return
}
