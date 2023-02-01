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
	"ByteTech-7355608/douyin-server/pkg/constants"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Messages struct {
	gorm.Model
	Uid     uint   `json:"uid"`
	To_uid  uint   `json:"to_uid"`
	Content string `json:"content"`
}

var DB *gorm.DB

// Init init DB
func Init() error {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return fmt.Errorf("数据库连接失败 %v", err)
	}
	fmt.Println("数据库连接成功！")
	return nil
}
