package handlers

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	interaction2 "ByteTech-7355608/douyin-server/service/interaction"
	"ByteTech-7355608/douyin-server/util"
	"context"
)

var _ interaction.InteractionService = new(InteractionServiceImpl)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct {
	rpc *rpc.RPC
	svc *interaction2.Service
}

// NewInteractionServiceImpl returns a new InteractionServiceImpl with the provided base and RPC client
func NewInteractionServiceImpl() *InteractionServiceImpl {
	return &InteractionServiceImpl{}
}

func (s *InteractionServiceImpl) Init(rpc *rpc.RPC) {
	s.rpc = rpc
	s.svc = interaction2.NewService(rpc)
}

// FavoriteAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.DouyinFavoriteActionRequest) (resp *interaction.DouyinFavoriteActionResponse, err error) {
	Log.Infof("FavoriteAction args: %v", util.LogStr(req))
	resp, err = s.svc.FavoriteAction(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("FavoriteAction resp: %v", util.LogStr(resp))
	return resp, nil
}

// FavoriteList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	Log.Infof("FavoriteList args: %v", util.LogStr(req))
	resp, err = s.svc.FavoriteList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("FavoriteList resp: %v", util.LogStr(resp))
	return resp, nil
}

// CommentAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	Log.Infof("CommentAction req: %v", util.LogStr(req))
	resp, err = s.svc.CommentAction(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("CommentAction resp: %v", util.LogStr(resp))
	return resp, nil
}

// CommentList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentList(ctx context.Context, req *interaction.DouyinCommentListRequest) (resp *interaction.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	Log.Infof("CommentList req: %v", util.LogStr(req))
	resp, err = s.svc.CommentList(ctx, req)
	HandlerErr(resp, err)
	Log.Infof("CommentList resp: %v", util.LogStr(resp))
	return resp, nil
}
