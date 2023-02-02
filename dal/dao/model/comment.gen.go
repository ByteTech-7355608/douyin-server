package model

import "gorm.io/gorm"

const TableNameComment = "comment"

type Comment struct {
	gorm.Model
	Vid     uint   `gorm:"type:uint;not null"`
	Uid     uint   `gorm:"type:uint;not null"`
	Content string `gorm:"type:string;not null"`
}

// TableName User's table name
func (*Comment) TableName() string {
	return TableNameComment
}
