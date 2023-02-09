package util

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetCoverPic(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	// 这里把 snapshotPath 的 string 类型转换成 []string

	snapshotName = names[len(names)-1] + ".png"

	return snapshotName, nil
}
