package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/util"
	"bufio"
	"context"
	"os"
	"strconv"
	"time"
)

func (s *Service) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (r *base.DouyinPublishActionResponse, err error) {
	r = base.NewDouyinPublishActionResponse()
	filePath := "../../upload/"
	Name := strconv.FormatInt(time.Now().Unix(), 10)
	videoName := Name + "." + "mp4"
	file, err := os.OpenFile(filePath+videoName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		msg := "上传视频失败请重试"
		r.StatusMsg = &msg
		r.StatusCode = 500
		Log.Errorf("创建视频文件失败")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(req.Data)
	writer.Flush()
	play_url := "http://localhost:8888/upload/" + videoName
	picName, err := util.GetCoverPic(filePath+"3-8 章节小结.mp4", filePath+Name, 1)
	if err != nil {
		msg := "获取封面失败请重试"
		r.StatusMsg = &msg
		r.StatusCode = 500
		Log.Errorf("获取封面失败")
		return
	}
	conver_url := "http://localhost:8888/upload/" + picName
	uid := 1
	err = s.dao.Video.AddVideo(ctx, play_url, conver_url, req.Title, int64(uid))
	if err != nil {
		msg := "发布视频失败请重试"
		r.StatusMsg = &msg
		r.StatusCode = 500
		Log.Errorf("添加视频文件失败")
		return
	}
	msg := "发布视频成功"
	r.StatusMsg = &msg
	r.StatusCode = 200
	return
}
