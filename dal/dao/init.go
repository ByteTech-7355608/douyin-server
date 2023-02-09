package dao

import (
	"ByteTech-7355608/douyin-server/pkg/constants"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Dao struct {
	User  User
	Video Video
}

func InitDB() {
	var err error
	db, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
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

func InitMockDB() (mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(
		mysql.New(
			mysql.Config{
				Conn:                      mockDB,
				SkipInitializeWithVersion: true,
			}),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	return
}

func NewDao() *Dao {
	if db == nil {
		InitDB()
	}
	return &Dao{
		User:  User{},
		Video: Video{},
	}
}
