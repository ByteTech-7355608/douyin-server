package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"math/rand"
	"strconv"
	"time"

	"context"
	"errors"

	"gorm.io/gorm"
)

type User struct {
}

var signature = "hello world"

func (u *User) AddUser(ctx context.Context, username, password string) (id int64, err error) {
	user := &model.User{}

	// 检查当前用户名是否已经存在
	if err = db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			rand.Seed(time.Now().UnixNano())
			b := rand.Intn(constants.Image_nums)
			user = &model.User{
				Username:        username,
				Password:        util.EncryptPassword(password),
				Signature:       signature,
				Avatar:          constants.Avatar_url + strconv.Itoa(b),
				BackgroundImage: constants.Bg_url + strconv.Itoa(b),
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

func (u *User) QueryUser(ctx context.Context, userID int64) (user *model.User, err error) {
	if err = db.WithContext(ctx).Model(model.User{}).Where("id = ?", userID).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Log.Warnf("user %v not found", userID)
			return nil, constants.ErrUserNotExist
		}
		Log.Errorf("query user %v err: %v", userID, err)
		return nil, err
	}

	return user, nil
}
