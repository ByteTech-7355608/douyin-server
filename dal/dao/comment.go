package dao

import (
	daoModel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	KitexModel "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

func (u *User) AddComment(ctx context.Context, req *interaction.DouyinCommentActionRequest) (commentRet KitexModel.Comment, err error) {
	comment := &daoModel.Comment{
		Vid:     req.VideoId,
		UID:     *req.BaseReq.UserId,
		Content: *req.CommentText,
	}
	user := &daoModel.User{
		ID: *req.BaseReq.UserId,
	}
	if err = u.db.WithContext(ctx).First(&user).Error; err != nil {
		Log.Errorf("miss user ID: %d", err, *req.BaseReq.UserId)
		return
	}
	if err = u.db.WithContext(ctx).Model(daoModel.Comment{}).Create(comment).Error; err != nil {
		Log.Errorf("add comment err: %v, comment: %+v", err, comment)
		return
	}
	userRet := KitexModel.User{
		Id:            *req.BaseReq.UserId,
		Name:          user.Username,
		FollowCount:   &user.FollowCount,
		FollowerCount: &user.FollowerCount,
	}
	commentRet = KitexModel.Comment{
		Id:         comment.ID,
		User:       &userRet,
		Content:    *req.CommentText,
		CreateDate: comment.CreatedAt.Format("01:02"),
	}
	return commentRet, nil
}

func (u *User) DeleteComment(ctx context.Context, req *interaction.DouyinCommentActionRequest) (err error) {
	if err = u.db.WithContext(ctx).Where("ID = ?", req.CommentId).Delete(&daoModel.Comment{}).Error; err != nil {
		Log.Errorf("delete comment err: %v, commentId: %d", err, *req.CommentId)
		return
	}
	return nil
}
