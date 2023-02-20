package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"errors"

	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Relation struct {
}

// 获取用户的关注id列表
func (r *Relation) FollowidList(ctx context.Context, id int64) (user_ids []int64, err error) {
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("concerned_id").Where("concerner_id=? AND action=1", id).Find(&user_ids).Error; err != nil {
		Log.Errorf("get follow list fail,err:%v", err)
		return nil, err
	}
	return
}

// 获取用户的粉丝id列表
func (r *Relation) FolloweridList(ctx context.Context, uid int64) (user_ids []int64, err error) {
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("concerner_id").Where("concerned_id=? AND action = ?", uid, 1).Find(&user_ids).Error; err != nil {
		Log.Errorf("get followeridlist err: %v, userid:%v", err, uid)
		return nil, err
	}
	return
}

// IsUserFollowed 两个用户有是否关注 输入两个用户的Id a->b
func (r *Relation) IsUserFollowed(ctx context.Context, concernerID int64, concernedID int64) (isFollow bool, err error) {
	relation := model.Relation{}
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("action").Where("concerner_id = ? AND concerned_id = ?", concernerID, concernedID).First(&relation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// relation 中不存在关注的关系，也可表示未关注
			Log.Infof("IsUserFollowed err: %v, concerner_id: %d AND concerned_id: %d", err, concernerID, concernedID)
			return false, nil
		} else {
			// 数据库出错
			Log.Errorf("get follow relation err: %v, concerner_id: %d AND concerned_id: %d", err, concernerID, concernedID)
			return false, err
		}
	}
	return relation.Action, nil
}

func (r *Relation) UpsertRecord(ctx context.Context, record *model.Relation) (err error) {
	if err = db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "concerner_id"}, {Name: "concerned_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"action"}),
	}).Create(&record).Error; err != nil {
		Log.Errorf("upsert relation record err: %v", err)
	}
	return
}
