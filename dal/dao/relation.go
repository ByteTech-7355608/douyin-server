package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"errors"
	"fmt"

	"context"

	"gorm.io/gorm"
)

type Relation struct {
}

// IsUserFollowed 两个用户有是否关注 输入两个用户的Id a->b
func (r *Relation) IsUserFollowed(ctx context.Context, concernerID int64, concernedID int64) (isFollow bool, err error) {
	relation := model.Relation{}
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("action").Where("concerner_id = ? AND concerned_id = ?", concernerID, concernedID).First(&relation).Error; err != nil {
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
	return relation.Action, nil
}

// FollowListLen 获取关注的数目，输入用户id，输出用户关注的人数量的信息和报错信息
func (r *Relation) FollowListLen(ctx context.Context, id int64) (nums *int64, err error) {
	var len int64
	err = db.WithContext(ctx).Model(model.Relation{}).Where("concerner_id=?", id).Count(&len).Error
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
	err = db.WithContext(ctx).Model(model.Relation{}).Where("concerned_id=?", id).Count(&len).Error
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

// 更改action位
func (r *Relation) UpdatedRelation(ctx context.Context, record *model.Relation, action int32) error {
	return db.WithContext(ctx).Model(&record).Update("action", action).Error
}

// 查看关注列表
func (r *Relation) FollowList(ctx context.Context, id int64) (list []*model.User, err error) {
	var user_ids []int64
	err = db.Debug().Model(model.Relation{}).Select("concerned_id").Where("concerner_id=? AND action=1", id).Find(&user_ids).Error
	if err != nil {
		Log.Errorf("get follow list fail,err:%v", err)
		return
	}
	for _, i := range user_ids {
		var user *model.User
		err = db.Where("id=?", i).First(&user).Error
		if err != nil {
			Log.Errorf("get userinfo fail,err:%v", err)
			return
		}
		list = append(list, user)
	}
	return
}

// CheckRecord查看两个用户的记录是否存在
func (r *Relation) CheckRecord(ctx context.Context, concernerID int64, concernedID int64) (flag bool, record *model.Relation, err error) {
	fmt.Println(concernerID)
	fmt.Println(concernedID)
	if err = db.WithContext(ctx).Model(model.Relation{}).Where("concerner_id = ? AND concerned_id = ?", concernerID, concernedID).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// relation 中不存在关注的关系，也可表示未关注
			flag = false
			err = nil
			Log.Infof("IsUserFollowed err：%v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
			return
		} else {
			// 数据库出错
			flag = false
			Log.Errorf("get follow relation err: %v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
			return
		}
	}
	flag = true
	return
}

// AddRelation  a关注b
func (r *Relation) AddRelation(ctx context.Context, concernerID int64, concernedID int64) error {
	follow := model.Relation{
		ConcernerID: concernerID,
		ConcernedID: concernedID,
		Action:      true,
	}
	return db.WithContext(ctx).Create(&follow).Error
}
