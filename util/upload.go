package util

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = constants.QiniuServerAccessKey
var ScretKey = constants.QiniuServerSecretKey
var Bucket = constants.QiniuServerBucket
var Url = constants.QiniuServerUrl

func UploadFile(path string) (string, string, error) {
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
	err := formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, path, &putExtra)

	if err != nil {
		return "", "", err
	}
	videoUrl := Url + ret.Key
	picReader, err := GetCoverPic(videoUrl, 1)
	if err != nil {
		return "", "", err
	}
	var data []byte
	byte := data
	n, err := picReader.Read(byte)
	if err != nil {
		return videoUrl, "", err
	}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, picReader, int64(n), &putExtra)
	if err != nil {
		return videoUrl, "", err
	}
	picUrl := Url + ret.Key
	return videoUrl, picUrl, nil
}
