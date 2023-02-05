package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u *User) AddUser(ctx context.Context, username, password string) (id int64, err error) {
	user := &model.User{
		Username: username,
		Password: password,
	}
	if err = u.db.WithContext(ctx).Model(model.User{}).Create(user).Error; err != nil {
		Log.Errorf("add user err: %v, user: %+v", err, user)
		return
	}
	return user.ID, nil
}
