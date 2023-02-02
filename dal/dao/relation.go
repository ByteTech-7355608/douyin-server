package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)

type Relation struct {
	db *gorm.DB
}

type RelationStruct model.Relation

func (r *Relation) UserById(ctx context.Context, id uint) (user *model.User, err error) {
	err = r.db.WithContext(ctx).Select("id", "username", "follow_count", "follower_count").Where("id=?", id).First(&user).Error
	return
}

// 关注用户之后修改用户的关注数和被关注者的粉丝数
func (r *RelationStruct) AfterCreate(tx *gorm.DB) (err error) {
	concerner_id := r.Concerner_id
	concerned_id := r.Concerned_id
	err = tx.Model(&model.User{}).Where("id=?", concerner_id).UpdateColumn("follow_count", gorm.Expr("follow_count+1")).Error
	if err != nil {
		return
	}
	err = tx.Model(&model.User{}).Where("id=?", concerned_id).UpdateColumn("follower_count", gorm.Expr("follower_count+1")).Error
	if err != nil {
		return
	}
	return
}

func (r *Relation) isRecordExist(ctx context.Context, concerner_id uint, concerned_id uint) bool {
	var count int64
	err := r.db.WithContext(ctx).Model(&Relation{}).Where("concerner_id=? And concerned_id=?", concerner_id, concerned_id).Count(&count).Error
	return err == nil && count > 0
}

// 添加关注，输入关注信息，返回报错信息
func (r *Relation) RelationAdd(ctx context.Context, relation *model.Relation) error {
	return r.db.WithContext(ctx).Create(&relation).Error
}

// 取消关注之前，减少用户的关注数和被关注的粉丝数
func (r *RelationStruct) BeforeDelete(tx *gorm.DB) (err error) {
	concerner_id := r.Concerner_id
	concerned_id := r.Concerned_id
	err = tx.Model(&model.User{}).Where("id=?", concerner_id).UpdateColumn("follow_count", gorm.Expr("follow_count-1")).Error
	if err != nil {
		return
	}
	err = tx.Model(&model.User{}).Where("id=?", concerned_id).UpdateColumn("follower_count", gorm.Expr("follower_count-1")).Error
	if err != nil {
		return
	}
	return
}

// 取消关注，输入关注者id和被关注者id，输出报错信息
func (r *Relation) RelationDel(ctx context.Context, concerner_id uint, concerned_id uint) error {
	var relation *RelationStruct
	err := r.db.WithContext(ctx).Where("concerner_id=? And concerned_id=?", concerner_id, concerned_id).First(&relation).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(&RelationStruct{
		Model: gorm.Model{
			ID: relation.ID,
		},
		Concerner_id: concerner_id,
		Concerned_id: concerned_id,
	}).Error
}

// 获取关注列表，输入用户id，输出关注的人的信息和报错信息
func (r *Relation) FollowList(ctx context.Context, id uint) (users []*model.User, err error) {
	var relations []*RelationStruct
	err = r.db.WithContext(ctx).Select("concerned_id").Where("concerner_id=?", id).Find(&relations).Error
	if err != nil {
		return
	}
	ids := make([]uint, len(relations))
	for i := 0; i < len(relations); i++ {
		ids[i] = relations[i].Concerned_id
	}
	err = r.db.WithContext(ctx).Where("id in ?", ids).Find(&users).Error
	return
}

// 获取粉丝列表，输入用户id，输出粉丝列表和报错信息
func (r *Relation) FollowerList(ctx context.Context, id uint) (users []*model.User, err error) {
	var relations []*RelationStruct
	err = r.db.WithContext(ctx).Select("concerner_id").Where("concerned_id=?", id).Find(&relations).Error
	if err != nil {
		return
	}
	ids := make([]uint, len(relations))
	for i := 0; i < len(relations); i++ {
		ids[i] = relations[i].Concerner_id
	}
	err = r.db.WithContext(ctx).Where("id in ?", ids).Find(&users).Error
	return
}

// 判断一个用户是否关注另外一个用户，输入关注用户id和是否被关注用户id
func (r *Relation) IsConcern(ctx context.Context, concerner_id uint, concerned_id uint) bool {
	var count int64
	var relation *RelationStruct
	err := r.db.WithContext(ctx).Where("concerner_id=? AND concerned_id=?", concerner_id, concerned_id).Find(&relation).Count(&count).Error
	if err != nil {
		fmt.Println(err)
	}
	return count > 0 && err == nil
}

// 获取朋友列表，输入用户id，输出朋友信息和报错信息
func (r *Relation) FriendList(ctx context.Context, id uint) (users []*model.User, err error) {
	var relation *RelationStruct
	var rows *sql.Rows
	rows, err = r.db.WithContext(ctx).Model(&relation).Where("concerner_id=?", id).Rows()
	if err != nil {
		return
	}
	var count int64
	var user *model.User
	defer rows.Close()
	for rows.Next() {
		r.db.WithContext(ctx).ScanRows(rows, &relation)
		err = r.db.WithContext(ctx).Model(&Relation{}).Where("concerner_id=? AND concerned_id=?", relation.Concerned_id, id).Count(&count).Error
		if err != nil {
			return
		}
		if count > 0 {
			user, err = r.UserById(ctx, relation.Concerned_id)
			if err != nil {
				return
			}
			users = append(users, user)
		}
	}
	return
}
