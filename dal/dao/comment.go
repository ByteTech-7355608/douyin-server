package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	db *gorm.DB
}

type CommentStruct model.Comment

// 发布评论之前修改视频的评论数
func (c *CommentStruct) BeforeCreate(tx *gorm.DB) (err error) {
	vid := c.Vid
	err = tx.Model(&model.Video{}).Where("id=?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count+1")).Error
	if err != nil {
		return
	}
	return
}

// 发布评论，输入参数是评论数据，返回评论内容和报错信息
func (c *Comment) CommentAdd(ctx context.Context, comment *CommentStruct) (string, error) {
	return comment.Content, c.db.WithContext(ctx).Create(&comment).Error
}

// 删除评论之前修改视频评论数目
func (c *CommentStruct) BeforeDelete(tx *gorm.DB) (err error) {
	var comment *CommentStruct
	id := c.ID
	err = tx.Select("vid").Where("id=?", id).First(&comment).Error
	if err != nil {
		return
	}
	vid := comment.Vid
	err = tx.Model(&model.Video{}).Where("id=?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count-1")).Error
	if err != nil {
		return
	}
	return
}

// 删除评论，输入评论id，返回报错信息
func (c *Comment) CommentDel(ctx context.Context, id uint) error {
	return c.db.WithContext(ctx).Delete(&CommentStruct{
		Model: gorm.Model{
			ID: id,
		},
	}).Error
}

// 根据视频id查找评论
func (c *Comment) CommentListByVedio(ctx context.Context, vid uint) (comments []*CommentStruct, err error) {
	err = c.db.WithContext(ctx).Where("vid=?", vid).Find(&comments).Error
	return
}
