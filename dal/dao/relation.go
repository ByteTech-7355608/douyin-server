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
	tx := db.WithContext(ctx).Begin()
	err := tx.Model(&record).Update("action", action).Error
	concerner_id := record.ConcernerID
	concerned_id := record.ConcernedID
	if err != nil {
		tx.Rollback()
		Log.Errorf("update user:%v and user:%v relation err:%v", concerner_id, concerned_id, err)
		return err
	}
	if action == 1 {
		err = tx.Model(model.User{}).Where("id=?", concerner_id).UpdateColumn("follow_count", gorm.Expr("follow_count+1")).Error
		if err != nil {
			Log.Errorf("update user:%v follow_count err:%v", concerner_id, err)
			tx.Rollback()
			return err
		}
		err = tx.Model(model.User{}).Where("id=?", concerned_id).UpdateColumn("follower_count", gorm.Expr("follower_count+1")).Error
		if err != nil {
			Log.Errorf("update user:%v follower_count err:%v", concerned_id, err)
			tx.Rollback()
			return err
		}
		tx.Commit()
	} else {
		err = tx.Model(model.User{}).Where("id=?", concerner_id).UpdateColumn("follow_count", gorm.Expr("follow_count-1")).Error
		if err != nil {
			Log.Errorf("update user:%v follow_count err:%v", concerner_id, err)
			tx.Rollback()
			return err
		}
		err = tx.Model(model.User{}).Where("id=?", concerned_id).UpdateColumn("follower_count", gorm.Expr("follower_count-1")).Error
		if err != nil {
			Log.Errorf("update user:%v follower_count err:%v", concerned_id, err)
			tx.Rollback()
			return err
		}
		tx.Commit()
	}
	return nil
}

// CheckRecord查看两个用户的记录是否存在
func (r *Relation) CheckRecord(ctx context.Context, concernerID int64, concernedID int64) (flag bool, record *model.Relation, err error) {
	if err = db.WithContext(ctx).Model(model.Relation{}).Omit("created_at, updated_at, deleted_at").Where("concerner_id = ? AND concerned_id = ?", concernerID, concernedID).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// relation 中不存在关注的关系，也可表示未关注
			flag = false
			err = nil
			Log.Infof("CheckRecord err：%v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
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
	concerner_id := concernerID
	concerned_id := concernedID
	tx := db.WithContext(ctx).Begin()
	err := tx.Create(&follow).Error
	if err != nil {
		Log.Infof("AddRelation err: %v, concerner_id: %d AND concerned_id: %d", err, concernerID, concernedID)
		tx.Rollback()
		return err
	}
	err = tx.Model(model.User{}).Where("id=?", concerner_id).UpdateColumn("follow_count", gorm.Expr("follow_count+1")).Error
	if err != nil {
		Log.Infof("AddRelation err: %v, concerner_id: %d AND concerned_id: %d", err, concernerID, concernedID)
		tx.Rollback()
		return err
	}
	err = tx.Model(model.User{}).Where("id=?", concerned_id).UpdateColumn("follower_count", gorm.Expr("follower_count+1")).Error
	if err != nil {
		Log.Infof("AddRelation err: %v, concerner_id: %d AND concerned_id: %d", err, concernerID, concernedID)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
