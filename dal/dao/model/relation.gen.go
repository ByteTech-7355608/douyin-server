package model

import "gorm.io/gorm"

const TableNameRelation = "relation"

type Relation struct {
	gorm.Model
	Concerner_id uint `gorm:"type:uint;not null"`
	Concerned_id uint `gorm:"type:uint;not null"`
}

func (*Relation) TableName() string {
	return TableNameRelation
}
