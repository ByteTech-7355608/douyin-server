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

func (u *Messages) TableName() string {
	return "messages"
}

// GetUserMessages   get user messages by uid and to_uid
func GetUserMessages(ctx context.Context, uid, to_uid int64) ([]*string, error) {
	msg_list := make([]*string, 0)
	if err := DB.WithContext(ctx).Table("messages").Select("content").Where("uid = ? AND to_uid = ?", uid, to_uid).Find(&msg_list).Error; err != nil {
		return nil, err
	}
	return msg_list, nil
}

// CreateUser create user info
func CreateMessages(ctx context.Context, msg *Messages) error {
	res := DB.WithContext(ctx).Table("messages").Create(msg)
	return res.Error

}
func DeleteMessages(ctx context.Context, msg_id uint) error {
	return DB.WithContext(ctx).Delete(&Messages{}, msg_id).Error
}
