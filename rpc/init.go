package rpc

import (
	"ByteTech-7355608/douyin-server/rpc/douyin/basecli"
	"ByteTech-7355608/douyin-server/rpc/douyin/interactioncli"
	socialcli "ByteTech-7355608/douyin-server/rpc/douyin/social"

	"github.com/golang/mock/gomock"
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
		base:        basecli.NewClient(nil),
		interaction: interactioncli.NewClient(nil),
		social:      socialcli.NewClient(nil),
	}
}

func NewMockRPC(ctrl *gomock.Controller) *RPC {
	return &RPC{
		base:        basecli.NewClient(basecli.NewMockClient(ctrl)),
		interaction: interactioncli.NewClient(interactioncli.NewMockClient(ctrl)),
		social:      socialcli.NewClient(socialcli.NewMockClient(ctrl)),
	}
}
