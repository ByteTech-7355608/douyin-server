package dao

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

type Message struct {
}

func (m *Message) QueryMessageList(ctx context.Context, fid, tid int64) (messageList []*model.Message, err error) {
	tx := db.WithContext(ctx).Model(model.Message{})
	tx.Where("uid = ? AND to_uid = ?", fid, tid)
	tx.Or("uid = ? AND to_uid = ?", tid, fid)
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
