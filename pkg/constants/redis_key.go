package constants

import (
	"fmt"
	"strconv"
)

func GetUserMsgKey(userID int64) string {
	return fmt.Sprintf("user_message_%d", userID)
}

func GetUserLikeListKey(userID int64) string {
	return fmt.Sprintf("user_like_list_%d", userID)
}

func GetUserFollowListKey(userID int64) string {
	return fmt.Sprintf("user_follow_list_%d", userID)
}

func GetUserFollowerListKey(userID int64) string {
	return fmt.Sprintf("user_follower_list_%d", userID)
}

func GetVideoMsgKey(videoID int64) string {
	return fmt.Sprintf("video_message_%d", videoID)
}

func GetIDFromUserMsgKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[13:], 10, 64)
	return
}

func GetIDFromUserLikeListKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[15:], 10, 64)
	return
}

func GetIDFromUserFollowListKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[17:], 10, 64)
	return
}

func GetIDFromUserFollowerListKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[19:], 10, 64)
	return
}

func GetIDFromVideoMsgKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[14:], 10, 64)
	return
}

func GetFavoriteLmtKey(ipaddr string) string {
	return "favorite_limit_" + ipaddr
}
