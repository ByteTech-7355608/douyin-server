package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/interaction"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// CommentAction .
// @router /douyin/user/action [POST]
func (h *Handler) CommentAction(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinCommentActionRequest{}
	rpcReq := &rpc.DouyinCommentActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Interaction().Client().CommentAction(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := rpc.DouyinCommentActionResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}

}
