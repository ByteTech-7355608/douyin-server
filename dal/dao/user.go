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
	// 根据用户名查询
	err = db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(user).Error
	// 1.用户名已存在
	if err == nil {
		err = constants.ErrUserExist
		Log.Errorf("check user err: %v, user: %+v", err, user)
		return
	}
	// 2.用户名未存在--新增用户
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
	// 3.数据库错误
	err = constants.ErrCreateRecord
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
	if err = db.WithContext(ctx).Model(model.User{}).Omit("created_at, updated_at, deleted_at").First(&user, uid).Error; err != nil {
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	Log.Warnf("user %v not found", uid)
		//	err = constants.ErrUserNotExist
		//}
		Log.Errorf("FindUserById  err: %v, uid: %+v", err, uid)
		return
	}
	return user, nil
}
