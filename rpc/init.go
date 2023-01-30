package rpc

import (
	"ByteTech-7355608/douyin-server/rpc/douyin/basecli"
	"ByteTech-7355608/douyin-server/rpc/douyin/interactioncli"
	"github.com/golang/mock/gomock"
)

type RPC struct {
	base        *basecli.Client
	interaction *interactioncli.Client
}

func (b *RPC) Base() *basecli.Client {
	return b.base
}

func (b *RPC) Interaction() *interactioncli.Client {
	return b.interaction
}

func NewRPC() *RPC {
	return &RPC{
		base:        basecli.NewClient(nil),
		interaction: interactioncli.NewClient(nil),
	}
}

func NewMockRPC(ctrl *gomock.Controller) *RPC {
	return &RPC{
		base:        basecli.NewClient(basecli.NewMockClient(ctrl)),
		interaction: interactioncli.NewClient(interactioncli.NewMockClient(ctrl)),
	}
}
