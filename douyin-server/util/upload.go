package util

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func UploadVideo(file *multipart.FileHeader, c *app.RequestContext) (videoUrl, picUrl string, err error) {
	unix_str := strconv.FormatInt(time.Now().Unix(), 10)
	suffixPath := "../../upload/"
	videoName := unix_str + file.Filename
	file_path := suffixPath + videoName
	err = c.SaveUploadedFile(file, file_path)
	if err != nil {
		return
	}
	picName, err := GetCoverPic(file_path, suffixPath+unix_str, 1, unix_str)
	if err != nil {
		return
	}
	videoUrl = constants.UploadAddr + videoName
	picUrl = constants.UploadAddr + picName
	return
}
