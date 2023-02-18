package util

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = constants.QiniuServerAccessKey
var ScretKey = constants.QiniuServerSecretKey
var Bucket = constants.QiniuServerBucket
var Url = constants.QiniuServerUrl

func UploadFile(file multipart.FileHeader) (string, string, error) {
	fileSize := file.Size
	name := file.Filename
	fileInfo, err := file.Open()
	if err != nil {
		return "", "", err
	}
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, ScretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, fileInfo, fileSize, &putExtra)

	if err != nil {
		return "", "", err
	}
	videoUrl := Url + ret.Key
	picReader, err := GetCoverPic(videoUrl, "../../", 1, name)
	var data []byte
	byte := data
	n, err := picReader.Read(byte)
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, picReader, int64(n), &putExtra)
	if err != nil {
		return videoUrl, "", err
	}
	picUrl := Url + ret.Key
	return videoUrl, picUrl, nil
}
