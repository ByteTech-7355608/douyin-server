package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u *User) AddUser(ctx context.Context, username, password string) (id int64, err error) {
	user := &model.User{}

	// 检查当前用户名是否已经存在
	if err = db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user = &model.User{
				Username: username,
				Password: util.EncryptPassword(password),
			}
			if err = db.WithContext(ctx).Model(model.User{}).Create(user).Error; err != nil {
				Log.Errorf("add user err: %v, user: %+v", err, user)
				return
			}
			return user.ID, nil
		}
		return
	}

	err = constants.ErrUserExist
	logrus.Errorf("check user err: %v, user: %+v", err, user)
	return
}

func (u *User) CheckUser(ctx context.Context, username, password string) (id int64, err error) {
	user := &model.User{}

	if err = db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = constants.ErrUserNotExist
		}
		logrus.Errorf("check user err: %v, user: %+v", err, user)
		return
	}

	// 检查密码是否正确
	if util.EncryptPassword(password) != user.Password {
		err = constants.ErrInvalidPassword
		logrus.Errorf("check user err: %v, user: %+v", err, user)

		return
	}

	return user.ID, nil
}
