package handlers

import (
	social "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	social2 "ByteTech-7355608/douyin-server/service/social"
	"ByteTech-7355608/douyin-server/util"
	"context"
)

var _ social.SocialService = new(SocialServiceImpl)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct {
	rpc *rpc.RPC
	svc *social2.Service
}

// NewSocialServiceImpl returns a new SocialServiceImpl with the provided base and RPC client
func NewSocialServiceImpl() *SocialServiceImpl {
	return &SocialServiceImpl{}
}

func (s *SocialServiceImpl) Init(rpc *rpc.RPC) {
	s.rpc = rpc
	s.svc = social2.NewService(rpc)
}

// FollowAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FollowAction(ctx context.Context, req *social.DouyinFollowActionRequest) (resp *social.DouyinFollowActionResponse, err error) {
	Log.Infof("FavoriteAction args: %v", util.LogStr(req))
	resp, err = s.svc.FollowAction(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("FavoriteAction resp: %v", util.LogStr(resp))
	return resp, nil
}

// FollowList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FollowList(ctx context.Context, req *social.DouyinFollowingListRequest) (resp *social.DouyinFollowingListResponse, err error) {
	Log.Infof("FollowList req: %v", util.LogStr(req))
	resp, err = s.svc.FollowList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("FollowList resp: %v", util.LogStr(resp))
	return resp, nil
}

// FollowerList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FollowerList(ctx context.Context, req *social.DouyinFollowerListRequest) (resp *social.DouyinFollowerListResponse, err error) {
	Log.Infof("FollowerList req: %v", util.LogStr(req))
	resp, err = s.svc.FollowerList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("FollowerList resp: %v", util.LogStr(resp))
	return resp, nil
}

// FriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FriendList(ctx context.Context, req *social.DouyinRelationFriendListRequest) (resp *social.DouyinRelationFriendListResponse, err error) {
	Log.Infof("FriendListList req: %v", util.LogStr(req))
	resp, err = s.svc.FriendList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("FriendList resp: %v", util.LogStr(resp))
	return resp, nil
}

// MessageList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) MessageList(ctx context.Context, req *social.DouyinMessageChatRequest) (resp *social.DouyinMessageChatResponse, err error) {
	Log.Infof("MessageList req: %v", util.LogStr(req))
	resp, err = s.svc.MessageList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("MessageList resp: %v", util.LogStr(resp))
	return resp, nil
}

// SendMessage implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) SendMessage(ctx context.Context, req *social.DouyinMessageActionRequest) (resp *social.DouyinMessageActionResponse, err error) {
	Log.Infof("SendMessage req: %v", util.LogStr(req))
	resp, err = s.svc.SendMessage(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("SendMessage resp: %v", util.LogStr(resp))
	return resp, nil
}
