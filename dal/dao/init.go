package dao

import (
	"ByteTech-7355608/douyin-server/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	User  User
	Video Video
}

func getDB() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	return
}

func NewDao(db *gorm.DB) *Dao {
	if db == nil {
		db = getDB()
	}
	return &Dao{
		User:  User{db: db},
		Video: Video{db: db},
	}
}
