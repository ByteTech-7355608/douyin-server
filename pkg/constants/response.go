package constants

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type ResponseData struct {
	RespCode ResCode     `json:"respcode"`
	RespMsg  interface{} `json:"respmsg"`
}

func Response(ctx context.Context, c *app.RequestContext, code ResCode) {
	rd := &ResponseData{
		RespCode: code,
		RespMsg:  code.Msg(),
	}
	c.JSON(http.StatusOK, rd)
}
