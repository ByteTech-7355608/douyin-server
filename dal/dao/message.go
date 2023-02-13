package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"

	"gorm.io/gorm"
)

type Message struct {
}

// 获得从a发往b的最新消息
func (m *Message) GetLastMessageByUid(ctx context.Context, uida, uidb int64) (msg model.Message, err error) {
	tx := db.WithContext(ctx).Model(model.Message{}).Where("uid = ? AND to_uid = ?", uida, uidb)
	if err = tx.Order("created_at desc").First(&msg).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 没找到a发往b的信息
			return msg, nil
		} else {
			Log.Errorf("get last message by uid err : %v, from_id %v, to_id %v", err, uida, uidb)
			return msg, err
		}
	}

	return msg, nil
}
