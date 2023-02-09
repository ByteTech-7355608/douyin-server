package dao

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

type Video struct {
}

func (v *Video) GetPublishVideoListByUserId(ctx context.Context, uid int64) (videlist []model.Video, err error) {
	if err = db.WithContext(ctx).Model(model.Video{}).Where("uid = ?", uid).Find(&videlist).Error; err != nil {
		Log.Errorf("search user published videos err: %v, user id: %d", err, uid)
		return
	}
	return
}
