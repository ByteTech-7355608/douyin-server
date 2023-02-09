package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/interaction"
	apiModel "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/interaction"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	rpcModel "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// CommentList
// @router /douyin/comment/list/ [GET]
func (h *Handler) CommentList(ctx context.Context, c *app.RequestContext) {
	req := apiModel.DouyinCommentListRequest{}
	rpcReq := &rpcModel.DouyinCommentListRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Interaction().Client().CommentList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := rpcModel.DouyinCommentListResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func (h *Handler) FavoriteList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinFavoriteListRequest{}
	rpcReq := &rpc.DouyinFavoriteListRequest{}

	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Interaction().Client().FavoriteList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &rpc.DouyinFavoriteListResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

// CommentAction .
// @router /douyin/user/action [POST]
func (h *Handler) CommentAction(ctx context.Context, c *app.RequestContext) {
	req := api.DouyinCommentActionRequest{}
	rpcReq := rpc.DouyinCommentActionRequest{}
	if h.Pre(ctx, c, &req, &rpcReq) {
		userID, ok := c.Get("userid")
		if !ok {
			return
		}
		username, ok := c.Get("username")
		if !ok {
			return
		}
		rpcReq.BaseReq = new(model.BaseReq)
		rpcReq.BaseReq.UserId = new(int64)
		rpcReq.BaseReq.Username = new(string)
		*rpcReq.BaseReq.UserId = userID.(int64)
		*rpcReq.BaseReq.Username = username.(string)
		rpcResp, err := h.RPC().Interaction().Client().CommentAction(ctx, &rpcReq)
		if err != nil {
			return
		}
		resp := rpc.DouyinCommentActionResponse{}
		h.After(ctx, c, &resp, &rpcResp, err)
	}
}
