package main

import (
	"ByteTech-7355608/douyin-server/cmd/api"
	"ByteTech-7355608/douyin-server/cmd/handlers"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base/baseservice"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction/interactionservice"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social/socialservice"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/rpc"
	"fmt"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

// main 服务入口，一个main启动多个服务
func main() {
	psm := os.Getenv("ServiceName")
	var svr server.Server
	switch psm {
	case constants.ApiServiceName:
		startDouyinApi()
		return
	case constants.BaseServiceName:
		svr = startDouyinBase()
	case constants.InteractionServiceName:
		svr = startDouyinInteraction()
	case constants.SocialServiceName:
		svr = startDouyinSocial()
	}
	logrus.Info("start service: %s", psm)
	if svr == nil {
		panic(fmt.Sprintf("no server for (%s) to run, support PSM: %v", psm,
			[]string{constants.ApiServiceName, constants.BaseServiceName, constants.InteractionServiceName}))
	}
	err := svr.Run()
	if err != nil {
		panic(err)
	}
}

func startDouyinApi() {
	svc := api.NewDouyinApiHertz()
	svc.Spin()
}

func startDouyinBase() (svr server.Server) {
	svc := handlers.NewBaseServiceImpl()
	svc.Init(rpc.NewRPC())
	svr = baseservice.NewServer(svc, serverOptions(constants.BaseServiceName, constants.BaseTCPAddr)...)
	return
}

func startDouyinInteraction() (svr server.Server) {
	svc := handlers.NewInteractionServiceImpl()
	svc.Init(rpc.NewRPC())
	svr = interactionservice.NewServer(svc, serverOptions(constants.InteractionServiceName, constants.InteractionTCPAddr)...)
	return
}

func startDouyinSocial() (svr server.Server) {
	svc := handlers.NewSocialServiceImpl()
	svc.Init(rpc.NewRPC())
	svr = socialservice.NewServer(svc, serverOptions(constants.SocialServiceName, constants.SocialTCPAddr)...)
	return
}

// serverOptions server启动配置
func serverOptions(serverName, tcpAddr string) (opts []server.Option) {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	//Init()
	opts = []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serverName}), // server name
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),
	}
	return
}
