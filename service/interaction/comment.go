package interaction

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

const (
	KAddType    = 1
	KDeleteType = 2
)

func (s *Service) CommentList(ctx context.Context, req *interaction.DouyinCommentListRequest) (resp *interaction.DouyinCommentListResponse, err error) {
	resp = interaction.NewDouyinCommentListResponse()
	res, err := s.dao.Comment.QueryCommentList(ctx, req.GetVideoId())
	if err != nil {
		Log.Errorf("quary comment list err: %v", err)
		return
	}
	commentList := make([]*model.Comment, 0)
	for _, v := range res {
		var user *model.User
		u, ok := s.dao.User.FindUserById(ctx, v.UID)
		if ok != nil {
			continue
		}
		user.Id = u.ID
		user.Name = u.Username
		comment := &model.Comment{
			Id:         v.ID,
			User:       user,
			Content:    v.Content,
			CreateDate: v.CreatedAt.String(),
		}
		commentList = append(commentList, comment)
	}
	resp.StatusCode = 200
	resp.CommentList = commentList
	return
}

func (s *Service) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	resp = interaction.NewDouyinCommentActionResponse()
	switch req.ActionType {
	case KAddType:
		comment, err := s.dao.Comment.AddComment(ctx, req)
		if err != nil {
			Log.Errorf("add comment err: %v", err)
			return nil, err
		}
		resp.Comment = &comment
	case KDeleteType:
		err = s.dao.Comment.DeleteComment(ctx, req)
		if err != nil {
			Log.Errorf("delete comment err: %v", err)
			return nil, err
		}
	}
	return
}
