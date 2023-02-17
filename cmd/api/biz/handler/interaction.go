package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/interaction"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// CommentList
// @router /douyin/comment/list/ [GET]
func (h *Handler) CommentList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinCommentListRequest{}
	rpcReq := &rpc.DouyinCommentListRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Interaction().Client().CommentList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinCommentListResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func (h *Handler) FavoriteList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinFavoriteListRequest{}
	rpcReq := &rpc.DouyinFavoriteListRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Interaction().Client().FavoriteList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinFavoriteListResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// CommentAction .
// @router /douyin/user/action [POST]
func (h *Handler) CommentAction(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinCommentActionRequest{}
	rpcReq := &rpc.DouyinCommentActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Interaction().Client().CommentAction(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinCommentActionResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// FavoriteAction .
// @router /douyin/favorite/action [POST]
func (h *Handler) FavoriteAction(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinFavoriteActionRequest{}
	rpcReq := &rpc.DouyinFavoriteActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Interaction().Client().FavoriteAction(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinCommentActionResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}

}
