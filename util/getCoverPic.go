package util

import (
	"bytes"
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

//func GetCoverPic(videoPath, snapshotPath string, frameNum int, name string) (snapshotName string, err error) {
//	bytes.NewReader()
//	buf := bytes.NewBuffer(nil)
//	err = ffmpeg.Input(videoPath).
//		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
//		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
//		WithOutput(buf, os.Stdout).
//		Run()
//	if err != nil {
//		panic(err)
//		return "", err
//	}
//	img, err := imaging.Decode(buf)
//	if err != nil {
//		panic(err)
//		return "", err
//	}
//
//	err = imaging.Save(img, snapshotPath+".png")
//	if err != nil {
//		panic(err)
//		return "", err
//	}
//
//	snapshotName = name + ".png"
//
//	return snapshotName, nil
//}

func GetCoverPic(videoPath, snapshotPath string, frameNum int, name string) (io.Reader, error) {
	bytes.NewReader(nil)
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
		return nil, err
	}

	return buf, nil
}
