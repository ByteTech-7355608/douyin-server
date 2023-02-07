// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	 "gorm.io/plugin/soft_delete"
)

const TableNameRelation = "relation"

// Relation mapped from table <relation>
type Relation struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ConcernerID int64          `gorm:"column:concerner_id;not null" json:"concerner_id"`
	ConcernedID int64          `gorm:"column:concerned_id;not null" json:"concerned_id"`
	Action      bool           `gorm:"column:action" json:"action"`
}

// TableName Relation's table name
func (*Relation) TableName() string {
	return TableNameRelation
}
