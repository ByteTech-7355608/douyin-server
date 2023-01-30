package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMockDB() (gormDB *gorm.DB, mock sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(
	//sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	if err != nil {
		panic(err)
	}
	gormDB, err = gorm.Open(
		mysql.New(
			mysql.Config{
				Conn:                      db,
				SkipInitializeWithVersion: true,
			}),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	return
}
