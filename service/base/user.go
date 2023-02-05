package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	. "ByteTech-7355608/douyin-server/pkg/configs"
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
	return
}
