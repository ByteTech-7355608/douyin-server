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
	"context"
)

func (u *User) TableName() string {
	return "user"
}

// GetUserInfo   get user info by ID
func GetUserInfo(ctx context.Context, userID int64) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users *User) (uint, error) {
	res := DB.WithContext(ctx).Create(users)
	return users.ID, res.Error

}

// QueryUser query user info by name
func QueryUser(ctx context.Context, userName string) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
