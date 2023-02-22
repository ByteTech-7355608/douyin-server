package interaction

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
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
		u, ok := s.dao.User.QueryUser(ctx, v.UID)
		if ok != nil {
			continue
		}
		user := &model.User{
			Id:     u.ID,
			Name:   u.Username,
			Avatar: &u.Avatar,
		}
		comment := &model.Comment{
			Id:         v.ID,
			User:       user,
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		commentList = append(commentList, comment)
	}
	resp.StatusCode = 200
	resp.CommentList = commentList
	return
}

func (s *Service) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	resp = interaction.NewDouyinCommentActionResponse()
	vid := req.GetVideoId()
	if s.cache.Video.IsExists(ctx, vid) == 0 {
		// 评论视频缓存不存在
		video, err := s.dao.Video.QueryVideoByID(ctx, vid)
		if err != nil {
			Log.Errorf("query video %v err: %v", vid, err)
			return resp, err
		}
		videoModel := &cache.VideoModel{
			Id:            video.ID,
			AuthorID:      video.UID,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
		}
		if !s.cache.Video.SetVideoMessage(ctx, videoModel) {
			Log.Errorf("set video message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	}

	switch req.ActionType {
	case KAddType:
		comment, err := s.dao.Comment.AddComment(ctx, req)
		if err != nil {
			Log.Errorf("add comment err: %v", err)
			return nil, err
		}
		resp.Comment = &comment
		s.cache.Video.IncrVideoField(ctx, vid, "comment_count", 1)
	case KDeleteType:
		err = s.dao.Comment.DeleteComment(ctx, req)
		if err != nil {
			Log.Errorf("delete comment err: %v", err)
			return nil, err
		}
		s.cache.Video.IncrVideoField(ctx, vid, "comment_count", -1)
	}
	return
}
