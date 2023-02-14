package dao

import (
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var db *gorm.DB

type Dao struct {
	User     User
	Video    Video
	Like     Like
	Relation Relation
	Comment  Comment
	Message  Message
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
		Log.Errorf("Init DB err: %+v", err)
		panic(err)
	}
	err = db.Use(gormopentracing.New())
	if err != nil {
		Log.Errorf("Use gorm-opentracing err: %+v", err)
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
		User:     User{},
		Video:    Video{},
		Like:     Like{},
		Relation: Relation{},
		Comment:  Comment{},
		Message:  Message{},
	}
}
