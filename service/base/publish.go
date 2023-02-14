package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"bufio"
	"context"
	"os"
	"strconv"
	"time"
)

func (s *Service) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (resp *base.DouyinPublishListResponse, err error) {
	resp = base.NewDouyinPublishListResponse()

	userInstance, err := s.dao.User.FindUserById(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get user err: %v", err)
		return
	}

	videoList, err := s.dao.Video.GetPublishVideoListByUserId(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get publish list err : %v", err)
		return
	}

	//user 类型转换
	user := &model.User{
		Id:            userInstance.ID,
		Name:          userInstance.Username,
		FollowCount:   &userInstance.FollowCount,
		FollowerCount: &userInstance.FollowerCount,
	}

	//video 类型转换
	var videos []*model.Video
	for _, videoInstance := range videoList {
		video := &model.Video{
			Id:            videoInstance.ID,
			PlayUrl:       videoInstance.PlayURL,
			CoverUrl:      videoInstance.CoverURL,
			FavoriteCount: videoInstance.FavoriteCount,
			CommentCount:  videoInstance.CommentCount,
			Title:         videoInstance.Title,
			Author:        user,
		}
		videos = append(videos, video)
	}
	resp.SetVideoList(videos)
	return
}

func (s *Service) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (r *base.DouyinPublishActionResponse, err error) {
	r = base.NewDouyinPublishActionResponse()
	user_id := *req.BaseReq.UserId
	filePath := "../../upload/"
	addr := constants.UploadAddr
	Name := strconv.FormatInt(time.Now().Unix(), 10)
	videoName := Name + "." + "mp4"
	file, err := os.OpenFile(filePath+videoName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Log.Errorf("创建视频文件失败")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.Write(req.Data)
	if err != nil {
		Log.Errorf("视频文件写入错误")
		return
	}
	writer.Flush()
	play_url := addr + videoName
	picName, err := util.GetCoverPic(filePath+videoName, filePath+Name, 1)
	if err != nil {
		picName = "default.jpg"
		err = nil
	}
	conver_url := addr + picName
	err = s.dao.Video.AddVideo(ctx, play_url, conver_url, req.Title, user_id)
	if err != nil {
		Log.Errorf("添加视频文件失败")
		return
	}
	r.StatusCode = 200
	return
}
