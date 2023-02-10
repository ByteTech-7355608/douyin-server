package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"errors"

	"context"

	"gorm.io/gorm"
)

type Like struct {
}

func (l *Like) GetFavoriteVideoListByUserId(ctx context.Context, id int64) (videoList []model.Video, err error) {
	var userLikes []model.Like
	videoList = make([]model.Video, 0)
	// 根据 uid = id 和 action = 1 找到 对应的 vid
	if err = db.WithContext(ctx).Model(model.Like{}).Select("vid").Where("uid = ? AND action = ?", id, 1).Find(&userLikes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 该用户没有点赞过的视频，不影响返回结果
			Log.Infof("GetFavoriteVideoListByUserId err：%v", err)
			return videoList, nil
		} else {
			// 数据库出错
			Log.Errorf("find like err: %v, uid: %d", err, id)
			return nil, err
		}
	}
	for _, userLike := range userLikes {
		var video model.Video
		if err = db.WithContext(ctx).Model(model.Video{}).Omit("created_at, updated_at, deleted_at").Where("id = ?", userLike.Vid).First(&video).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 某一个点赞的视频未找到，不影响返回结果
				Log.Infof("%v, video id: %d", err, userLike.Vid)
				continue
			} else {
				// 数据库出错
				Log.Errorf("find video err: %v, video id: %d", err, userLike.Vid)
				return nil, err
			}
		}
		videoList = append(videoList, video)
	}
	return
}

// IsLike 用户是否点赞了视频
func (l *Like) IsLike(ctx context.Context, uid, vid int64) (like bool, err error) {
	var id int64
	// TODO 建立唯一索引，提高查询效率
	if err = db.WithContext(ctx).Model(model.Like{}).Select("id").Where("uid = ? AND vid = ?", uid, vid).Find(&id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		Log.Errorf("query relation between user %v and video %v err: %v", uid, vid, err)
		return false, err
	}
	return true, nil
}
