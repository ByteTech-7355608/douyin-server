package interactionimport

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/rpc"

	"gorm.io/gorm"
)

type Service struct {
	dao *dao.Dao
	rpc *rpc.RPC
}

func NewService(db *gorm.DB, rpc *rpc.RPC) *Service {
	return &Service{
		dao: dao.NewDao(db),
		rpc: rpc,
	}
}
