package cache

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
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
