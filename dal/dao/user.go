package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	"context"

	"github.com/sirupsen/logrus"
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
		logrus.Errorf("add user err: %v, user: %+v", err, user)
		return
	}

	return user.ID, nil
}

func (u *User) CheckUser(ctx context.Context, username, password string) (id int64, flag bool, err error) {
	user := &model.User{}

	if err = u.db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		logrus.Errorf("check user err: %v, user: %+v", err, user)
		return
	}

	// 检查密码是否正确
	if user.Password != password {
		return 0, false, nil
	}

	return user.ID, true, nil
}
