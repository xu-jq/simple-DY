/*
 * @Date: 2023-01-23 14:50:51
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-23 15:16:03
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/ffmpeg/ExtractFirstFrame.go
 * @Description:
 */
package ffmpeg

import (
	"bytes"
	"fmt"
	"io"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		zap.L().Error("读取视频首帧失败！错误信息为：" + err.Error())
	}
	return buf
}

func ExtractFirstFrame(videoPath, imagePath string) {
	reader := ExampleReadFrameAsJpeg(videoPath, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		zap.L().Error("字节流转换为图片失败！错误信息为：" + err.Error())
	}
	err = imaging.Save(img, imagePath)
	if err != nil {
		zap.L().Error("保存图片失败！错误信息为：" + err.Error())
	}
}
