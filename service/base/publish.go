package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/jwt"
	"ByteTech-7355608/douyin-server/util"
	"bufio"
	"context"
	"os"
	"strconv"
	"time"
)

func (s *Service) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (r *base.DouyinPublishActionResponse, err error) {
	r = base.NewDouyinPublishActionResponse()
	myclaim, err := jwt.ParseToken(req.Token)
	if err != nil {
		Log.Errorf("解析token失败")
		return
	}
	filePath := "../../upload/"
	Name := strconv.FormatInt(time.Now().Unix(), 10)
	videoName := Name + "." + "mp4"
	file, err := os.OpenFile(filePath+videoName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Log.Errorf("创建视频文件失败")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(req.Data)
	writer.Flush()
	play_url := "http://localhost:8888/upload/" + videoName
	picName, err := util.GetCoverPic(filePath+videoName, filePath+Name, 1)
	if err != nil {
		picName = "default.jpg"
		err = nil
	}
	conver_url := "http://localhost:8888/upload/" + picName
	err = s.dao.Video.AddVideo(ctx, play_url, conver_url, req.Title, myclaim.UserID)
	if err != nil {
		Log.Errorf("添加视频文件失败")
		return
	}
	r.StatusCode = 200
	msg := "视频投稿成功"
	r.StatusMsg = &msg
	return
}
