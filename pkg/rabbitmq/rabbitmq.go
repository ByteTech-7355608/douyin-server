package rabbitmq

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMsg struct {
	Uid      int64  `json:"uid"`
	Title    string `json:"title"`
	VideoUrl string `json:"VideoUrl"`
}

const Username = constants.RabbitUsername
const Password = constants.RabbitPassword
const Address = constants.RabbitAddress

func Produce(uid int64, title, url string) error {
	dsn := fmt.Sprintf("amqp://%v:%v@%v/", Username, Password, Address)
	conn, err := amqp.Dial(dsn)
	defer conn.Close()
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	msg := RabbitMsg{
		Uid:      uid,
		Title:    title,
		VideoUrl: url,
	}
	byte_data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = ch.ExchangeDeclare(
		"douyin_exchange",
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"douyin_exchange",
		"upload",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         byte_data,
		})
	if err != nil {
		return err
	}
	return nil
}

func Consume(ctx context.Context) {
	dsn := fmt.Sprintf("amqp://%v:%v@%v/", Username, Password, Address)
	conn, err := amqp.Dial(dsn)
	defer conn.Close()
	if err != nil {
		Log.Errorf("create connection,err:%v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		Log.Errorf("create channel,err:%v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"douyin_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		Log.Errorf("create queue,err:%v", err)
	}
	err = ch.QueueBind(
		queue.Name,
		"upload",
		"douyin_exchange",
		false,
		nil,
	)
	if err != nil {
		Log.Errorf("create bind,err:%v", err)
	}
	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	db := dao.NewDao()
	for msg := range msgs {
		msg.Ack(true)
		go func(msg amqp.Delivery) {
			byte_data := msg.Body
			var publishInfo RabbitMsg
			err := json.Unmarshal(byte_data, &publishInfo)
			if err != nil {
				Log.Errorf("unmarshal publish request fail,err:%v", err)
			}
			videoUrl, pictureUrl, err := util.UploadFile(publishInfo.VideoUrl)
			if err != nil {
				Produce(publishInfo.Uid, publishInfo.Title, publishInfo.VideoUrl)
				Log.Errorf("upload file err:%v", err)
			}
			err = db.Video.AddVideo(ctx, videoUrl, pictureUrl, publishInfo.Title, publishInfo.Uid)
			if err != nil {
				Produce(publishInfo.Uid, publishInfo.Title, publishInfo.VideoUrl)
				Log.Errorf("publish action insert db err : %v", err)
			}
			err = os.Remove(publishInfo.VideoUrl)
			if err != nil {
				Log.Errorf("delete video file err:%v", err)
			}
		}(msg)
	}
}
