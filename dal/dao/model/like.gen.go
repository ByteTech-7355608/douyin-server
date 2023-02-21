// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

const TableNameLike = "like"

// Like mapped from table <like>
type Like struct {
	ID        int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time             `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time             `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UID       int64                 `gorm:"column:uid" json:"uid"`
	Vid       int64                 `gorm:"column:vid" json:"vid"`
	Action    bool                  `gorm:"column:action" json:"action"`
}

// TableName Like's table name
func (*Like) TableName() string {
	return TableNameLike
}
