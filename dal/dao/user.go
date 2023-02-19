package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"math/rand"
	"time"

	"context"
	"errors"

	"gorm.io/gorm"
)

type User struct {
}

var bg_url = []string{
	"http://rqa2iqcgg.hd-bkt.clouddn.com/bg/ec1e98a2301ca281678c938128201e07.jpeg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/bg/d4605e47c4363cb18fdb29b1cead5bd0.jpeg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/bg/a85dd0618fdc2f5580d564ca3a009c01.jpeg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/bg/5a6208761f79c35feb47cc5358341c50.jpeg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/bg/2c170f3d0bc839800fbb6f2571033638.jpeg",
}

var avatar_url = []string{
	"http://rqa2iqcgg.hd-bkt.clouddn.com/avatar/de5ead300d8d568005131a1790b75ef1.jpg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/avatar/9074b82f422d6076524db6dc78c1ca55.jpg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/avatar/5368d72ddbb82d8e963ca776f47495de.jpg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/avatar/33eb353766f8b04686cfa64180e4e8ca.jpg",
	"http://rqa2iqcgg.hd-bkt.clouddn.com/avatar/30b626738756934f1412c57b5efd3cf1.jpg",
}
var signature = "hello world"

func (u *User) AddUser(ctx context.Context, username, password string) (id int64, err error) {
	user := &model.User{}

	// 检查当前用户名是否已经存在
	if err = db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			rand.Seed(time.Now().UnixNano())
			b := rand.Intn(len(avatar_url))
			user = &model.User{
				Username:        username,
				Password:        util.EncryptPassword(password),
				Signature:       signature,
				Avatar:          avatar_url[b],
				BackgroundImage: bg_url[b],
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
