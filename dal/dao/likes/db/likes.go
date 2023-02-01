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

func (u *Likes) TableName() string {
	return "likes"
}

// GetUserLikes   get user likes video list by uid
func GetUserLikes(ctx context.Context, userID uint) ([]uint, error) {
	res := make([]uint, 0)
	err := DB.WithContext(ctx).Table("likes").Select("vid").Where("uid=?", userID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateLikes create user likes
func CreateLikes(ctx context.Context, lk *Likes) error {
	return DB.WithContext(ctx).Table("likes").Create(lk).Error
}

// DeleteLikes cancle user likes
func DeleteLikes(ctx context.Context, id uint) error {
	return DB.WithContext(ctx).Delete(&Likes{}, id).Error
}
