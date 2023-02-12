package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

type Video struct {
}

func (v *Video) GetPublishVideoListByUserId(ctx context.Context, uid int64) (videoList []model.Video, err error) {
	if err = db.WithContext(ctx).Model(model.Video{}).Omit("created_at, updated_at, deleted_at").Where("uid = ?", uid).Find(&videoList).Error; err != nil {
		Log.Errorf("search user published videos err: %v, user id: %d", err, uid)
		return
	}
	return
}
