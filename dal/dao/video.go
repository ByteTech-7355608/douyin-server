package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"

	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
)

type Video struct {
}

func (v *Video) GetPublishVideoListByUserId(ctx context.Context, uid int64) (videoList []model.Video, err error) {
	if err = db.WithContext(ctx).Model(model.Video{}).Omit("created_at, updated_at, deleted_at").Where("uid = ?", uid).Find(&videoList).Error; err != nil {
		Log.Errorf("search user published videos err: %v, user id: %d", err, uid)
		return nil, err
	}
	return
}

func (v *Video) QueryVideoByTime(ctx context.Context, latestTime int64) (videos []*model.Video, err error) {
	tx := db.WithContext(ctx).Model(model.Video{}).Where("unix_timestamp(created_at) <= ?", latestTime)
	if err = tx.Order("created_at desc").Limit(constants.VideoCountLimit).Find(&videos).Error; err != nil {
		Log.Errorf("query video by time err: %v, latestTime: %v", err, latestTime)
		return nil, err
	}
	return
}

func (v *Video) AddVideo(ctx context.Context, playUrl string, coverUrl string, title string, uid int64) (err error) {
	video := model.Video{
		Title:    title,
		PlayURL:  playUrl,
		CoverURL: coverUrl,
		UID:      uid,
	}
	if err = db.WithContext(ctx).Create(&video).Error; err != nil {
		Log.Errorf("add video err:%v", err)
	}
	return
}
