/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 19:24:58
 * @FilePath: /simple-DY/DY-srvs/video-srv/main.go
 * @Description: 主程序
 */
package main

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/initialize"

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

	// 初始化数据库
	initialize.InitDb()
	zap.L().Info("数据库初始化成功！")

	// 初始化服务
	zap.L().Info("开始进行服务初始化！")
	initialize.InitHandler()
}
