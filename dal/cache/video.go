package cache

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
)

type Video struct {
}

type VideoModel struct {
	Id            int64  `json:"id" redis:"id"`
	AuthorID      int64  `json:"author_id" redis:"author_id"`
	PlayUrl       string `json:"play_url" redis:"play_url"`
	CoverUrl      string `json:"cover_url" redis:"cover_url"`
	FavoriteCount int64  `json:"favorite_count" redis:"favorite_count"`
	CommentCount  int64  `json:"comment_count" redis:"comment_count"`
	Title         string `json:"title" redis:"title"`
}

func Video2VideoModel(video *model.Video) *VideoModel {
	return &VideoModel{
		Id:            video.Id,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	}
}

func (v *Video) IsExists(ctx context.Context, vids ...int64) int64 {
	keys := make([]string, len(vids))
	for i, vid := range vids {
		keys[i] = constants.GetVideoMsgKey(vid)
	}
	return Exists(ctx, keys...)
}

func (v *Video) SetVideoMessage(ctx context.Context, video *VideoModel) (ok bool) {
	return HSet(ctx, constants.GetVideoMsgKey(video.Id), video)
}

func (v *Video) IncrVideoField(ctx context.Context, vid int64, field string, incr int64) (ok bool) {
	return HIncr(ctx, constants.GetVideoMsgKey(vid), field, incr)
}

func (v *Video) GetVideoFields(ctx context.Context, vid int64, field ...string) []interface{} {
	return HMGet(ctx, constants.GetVideoMsgKey(vid), field...)
}
