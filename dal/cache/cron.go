package cache

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"strconv"
	"time"
)

func SyncDataToDB() {
	r := NewRedisCache()
	ctx := context.Background()
	go r.UpdateUserMsgToDB(ctx)
	go r.UpdateUserLikeListToDB(ctx)
	go r.UpdateUserFollowListToDB(ctx)
	go r.UpdateVideoMsgToDB(ctx)
}

func (r *RedisCache) UpdateUserMsgToDB(ctx context.Context) {
	begin := time.Now().Unix()
	Log.Info("start sync user message to db")
	var cursor uint64
	keys, _, err := cli.Scan(ctx, cursor, "user_message_*", 0).Result()
	if err != nil {
		Log.Errorf("scan user message keys err: %v", err)
		return
	}
	for _, key := range keys {
		uid := constants.GetIDFromUserMsgKey(key)
		user, err := r.User.GetUserMessage(ctx, uid)
		if err != nil {
			Log.Errorf("get user %v message from redis err: %v", uid, err)
			continue
		}
		userMap := map[string]interface{}{
			"id":               user.Id,
			"username":         user.Name,
			"follow_count":     user.FollowCount,
			"follower_count":   user.FollowerCount,
			"total_favorited":  user.TotalFavorited,
			"work_count":       user.WorkCount,
			"favorite_count":   user.FavoriteCount,
			"avatar":           user.Avatar,
			"signature":        user.Signature,
			"background_image": user.BackgroundImage,
		}
		err = r.dao.User.UpdateUser(ctx, uid, &userMap)
		if err != nil {
			Log.Errorf("update user %v to db err: %v", uid, err)
			continue
		}
	}
	Log.Infof("sync user message to db end, takes %v s", time.Now().Unix()-begin)
	return
}

func (r *RedisCache) UpdateUserLikeListToDB(ctx context.Context) {
	begin := time.Now().Unix()
	Log.Info("start sync user like list to db")
	var cursor uint64
	keys, _, err := cli.Scan(ctx, cursor, "user_like_list_*", 0).Result()
	if err != nil {
		Log.Errorf("scan user like list keys err: %v", err)
		return
	}
	for _, key := range keys {
		uid := constants.GetIDFromUserLikeListKey(key)
		list := r.Like.GetFavoriteList(ctx, uid)
		for k, v := range list {
			vid, err := strconv.ParseInt(k, 10, 64) //nolint: staticcheck
			action, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				Log.Errorf("parse int err: %v", err)
				continue
			}
			err = r.dao.Like.UpsertRecord(ctx, &model.Like{
				UID:    uid,
				Vid:    vid,
				Action: action == 1,
			})
			if err != nil {
				Log.Errorf("upsert like %v:%v:%v err: %v", uid, vid, action, err)
				continue
			}
		}
	}
	Log.Infof("sync user like list to db end, takes %v s", time.Now().Unix()-begin)
	return
}

func (r *RedisCache) UpdateUserFollowListToDB(ctx context.Context) {
	begin := time.Now().Unix()
	Log.Info("start sync user follow list to db")
	var cursor uint64
	keys, _, err := cli.Scan(ctx, cursor, "user_follow_list_*", 0).Result()
	if err != nil {
		Log.Errorf("scan user follow list keys err: %v", err)
		return
	}
	for _, key := range keys {
		uid := constants.GetIDFromUserFollowListKey(key)
		list := r.Relation.GetFollowFullList(ctx, uid)
		for k, v := range list {
			uid2, err := strconv.ParseInt(k, 10, 64) //nolint: staticcheck
			action, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				Log.Errorf("parse int err: %v", err)
				continue
			}
			err = r.dao.Relation.UpsertRecord(ctx, &model.Relation{
				ConcernerID: uid,
				ConcernedID: uid2,
				Action:      action == 1,
			})
			if err != nil {
				Log.Errorf("upsert relation %v:%v:%v err: %v", uid, uid2, action, err)
				continue
			}
		}
	}
	Log.Infof("sync user follow list to db end, takes %v s", time.Now().Unix()-begin)
	return
}

func (r *RedisCache) UpdateVideoMsgToDB(ctx context.Context) {
	begin := time.Now().Unix()
	Log.Info("start sync video message to db")
	var cursor uint64
	keys, _, err := cli.Scan(ctx, cursor, "video_message_*", 0).Result()
	if err != nil {
		Log.Errorf("scan video message keys err: %v", err)
		return
	}
	for _, key := range keys {
		vid := constants.GetIDFromVideoMsgKey(key)
		video, err := r.Video.GetVideoMessage(ctx, vid)
		if err != nil {
			Log.Errorf("get video %v message from redis err: %v", vid, err)
			continue
		}
		videoMap := map[string]interface{}{
			"id":             video.Id,
			"play_url":       video.PlayUrl,
			"cover_url":      video.CoverUrl,
			"favorite_count": video.FavoriteCount,
			"comment_count":  video.CommentCount,
			"title":          video.Title,
			"uid":            video.AuthorID,
		}
		err = r.dao.Video.UpdateVideo(ctx, vid, &videoMap)
		if err != nil {
			Log.Errorf("update video %v to db err: %v", vid, err)
			continue
		}
	}
	Log.Infof("sync video message to db end, takes %v s", time.Now().Unix()-begin)
	return
}
