package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"context"

	"github.com/sirupsen/logrus"
)

func (s *Service) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest) (resp *base.DouyinUserRegisterResponse, err error) {
	resp = base.NewDouyinUserRegisterResponse()
	id, err := s.dao.User.AddUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		logrus.Errorf("add user err: %v", err)
		return nil, err
	}
	resp.StatusCode = 200
	resp.UserId = int64(id)
	return
}

func (s *Service) UserLogin(ctx context.Context, req *base.DouyinUserLoginRequest) (resp *base.DouyinUserLoginResponse, err error) {
	resp = base.NewDouyinUserLoginResponse()

	id, ok, err := s.dao.User.CheckUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		logrus.Errorf("user login err: %v", err)
		return nil, err
	}
	if !ok {
		resp.StatusCode = 444
		return
	}
	resp.StatusCode = 200
	resp.UserId = int64(id)
	return
}
