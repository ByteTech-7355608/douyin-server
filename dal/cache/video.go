package cache

import "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

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
