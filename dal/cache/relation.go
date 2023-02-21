package cache

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"strconv"
)

type Relation struct {
	dao *dao.Dao
}

func (r *Relation) FollowIsExists(ctx context.Context, uids ...int64) int64 {
	keys := make([]string, len(uids))
	for i, uid := range uids {
		keys[i] = constants.GetUserFollowListKey(uid)
	}
	return Exists(ctx, keys...)
}

func (r *Relation) FollowerIsExists(ctx context.Context, uids ...int64) int64 {
	keys := make([]string, len(uids))
	for i, uid := range uids {
		keys[i] = constants.GetUserFollowerListKey(uid)
	}
	return Exists(ctx, keys...)
}

func (r *Relation) IsFollow(ctx context.Context, from_id, to_id int64) bool {
	res := HMGet(ctx, constants.GetUserFollowListKey(from_id), strconv.FormatInt(to_id, 10))
	if res[0] == nil || res[0].(string) == "0" {
		return false
	}
	return true
}

func (r *Relation) SetFollowList(ctx context.Context, userID int64, kv ...string) bool {
	return HSet(ctx, constants.GetUserFollowListKey(userID), kv)
}

func (r *Relation) SetFollowerList(ctx context.Context, userID int64, kv ...string) bool {
	return HSet(ctx, constants.GetUserFollowerListKey(userID), kv)
}

func (r *Relation) FollowAction(ctx context.Context, from_id, to_id int64, action int64) bool {
	b1 := HIncr(ctx, constants.GetUserFollowListKey(from_id), strconv.FormatInt(to_id, 10), action)
	b2 := HIncr(ctx, constants.GetUserFollowerListKey(to_id), strconv.FormatInt(from_id, 10), action)
	return b1 && b2
}

func (r *Relation) GetFollowList(ctx context.Context, userID int64) (followList []int64) {
	followList = make([]int64, 0)
<<<<<<< HEAD
	res := HGetAll(ctx, constants.GetUserFollowerListKey(userID))
=======
	res := HGetAll(ctx, constants.GetUserFollowListKey(userID))
>>>>>>> origin/syx-dev-redis
	for k, v := range res {
		uid, _ := strconv.ParseInt(k, 10, 64)
		action, _ := strconv.ParseInt(v, 10, 64)
		if action == 1 {
			followList = append(followList, uid)
		}
	}
	return
}

func (r *Relation) GetFollowerList(ctx context.Context, userID int64) (followerList []int64) {
	followerList = make([]int64, 0)
	res := HGetAll(ctx, constants.GetUserFollowerListKey(userID))
	for k, v := range res {
		uid, _ := strconv.ParseInt(k, 10, 64)
		action, _ := strconv.ParseInt(v, 10, 64)
		if action == 1 {
			followerList = append(followerList, uid)
		}
	}
	return
}
<<<<<<< HEAD
=======

func (r *Relation) GetFollowFullList(ctx context.Context, userID int64) map[string]string {
	return HGetAll(ctx, constants.GetUserFollowListKey(userID))
}
>>>>>>> origin/syx-dev-redis
