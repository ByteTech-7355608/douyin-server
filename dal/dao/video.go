package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

type Video struct {
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
		return
	}
	return
}
