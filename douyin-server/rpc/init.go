package rpc

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/pkg/tracer"
	"ByteTech-7355608/douyin-server/rpc/douyin/basecli"
	"ByteTech-7355608/douyin-server/rpc/douyin/interactioncli"
	"ByteTech-7355608/douyin-server/rpc/douyin/socialcli"
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/golang/mock/gomock"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

type RPC struct {
	base        *basecli.Client
	interaction *interactioncli.Client
	social      *socialcli.Client
}

func (r *RPC) Base() *basecli.Client {
	return r.base
}

func (r *RPC) Interaction() *interactioncli.Client {
	return r.interaction
}

func (r *RPC) Social() *socialcli.Client {
	return r.social
}

func NewRPC() *RPC {
	return &RPC{
		base:        basecli.NewClient(nil, clientOptions(constants.BaseServiceName)...),
		interaction: interactioncli.NewClient(nil, clientOptions(constants.InteractionServiceName)...),
		social:      socialcli.NewClient(nil, clientOptions(constants.SocialServiceName)...),
	}
}

func NewMockRPC(ctrl *gomock.Controller) *RPC {
	return &RPC{
		base:        basecli.NewClient(basecli.NewMockClient(ctrl)),
		interaction: interactioncli.NewClient(interactioncli.NewMockClient(ctrl)),
		social:      socialcli.NewClient(socialcli.NewMockClient(ctrl)),
	}
}

// clientOptions client启动配置
func clientOptions(serviceName string) (opts []client.Option) {
	kTracer, _ := tracer.InitTracer("kitex.client::" + serviceName)
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	opts = []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName, Tags: make(map[string]string)}), //service base info
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3 * time.Second),            // rpc timeout
		client.WithConnectTimeout(50 * time.Millisecond),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(opentracing.NewClientSuite(kTracer, func(c context.Context) string {
			endpoint := rpcinfo.GetRPCInfo(c).To()
			return endpoint.ServiceName() + "::" + endpoint.Method()
		})), // tracer
		client.WithResolver(r), // resolver
	}
	return
}
