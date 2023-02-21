package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"errors"

	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Like struct {
}

func (l *Like) QueryUserLikeRecords(ctx context.Context, uid int64) (userLikes []model.Like, err error) {
	if err = db.WithContext(ctx).Model(model.Like{}).Select("vid").Where("uid = ? AND action = ?", uid, 1).Find(&userLikes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 该用户没有点赞过的任何视频，不影响返回结果
			Log.Infof("uid %v dont like any videos", err)
			return
		}
		// 数据库出错
		Log.Errorf("find like err: %v, uid: %d", err, uid)
	}
	return
}

// IsLike 用户是否点赞了视频
func (l *Like) IsLike(ctx context.Context, uid, vid int64) (like bool, err error) {
	var action bool
	// TODO 建立唯一索引，提高查询效率
	if err = db.WithContext(ctx).Model(model.Like{}).Select("action").Where("uid = ? AND vid = ?", uid, vid).Find(&action).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		Log.Errorf("query relation between user %v and video %v err: %v", uid, vid, err)
		return false, err
	}
	return action, nil
}

func (l *Like) QueryRecord(ctx context.Context, uid, vid int64) (like *model.Like, err error) {
	if err = db.WithContext(ctx).Model(model.Like{}).Where("uid = ? AND vid = ?", uid, vid).First(&like).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Log.Infof("%v not like %v", uid, vid)
			return
		}
		Log.Errorf("query relation between user %v and video %v err: %v", uid, vid, err)
	}
	return
}

func (l *Like) UpsertRecord(ctx context.Context, record *model.Like) (err error) {
	if err = db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}, {Name: "vid"}},
		DoUpdates: clause.AssignmentColumns([]string{"action"}),
	}).Create(&record).Error; err != nil {
		Log.Errorf("upsert like record err: %v", err)
	}
	return
}
