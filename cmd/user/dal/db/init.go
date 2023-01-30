// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

var DB *gorm.DB

// Init init DB
func Init() error {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456@(127.0.0.1:3305)/db1?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		return fmt.Errorf("数据库连接失败 %v", err)
	}
	fmt.Println("数据库连接成功！")
	if DB.Migrator().HasTable(&User{}) {
		DB.Migrator().DropTable(&User{})
	}
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	return nil
}
