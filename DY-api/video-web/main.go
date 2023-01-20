/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 16:54:18
 * @FilePath: /simple-DY/DY-api/video-web/main.go
 * @Description: 主程序
 */
package main

import (
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/initialize"

	"go.uber.org/zap"
)

func main() {

	debug := true // 线下环境为True，线上环境为False

	// 初始化日志配置
	initialize.InitLogger()
	zap.L().Info("日志配置初始化成功！")

	// 初始化全局配置
	global.GlobalConfig = initialize.InitConfig(debug)
	zap.L().Info("全局配置初始化成功！")

	// 初始化路由
	r := initialize.Routers(debug)
	zap.L().Info("路由初始化成功！")

	// 运行主程序
	r.Run(":" + global.GlobalConfig.MainServerPort)
}
