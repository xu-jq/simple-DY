/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 16:46:58
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
	Trans                  ut.Translator
	GlobalConfig           config.Config
	FeedSrvClient          pb.FeedClient
	PublishActionSrvClient pb.PublishActionClient
	PublishListSrvClient   pb.PublishListClient
	UserInfoSrvClient      pb.UserInfoClient
	UserLoginSrvClient     pb.UserLoginClient
	UserRegisterSrvClient  pb.UserRegisterClient
	SocialServiceClient    pb.SocialServiceClient
)
