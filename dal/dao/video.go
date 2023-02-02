package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	db *gorm.DB
}

// 视频投稿，输入视频信息，报错则返回错误
func (v *Video) VideoAdd(ctx context.Context, video *model.Video) error {
	return v.db.WithContext(ctx).Create(&video).Error
}

// 删除视频，输入视频id，报错则返回错误
func (v *Video) VideoDel(ctx context.Context, id int) error {
	return v.db.WithContext(ctx).Delete(&Video{}, id).Error
}

// 查询某个时间点之前的视频，输入参数latestTime的格式为：年-月-日 时：分：秒，返回视频列表（30个），报错则返回错误，以及当前查到的时间点最后的时间
func (v *Video) VideoListByTime(ctx context.Context, latestTime string) (videos []*model.Video, err error, nextTime string) {
	var count int64
	format, _ := time.Parse("2006-01-02 15:04:05", latestTime)
	err = v.db.WithContext(ctx).Where("created_at<=?", format).Limit(30).Order("created_at Desc").Find(&videos).Count(&count).Error
	if err != nil || count == 0 {
		nextTime = time.Now().Format("2006-01-02 15:04:05")
		return
	}
	nextTime = videos[0].CreatedAt.Format("2006-01-02 15:04:05")
	return
}

// 查询用户发布的视频
func (v *Video) VideoListByUser(ctx context.Context, uid int) (videos []*model.Video, err error) {
	err = v.db.WithContext(ctx).Where("uid=?", uid).Order("created_at Desc").Find(&videos).Error
	return
}
