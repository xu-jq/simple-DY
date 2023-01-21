/*
 * @Date: 2023-01-20 19:05:40
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-21 12:22:09
 * @FilePath: /simple-DY/DY-srvs/video-srv/initialize/handler.go
 * @Description: 服务协程
 */
package initialize

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/handler"
)

func InitHandler() {
	global.Wg.Add(6)
	go handler.FeedService(global.GlobalConfig.GRPCServerFeedPort)
	go handler.PublishActionService(global.GlobalConfig.GRPCServerPublishActionPort)
	go handler.PublishListService(global.GlobalConfig.GRPCServerPublishListPort)
	go handler.UserInfoService(global.GlobalConfig.GRPCServerUserInfoPort)
	go handler.UserLoginService(global.GlobalConfig.GRPCServerUserLoginPort)
	go handler.UserRegisterService(global.GlobalConfig.GRPCServerUserRegisterPort)
	global.Wg.Wait()
}
