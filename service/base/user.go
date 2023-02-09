package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/jwt"
	"context"
)

func (s *Service) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest) (resp *base.DouyinUserRegisterResponse, err error) {
	resp = base.NewDouyinUserRegisterResponse()

	id, err := s.dao.User.AddUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		Log.Errorf("add user err: %v", err)
		return nil, err
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
		return nil, err
	}

	resp.UserId = id
	resp.Token, err = jwt.GenToken(id, req.GetUsername())
	return
}
