package cache

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"strconv"
)

type Relation struct {
}

func (r *Relation) IsExists(ctx context.Context, uids ...int64) int64 {
	keys := make([]string, len(uids))
	for i, uid := range uids {
		keys[i] = constants.GetUserFollowListKey(uid)
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
func (r *Relation) FollowAction(ctx context.Context, from_id, to_id int64, action int64) bool {
	return HIncr(ctx, constants.GetUserFollowListKey(from_id), strconv.FormatInt(to_id, 10), action)
}
