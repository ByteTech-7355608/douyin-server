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

func (m *Message) QueryMessageList(ctx context.Context, fid, tid, preMsgTime int64) (messageList []*model.Message, err error) {
	tx := db.WithContext(ctx).Model(model.Message{})
	// UNIX_TIMESTAMP()按照0时区计算，数据库存的是东八区，减去差值
	tx.Where("(UNIX_TIMESTAMP(created_at) - 28800) * 1000 > ?", preMsgTime)
	tx.Where("(uid = ? AND to_uid = ?) OR (uid = ? AND to_uid = ?)", fid, tid, tid, fid)
	tx.Order("created_at")
	if err = tx.Find(&messageList).Error; err != nil {
		Log.Errorf("query message list err: %v, fid: %v, tid: %v", err, fid, tid)
	}
	return
}

func (m *Message) CreateRecord(ctx context.Context, message *model.Message) (err error) {
	if err = db.WithContext(ctx).Create(message).Error; err != nil {
		Log.Errorf("insert message err: %v", err)
	}
	return
}
