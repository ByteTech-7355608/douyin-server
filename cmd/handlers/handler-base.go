package handlers

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	. "ByteTech-7355608/douyin-server/pkg/configs"
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
	s.svc = base2.NewService(rpc)
}

// Feed implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Feed(ctx context.Context, req *base.DouyinFeedRequest) (resp *base.DouyinFeedResponse, err error) {
	Log.Infof("Feed req: %v", util.LogStr(req))
	resp, err = s.svc.Feed(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("Feed resp: %v", util.LogStr(resp))
	return resp, nil
}

// UserRegister implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest) (resp *base.DouyinUserRegisterResponse, err error) {
	Log.Infof("UserRegister req: %v", util.LogStr(req))
	resp, err = s.svc.UserRegister(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("UserRegister resp: %v", util.LogStr(resp))
	return resp, nil
}

// UserLogin implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserLogin(ctx context.Context, req *base.DouyinUserLoginRequest) (resp *base.DouyinUserLoginResponse, err error) {
	logrus.Infof("UserLogin args: %v", util.LogStr(req))
	resp, err = s.svc.UserLogin(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("UserLogin resp: %v", util.LogStr(resp))
	return resp, nil
}

// UserMsg implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserMsg(ctx context.Context, req *base.DouyinUserRequest) (resp *base.DouyinUserResponse, err error) {
	// TODO: Your code here...
	Log.Infof("UserMsg req: %v", util.LogStr(req))
	resp, err = s.svc.UserMsg(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("UserMsg resp: %v", util.LogStr(resp))
	return resp, nil
}

// PublishAction implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (resp *base.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (resp *base.DouyinPublishListResponse, err error) {
	logrus.Infof("PublishList args: %v", util.LogStr(req))
	resp, err = s.svc.PublishList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("PublishList resp: %v", util.LogStr(resp))
	return resp, nil
}
