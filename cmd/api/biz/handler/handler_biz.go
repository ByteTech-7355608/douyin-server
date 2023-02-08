package handler

import (
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/util"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

type Handler struct {
	rpc *rpc.RPC
}

func (h *Handler) RPC() *rpc.RPC {
	return h.rpc
}

func NewHandler(rpc *rpc.RPC) *Handler {
	return &Handler{
		rpc: rpc,
	}
}

// Pre 绑定参数，并且copy到rpcReq
// rpcReq: nil 或 指针
func (h *Handler) Pre(ctx context.Context, c *app.RequestContext, req interface{}, rpcReq interface{}) (ok bool) {
	if err := c.BindAndValidate(req); err != nil {
		Log.Errorf("bind to %v error: %v", req, err)
		c.String(consts.StatusBadRequest, err.Error())
		return false
	}

	Log.Infof("req %T: %v", req, util.LogStr(req))
	if rpcReq != nil {
		if err := copier.Copy(rpcReq, req); err != nil {
			Log.Errorf("copy from %T to %T error: %v", req, rpcReq, err)
			c.String(consts.StatusBadRequest, err.Error())
			return false
		}
		Log.Infof("rpc req %T: %v", rpcReq, util.LogStr(rpcReq))
	}
	return true
}

// After .
func (h *Handler) After(ctx context.Context, c *app.RequestContext, resp interface{}, rpcResp interface{}, err error) (ok bool) {
	Log.Infof("rpc resp: %v", util.LogStr(rpcResp))
	c.JSON(consts.StatusOK, rpcResp)
	return true
}

//// AfterData 只返回resp的data字段
//func (h *Handler) afterData(ctx context.Context, c *app.RequestContext, resp interface{}, rpcResp interface{}, err error) {
//	Log.Infof("resp type: %T, rpcResp type: %T err: %v", resp, rpcResp, err)
//	if rpcResp != nil && !reflect.ValueOf(rpcResp).IsZero() {
//		if err := copier.Copy(resp, rpcResp); err != nil {
//			Log.Errorf("copy from %T to %T error: %v", rpcResp, resp, err)
//			c.String(consts.StatusInternalServerError, err.Error())
//			return
//		}
//	}
//	Log.Infof("resp: %+v", resp)
//	Log.Infof("rpc resp: %+v", rpcResp)
//	return
//}

type ResponseData struct {
	RespCode int32  `json:"resp_code"`
	RespMsg  string `json:"resp_msg"`
}

func Response(ctx context.Context, c *app.RequestContext, status *constants.RespStatus) {
	rd := &ResponseData{
		RespCode: status.StatusCode,
		RespMsg:  status.Error(),
	}
	c.JSON(http.StatusOK, rd)
}
