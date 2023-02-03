package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"

	"context"

	"gorm.io/gorm"
)

type Relation struct {
}

func (r *Relation) IsUserFollowed(ctx context.Context, concernerID int64, concernedID int64) (isFollow bool, err error) {
	relation := model.Relation{}
	if err = db.WithContext(ctx).Model(model.Relation{}).Select("action").Where("concerner_id = ? AND concerned_id = ?", concernerID, concernedID).First(&relation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// relation 中不存在关注的关系，也可表示未关注
			Log.Infof("IsUserFollowed err：%v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
			return false, nil
		} else {
			// 数据库出错
			Log.Errorf("get follow relation err: %v, concerner_id: %d AND concerned_id：%d", err, concernerID, concernedID)
			return false, err
		}
	}
	return relation.Action, nil
}
