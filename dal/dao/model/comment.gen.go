package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID        int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time             `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time             `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Vid       int64                 `gorm:"column:vid;not null" json:"vid"`
	UID       int64                 `gorm:"column:uid;not null" json:"uid"`
	Content   string                `gorm:"column:content;not null" json:"content"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
