package cache

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"strconv"
)

type User struct {
}

type UserModel struct {
	Id              int64  `json:"id" redis:"id"`
	Name            string `json:"name" redis:"name"`
	FollowCount     int64  `json:"follow_count,omitempty" redis:"follow_count"`
	FollowerCount   int64  `json:"follower_count,omitempty" redis:"follower_count"`
	Avatar          string `json:"avatar,omitempty" redis:"avatar"`
	BackgroundImage string `json:"background_image,omitempty" redis:"background_image"`
	Signature       string `json:"signature,omitempty" redis:"signature"`
	TotalFavorited  int64  `json:"total_favorited,omitempty" redis:"total_favorited"`
	WorkCount       int64  `json:"work_count,omitempty" redis:"work_count"`
	FavoriteCount   int64  `json:"favorite_count,omitempty" redis:"favorite_count"`
}

// User2UserModel .
func User2UserModel(user *model.User) (userModel *UserModel) {
	userModel = &UserModel{
		Id:   user.Id,
		Name: user.Name,
	}
	if user.FollowCount != nil {
		userModel.FollowCount = *user.FollowCount
	}
	if user.FollowerCount != nil {
		userModel.FollowerCount = *user.FollowerCount
	}
	if user.Avatar != nil {
		userModel.Avatar = *user.Avatar
	}
	if user.BackgroundImage != nil {
		userModel.BackgroundImage = *user.BackgroundImage
	}
	if user.Signature != nil {
		userModel.Signature = *user.Signature
	}
	if user.TotalFavorited != nil {
		userModel.TotalFavorited = *user.TotalFavorited
	}
	if user.WorkCount != nil {
		userModel.WorkCount = *user.WorkCount
	}
	if user.FavoriteCount != nil {
		userModel.FavoriteCount = *user.FavoriteCount
	}
	return
}

// UserModel2User .
func UserModel2User(userModel *UserModel) *model.User {
	var followCount, followerCount, totalFavorited, workCount, favoriteCount = userModel.FollowCount, userModel.FollowerCount, userModel.TotalFavorited, userModel.WorkCount, userModel.FavoriteCount
	var avatar, backgroundImage, signature = userModel.Avatar, userModel.BackgroundImage, userModel.Signature
	return &model.User{
		Id:              userModel.Id,
		Name:            userModel.Name,
		FollowCount:     &followCount,
		FollowerCount:   &followerCount,
		Avatar:          &avatar,
		BackgroundImage: &backgroundImage,
		Signature:       &signature,
		TotalFavorited:  &totalFavorited,
		WorkCount:       &workCount,
		FavoriteCount:   &favoriteCount,
	}
}

func Map2UserModel(mp map[string]string) (userModel *UserModel, err error) {
	var id, followCount, followerCount, totalFavorited, workCount, favoriteCount int64
	id, err = strconv.ParseInt(mp["id"], 10, 64)                              //nolint: staticcheck
	followCount, err = strconv.ParseInt(mp["follow_count"], 10, 64)           //nolint: staticcheck
	followerCount, err = strconv.ParseInt(mp["follower_count"], 10, 64)       //nolint: staticcheck
	totalFavorited, err = strconv.ParseInt(mp["total_favorited	"], 10, 64) //nolint: staticcheck
	workCount, err = strconv.ParseInt(mp["work_count"], 10, 64)               //nolint: staticcheck
	favoriteCount, err = strconv.ParseInt(mp["favorite_count"], 10, 64)
	if err != nil {
		Log.Errorf("parse int from map err: %v", err)
		return nil, err
	}
	return &UserModel{
		Id:              id,
		Name:            mp["name"],
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		Avatar:          mp["avatar"],
		BackgroundImage: mp["background_image"],
		Signature:       mp["signature"],
		TotalFavorited:  totalFavorited,
		WorkCount:       workCount,
		FavoriteCount:   favoriteCount,
	}, nil
}

// SetUserMessage
// 传参为UserModel类型，可以使用User2UserModel进行转换
func (u *User) SetUserMessage(ctx context.Context, user *UserModel) (ok bool) {
	return HSet(ctx, constants.GetUserMsgKey(user.Id), user)
}

// GetUserMessage
// 返回UserModel类型，如果需要model.User类型可以使用UserModel2User进行转换
func (u *User) GetUserMessage(ctx context.Context, userID int64) (user *UserModel, err error) {
	user, err = Map2UserModel(HGetAll(ctx, constants.GetUserMsgKey(userID)))
	if err != nil {
		Log.Warnf("get user %d message err: %v", userID, err)
		return nil, err
	}
	return
}

// GetUserFields
// 1. 根据key和字段名查找值，key或field不存在时，对应的值返回nil，需要调用方自己判断
// 2. 返回的类型都为string，调用方自行转换
func (u *User) GetUserFields(ctx context.Context, userID int64, field ...string) []interface{} {
	return HMGet(ctx, constants.GetUserMsgKey(userID), field...)
}

func (u *User) IncrUserField(ctx context.Context, userID int64, field string, incr int64) (ok bool) {
	return HIncr(ctx, constants.GetUserMsgKey(userID), field, incr)
}

func (u *User) DeleteUser(ctx context.Context, userID int64) (ok bool) {
	return Delete(ctx, constants.GetUserMsgKey(userID))
}
