package main

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/dao/model",
	})

	db, _ := gorm.Open(mysql.Open(constants.MySQLDefaultDSN))
	g.UseDB(db)

	g.GenerateAllTable()

	g.Execute()
}
