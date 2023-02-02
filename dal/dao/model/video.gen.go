package model

import "gorm.io/gorm"

const TableNameVedio = "vedio"

type Video struct {
	gorm.Model
	Play_url       string `gorm:"type:string;not null"`
	Cover_url      string `gorm:"type:string;not null"`
	Favorite_count uint   `gorm:"type:uint;Default:0"`
	Comment_count  uint   `gorm:"type:uint;Default:0"`
	Title          string `gorm:"type:string;not null"`
	Uid            uint   `gorm:"type:uint;not null"`
}

// TableName User's table name
func (*Video) TableName() string {
	return TableNameVedio
}
