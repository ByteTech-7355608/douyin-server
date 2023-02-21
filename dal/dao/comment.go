package dao

import (
	daoModel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	KitexModel "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/util"
	"context"

	"gorm.io/gorm"
)

type Comment struct {
}

func (c *Comment) QueryCommentList(ctx context.Context, vid int64) (res []*daoModel.Comment, err error) {
	res = make([]*daoModel.Comment, 0)
	err = db.WithContext(ctx).Model(daoModel.Comment{}).Where("vid = ?", vid).Find(&res).Error
	if err != nil {
		Log.Errorf("select comment list err: %v, videoId: %d", err, vid)
	}
	return
}

func (c *Comment) AddComment(ctx context.Context, req *interaction.DouyinCommentActionRequest) (commentRet *KitexModel.Comment, err error) {
	text := util.SensitiveMatch(*req.CommentText)
	comment := daoModel.Comment{
		Vid:     req.VideoId,
		UID:     *req.BaseReq.UserId,
		Content: text,
	}
	user := daoModel.User{}
	if err = db.WithContext(ctx).Select("username", "follow_count", "follower_count").Where("ID = ?", *req.BaseReq.UserId).Find(&user).Error; err != nil {
		Log.Errorf("miss user err: %v, ID: %d", err, *req.BaseReq.UserId)
		return
	}
	tx := db.Begin()
	if err = tx.WithContext(ctx).Model(daoModel.Comment{}).Create(&comment).Error; err != nil {
		Log.Errorf("add comment err: %v, comment: %+v", err, comment)
		tx.Rollback()
		return
	}
	if err = tx.WithContext(ctx).Model(daoModel.Video{}).Where("id = ?", req.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		Log.Errorf("incr comment_count err: %v, videoId: %d", err, req.VideoId)
		tx.Rollback()
		return
	}
	tx.Commit()
	userRet := KitexModel.User{
		Id:            *req.BaseReq.UserId,
		Name:          user.Username,
		FollowCount:   &user.FollowCount,
		FollowerCount: &user.FollowerCount,
	}
	commentRet = &KitexModel.Comment{
		Id:         comment.ID,
		User:       &userRet,
		Content:    text,
		CreateDate: comment.CreatedAt.Format("01:02"),
	}
	return commentRet, nil
}

func (c *Comment) DeleteComment(ctx context.Context, req *interaction.DouyinCommentActionRequest) (err error) {
	tx := db.Begin()
	if err = tx.WithContext(ctx).Where("ID = ?", req.CommentId).Delete(&daoModel.Comment{}).Error; err != nil {
		Log.Errorf("delete comment err: %v, commentId: %d", err, *req.CommentId)
		tx.Rollback()
		return
	}
	if err = tx.WithContext(ctx).Model(daoModel.Video{}).Where("id = ?", req.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", -1)).Error; err != nil {
		Log.Errorf("decr comment_count err: %v, videoId: %d", err, req.VideoId)
		tx.Rollback()
		return
	}
	tx.Commit()
	return nil
}
