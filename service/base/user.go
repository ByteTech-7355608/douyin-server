package base

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/jwt"
	"context"
)

func (s *Service) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest) (resp *base.DouyinUserRegisterResponse, err error) {
	resp = base.NewDouyinUserRegisterResponse()

	id, err := s.dao.User.AddUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		Log.Errorf("add user err: %v", err)
		return
	}

	resp.UserId = id
	resp.Token, err = jwt.GenToken(id, req.GetUsername())
	return
}

func (s *Service) UserLogin(ctx context.Context, req *base.DouyinUserLoginRequest) (resp *base.DouyinUserLoginResponse, err error) {
	resp = base.NewDouyinUserLoginResponse()

	id, err := s.dao.User.CheckUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		Log.Errorf("user login err: %v", err)
		return
	}

	resp.UserId = id
	resp.Token, err = jwt.GenToken(id, req.GetUsername())
	return
}

func (s *Service) UserMsg(ctx context.Context, req *base.DouyinUserRequest) (resp *base.DouyinUserResponse, err error) {
	resp = base.NewDouyinUserResponse()
	var user model2.User
	user.Id = req.UserId //被查看的用户id
	if s.cache.User.IsExists(ctx, user.Id) != 0 {
		usermsg, err := s.cache.User.GetUserMessage(ctx, user.Id)
		if err != nil {
			Log.Errorf("Get usermsg from redis err: %v", err)
			return resp, err
		}
		user = *cache.UserModel2User(usermsg)
	} else {
		userinfo, err := s.dao.User.FindUserById(ctx, user.Id)
		if err != nil {
			Log.Errorf("Get usermsg from db err: %v", err)
			return resp, err
		}
		user.Name = userinfo.Username
		user.FollowCount = &userinfo.FollowCount
		user.FollowerCount = &userinfo.FollowerCount
		user.FavoriteCount = &userinfo.FavoriteCount
		user.TotalFavorited = &userinfo.TotalFavorited
		user.WorkCount = &userinfo.WorkCount
		user.Avatar = &userinfo.Avatar
		user.BackgroundImage = &userinfo.BackgroundImage
		user.Signature = &userinfo.Signature
		cache_user := cache.User2UserModel(&user)
		s.cache.User.SetUserMessage(ctx, cache_user)
	}
	if s.cache.Relation.IsExists(ctx, *req.BaseReq.UserId) != 0 {
		user.IsFollow = s.cache.Relation.IsFollow(ctx, *req.BaseReq.UserId, user.Id)
	} else {
		user.IsFollow, err = s.dao.Relation.IsUserFollowed(ctx, *req.BaseReq.UserId, user.Id)
		if err != nil {
			Log.Errorf("Get user relation from db err: %v", err)
			return
		}
		var action int64
		if user.IsFollow {
			action = 1
		} else {
			action = 0
		}
		s.cache.Relation.FollowAction(ctx, *req.BaseReq.UserId, user.Id, action)
	}
	resp.SetUser(&user)
	return
}
