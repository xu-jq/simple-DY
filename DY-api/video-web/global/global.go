/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-03 10:29:15
 * @FilePath: /simple-DY/DY-api/video-web/global/global.go
 * @Description: 全局变量
 */
package global

import (
	"simple-DY/DY-api/video-web/config"
	socialpb "simple-DY/DY-api/video-web/proto/social"
	videopb "simple-DY/DY-api/video-web/proto/video"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans                  ut.Translator
	GlobalConfig           config.Config
	FeedSrvClient          videopb.FeedClient
	PublishActionSrvClient videopb.PublishActionClient
	PublishListSrvClient   videopb.PublishListClient
	UserInfoSrvClient      videopb.UserInfoClient
	UserLoginSrvClient     videopb.UserLoginClient
	UserRegisterSrvClient  videopb.UserRegisterClient
	SocialServiceClient    socialpb.SocialServiceClient
)
