package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"errors"

	"context"

	"gorm.io/gorm"
)

type Relation struct {
}

// AddRelation  a关注b
func (r *Relation) AddRelation(ctx context.Context, concernerID int64, concernedID int64) (err error) {
	follow:=model.Relation{
		ConcernerID: concernerID,
		ConcernedID: concernedID,
		Action: 1,
	}
	

}
// DeleteRelation  a取消关注b
func(r *Relation) DeleteRelation(ctx context.Context,concernerID int64, concernedID int64){

}

// IsUserFollowed 两个用户有是否关注 输入两个用户的Id a->b
func (r *Relation) IsUserFollowed(ctx context.Context, concernerID int64, concernedID int64) (isFollow bool, err error) {
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("action").Where("concerner_id = ? AND concerned_id = ?", concernerID, concernedID).First(&isFollow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// relation 中不存在关注的关系，也可表示未关注
			Log.Infof("IsUserFollowed err：%v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
			return false, nil
		} else {
			// 数据库出错
			Log.Errorf("get follow relation err: %v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
			return false, err
		}
	}
	return
}

// FollowListLen 获取关注的数目，输入用户id，输出用户关注的人数量的信息和报错信息
func (r *Relation) FollowListLen(ctx context.Context, id int64) (nums *int64, err error) {
	var len int64
	err = db.WithContext(ctx).Model(model.Relation{}).Where("concerner_id=? AND action = 1", id).Count(&len).Error
	if err != nil {
		Log.Errorf("get concerner-nums err: %v, userId: %+v", err, id)
		return
	}
	nums = &len
	err = nil
	return
}

// FollowerListLen 获取粉丝数目，输入用户id，输出用户的粉丝数目和报错信息
func (r *Relation) FollowerListLen(ctx context.Context, id int64) (nums *int64, err error) {
	var len int64
	err = db.WithContext(ctx).Model(model.Relation{}).Where("concerned_id=? AND action = 1", id).Count(&len).Error
	if err != nil {
		Log.Errorf("get concerned-nums err: %v, userId: %+v", err, id)
		return
	}
	nums = &len
	err = nil
	return
}

// IsFollower a是否关注b
func (r *Relation) IsFollower(ctx context.Context, a, b int64) (follower bool, err error) {
	var action bool
	// TODO 建立唯一索引，提高查询效率
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("action").Where("concerner_id = ? AND concerned_id = ?", a, b).Find(&action).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		Log.Errorf("query relation between %v and %v err: %v", a, b, err)
		return false, err
	}
	return action, nil
}
