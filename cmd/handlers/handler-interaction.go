package handlers

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/rpc"
	interaction2 "ByteTech-7355608/douyin-server/service/interaction"
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
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentList(ctx context.Context, req *interaction.DouyinCommentListRequest) (resp *interaction.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
