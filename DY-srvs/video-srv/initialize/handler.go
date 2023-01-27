/*
 * @Date: 2023-01-20 19:05:40
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 22:09:25
 * @FilePath: /simple-DY/DY-srvs/video-srv/initialize/handler.go
 * @Description: 服务协程
 */
package initialize

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/handler"
	"simple-DY/DY-srvs/video-srv/utils/rabbitmq"
)

func InitHandler() {
	global.Wg.Add(7)
	go handler.FeedService(global.GlobalConfig.GRPC.FeedPort)
	go handler.PublishActionService(global.GlobalConfig.GRPC.PublishActionPort)
	go handler.PublishListService(global.GlobalConfig.GRPC.PublishListPort)
	go handler.UserInfoService(global.GlobalConfig.GRPC.UserInfoPort)
	go handler.UserLoginService(global.GlobalConfig.GRPC.UserLoginPort)
	go handler.UserRegisterService(global.GlobalConfig.GRPC.UserRegisterPort)
	go rabbitmq.ConsumeSimple()
	global.Wg.Wait()
}
