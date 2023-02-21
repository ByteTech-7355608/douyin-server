package base

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"strconv"

	. "ByteTech-7355608/douyin-server/pkg/configs"
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
		if s.cache.Video.IsExists(ctx, v.Id) != 0 {
			needs := []string{"comment_count", "favorite_count"}
			res := s.cache.Video.GetVideoFields(ctx, v.Id, needs...)
			v.CommentCount, _ = strconv.ParseInt(res[0].(string), 10, 0)
			v.FavoriteCount, _ = strconv.ParseInt(res[1].(string), 10, 0)
		} else {
			v.CommentCount = video.CommentCount
			v.FavoriteCount = video.FavoriteCount
			v_model := cache.Video2VideoModel(v)
			if ok := s.cache.Video.SetVideoMessage(ctx, v_model); !ok {
				Log.Errorf("set video redis err: %v", err)
				return
			}
		}
		v.IsFavorite = false
		userID := req.GetBaseReq().GetUserId()
		if userID != 0 {
			if s.cache.Like.IsExists(ctx, userID) != 0 {
				v.IsFavorite = s.cache.Like.IsLike(ctx, userID, v.Id)
			} else {
				v.IsFavorite, err = s.dao.Like.IsLike(ctx, userID, video.ID)
				if err != nil {
					Log.Errorf("query relation between user %v and video %v err: %v", userID, video.UID, err)
					return
				}
				var action int64
				if v.IsFavorite {
					action = 1
				} else {
					action = -1
				}
				s.cache.Like.FavoriteAction(ctx, userID, video.ID, action)
			}

		}
		if s.cache.User.IsExists(ctx, video.UID) != 0 {
			author_model, err := s.cache.User.GetUserMessage(ctx, video.UID)
			if err != nil {
				Log.Errorf("Get usermessage from redis err: %v", err)
			}
			author := cache.UserModel2User(author_model)
			v.Author = author
		} else {
			author2, err := s.dao.User.QueryUser(ctx, video.UID)
			if err != nil {
				Log.Warnf("query user %v err: %v", video.UID, err)
				continue
			}
			v.Author = &model.User{
				Id:            author2.ID,
				Name:          author2.Username,
				FollowCount:   &author2.FollowCount,
				FollowerCount: &author2.FollowerCount,
				Avatar:        &author2.Avatar,
				IsFollow:      false,
			}
			author_model := cache.User2UserModel(v.Author)
			if ok := s.cache.User.SetUserMessage(ctx, author_model); !ok {
				Log.Errorf("set video redis err: %v", err)
			}
		}
		if userID != 0 {
			if s.cache.Relation.IsExists(ctx, userID) != 0 {
				v.Author.IsFollow = s.cache.Relation.IsFollow(ctx, userID, video.UID)
			} else {
				v.Author.IsFollow, _ = s.dao.Relation.IsFollower(ctx, userID, video.UID)
				var action int64
				if v.Author.IsFollow {
					action = 1
				} else {
					action = -1
				}
				s.cache.Relation.FollowAction(ctx, userID, video.ID, action)
			}
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
