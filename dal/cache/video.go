package cache

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	dbmodel "ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"strconv"
)

type Video struct {
	dao *dao.Dao
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

func DBVideo2VideoModel(video *dbmodel.Video) *VideoModel {
	return &VideoModel{
		Id:            video.ID,
		AuthorID:      video.UID,
		PlayUrl:       video.PlayURL,
		CoverUrl:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	}
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

func VideoModel2DBVideo(videoModel *VideoModel) *dbmodel.Video {
	return &dbmodel.Video{
		ID:            videoModel.Id,
		UID:           videoModel.AuthorID,
		PlayURL:       videoModel.PlayUrl,
		CoverURL:      videoModel.CoverUrl,
		FavoriteCount: videoModel.FavoriteCount,
		CommentCount:  videoModel.CommentCount,
		Title:         videoModel.Title,
	}
}

func Map2VideoModel(mp map[string]string) (videoModel *VideoModel, err error) {
	var id, authorID, favoriteCount, commentCount int64
	id, err = strconv.ParseInt(mp["id"], 10, 64)                        //nolint: staticcheck
	authorID, err = strconv.ParseInt(mp["author_id"], 10, 64)           //nolint: staticcheck
	favoriteCount, err = strconv.ParseInt(mp["favorite_count"], 10, 64) //nolint: staticcheck
	commentCount, err = strconv.ParseInt(mp["comment_count"], 10, 64)   //nolint: staticcheck
	if err != nil {
		Log.Errorf("parse int from map err: %v", err)
		return nil, err
	}

	return &VideoModel{
		Id:            id,
		AuthorID:      authorID,
		PlayUrl:       mp["play_url"],
		CoverUrl:      mp["cover_url"],
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		Title:         mp["title"],
	}, nil
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

func (v *Video) GetVideoMessage(ctx context.Context, vid int64) (video *VideoModel, err error) {
	video, err = Map2VideoModel(HGetAll(ctx, constants.GetVideoMsgKey(vid)))
	if err != nil {
		Log.Warnf("get video %d message err: %v", vid, err)
		return nil, err
	}
	return
}
