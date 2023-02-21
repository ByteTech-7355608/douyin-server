package cache

import (
	dbmodel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"fmt"
	"strconv"
)

type Like struct {
}

func (l *Like) IsExists(ctx context.Context, uids ...int64) int64 {
	keys := make([]string, len(uids))
	for i, uid := range uids {
		keys[i] = constants.GetUserLikeListKey(uid)
	}
	return Exists(ctx, keys...)
}

func (l *Like) IsLike(ctx context.Context, uid, vid int64) bool {
	res := HMGet(ctx, constants.GetUserLikeListKey(uid), strconv.FormatInt(vid, 10))

	if res[0] == nil || res[0].(string) == "0" {
		return false
	}
	return true
}

func (l *Like) SetFavoriteList(ctx context.Context, userID int64, kv ...string) bool {
	return HSet(ctx, constants.GetUserLikeListKey(userID), kv)
}

func (l *Like) FavoriteAction(ctx context.Context, uid, vid int64, action int64) bool {
	return HIncr(ctx, constants.GetUserLikeListKey(uid), strconv.FormatInt(vid, 10), action)
}

func (l *Like) SetLikeMessage(ctx context.Context, uid, vid int64, action bool) (ok bool) {
	data := make(map[string]interface{})
	data[fmt.Sprintf("%d", vid)] = action
	return HSet(ctx, constants.GetUserLikeListKey(uid), data)
}

func (l *Like) GetLikeField(ctx context.Context, uid int64, field ...string) []interface{} {
	return HMGet(ctx, constants.GetUserLikeListKey(uid), field...)
}

// GetAllUserLikes 获取当前用户的所有喜欢的videos vid;
func (l *Like) GetAllUserLikes(ctx context.Context, uid int64) (userLikes []dbmodel.Like) {
	userLikes = make([]dbmodel.Like, 0)
	res := HGetAll(ctx, constants.GetUserLikeListKey(uid))
	for k, v := range res {
		vid, _ := strconv.ParseInt(k, 10, 64)
		action, _ := strconv.ParseInt(v, 10, 64)
		if action == 1 {
			userLikes = append(userLikes, dbmodel.Like{
				Vid: vid,
			})
		}
	}
	return
}
