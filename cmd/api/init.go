package api

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"
	"ByteTech-7355608/douyin-server/cmd/api/biz/router"
	"ByteTech-7355608/douyin-server/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertztracer "github.com/hertz-contrib/tracer/hertz"
	"github.com/opentracing/opentracing-go"
)

func NewDouyinApiHertz() *server.Hertz {
	svc := server.Default(
		server.WithTracer(hertztracer.NewTracer(opentracing.GlobalTracer(), func(c *app.RequestContext) string {
			return "hertz.server" + "::" + c.FullPath()
		})))
	svc.Use(hertztracer.ServerCtx())
	h := handler.NewHandler(rpc.NewRPC())
	router.Register(svc, h)
	return svc
}
