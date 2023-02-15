package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"errors"

	"context"

	"gorm.io/gorm"
)

type User struct {
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
	Log.Errorf("check user err: %v, user: %+v", err, user)
	return
}

func (u *User) CheckUser(ctx context.Context, username, password string) (id int64, err error) {
	user := &model.User{}

	if err = db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = constants.ErrUserNotExist
		}
		Log.Errorf("check user err: %v, user: %+v", err, user)
		return
	}

	// 检查密码是否正确
	if util.EncryptPassword(password) != user.Password {
		err = constants.ErrInvalidPassword
		Log.Errorf("check user err: %v, user: %+v", err, user)

		return
	}

	return user.ID, nil
}

func (u *User) FindUserById(ctx context.Context, uid int64) (user model.User, err error) {
	if err = db.WithContext(ctx).Model(model.User{}).Omit("created_at, updated_at, deleted_at").Where("id = ?", uid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = constants.ErrUserNotExist
		}
		Log.Errorf("FindUserById  err: %v, uid: %+v", err, uid)
		return
	}

	return user, nil
}

// // 使用id获取用户信息
//
//	func (u *User) FindUserNameById(ctx context.Context, id int64) (username string, err error) {
//		user:=model.User{}
//		if err = db.WithContext(ctx).Model(model.User{}).Select("username","follow_count","follower_count").Where("id = ?", id).Fin(&user).Error; err != nil {
//			Log.Errorf("find username err: %v, user_id: %+v", err, id)
//			return "", err
//		}
//		username = name
//		err = nil
//		return
//	}
func (u *User) QueryUser(ctx context.Context, userID int64) (user *model.User, err error) {
	if err = db.WithContext(ctx).Model(model.User{}).Where("id = ?", userID).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Log.Warnf("user %v not found", userID)
			return nil, constants.ErrUserNotExist
		}
		Log.Errorf("query user %v err: %v", userID, err)
		return nil, err
	}
	return
}
