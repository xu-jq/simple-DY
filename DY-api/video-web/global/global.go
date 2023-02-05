/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:11:35
 * @FilePath: /simple-DY/DY-api/video-web/global/global.go
 * @Description: 全局变量
 */
package global

import (
	"simple-DY/DY-api/video-web/config"
	pb "simple-DY/DY-api/video-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans                 ut.Translator
	GlobalConfig          config.Config
	VideoServiceClient    pb.VideoServiceClient
	SocialServiceClient   pb.SocialServiceClient
	InteractServiceClient pb.InteractServiceClient
)
