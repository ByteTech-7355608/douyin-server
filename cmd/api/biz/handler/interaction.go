package handler

import (
	"context"

	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/interaction"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"

	"github.com/cloudwego/hertz/pkg/app"
)

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
