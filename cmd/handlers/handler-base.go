package handlers

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/rpc"
	base2 "ByteTech-7355608/douyin-server/service/base"
	"ByteTech-7355608/douyin-server/util"
	"context"

	"github.com/sirupsen/logrus"
)

var _ base.BaseService = new(BaseServiceImpl)

// BaseServiceImpl implements the last service interface defined in the IDL.
type BaseServiceImpl struct {
	rpc *rpc.RPC
	svc *base2.Service
}

// NewBaseServiceImpl returns a new BaseServiceImpl with the provided base and RPC client
func NewBaseServiceImpl() *BaseServiceImpl {
	return &BaseServiceImpl{}
}

func (s *BaseServiceImpl) Init(rpc *rpc.RPC) {
	s.rpc = rpc
	s.svc = base2.NewService(nil, rpc)
}

// Feed implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Feed(ctx context.Context, req *base.DouyinFeedRequest) (resp *base.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest) (resp *base.DouyinUserRegisterResponse, err error) {
	logrus.Infof("UserRegister args: %v", util.LogStr(req))
	return s.svc.UserRegister(ctx, req)
}

// UserLogin implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserLogin(ctx context.Context, req *base.DouyinUserLoginRequest) (resp *base.DouyinUserLoginResponse, err error) {
	logrus.Infof("UserLogin args: %v", util.LogStr(req))
	return s.svc.UserLogin(ctx, req)
}

// UserMsg implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserMsg(ctx context.Context, req *base.DouyinUserRequest) (resp *base.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (resp *base.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (resp *base.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
