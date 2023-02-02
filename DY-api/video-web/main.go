/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 19:41:36
 * @FilePath: /simple-DY/DY-api/video-web/main.go
 * @Description: 主程序
 */
package main

import (
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/initialize"
	"simple-DY/DY-api/video-web/utils/consul"

	"go.uber.org/zap"
)

func main() {

	debug := true // 线下环境为True，线上环境为False

	// 初始化日志配置
	initialize.InitLogger()
	zap.L().Info("日志配置初始化成功！")

	var err error

	// 初始化全局配置
	global.GlobalConfig, err = initialize.InitConfig(debug)
	if err != nil {
		zap.L().Error("配置读取失败！错误信息：" + err.Error())
		return
	}
	zap.L().Info("全局配置初始化成功！")

	// 初始化路由
	r := initialize.Routers(debug)
	zap.L().Info("路由初始化成功！")

	// GRPC翻译器（不知道怎么用）
	err = initialize.InitTrans("en")
	if err != nil {
		zap.L().Error("GRPC翻译器失败！错误信息：" + err.Error())
		return
	}
	zap.L().Info("GRPC翻译器初始化成功！")

	// 初始化GRPC连接
	initialize.InitSrvConn()

	// 服务注册
	registerClient := consul.NewRegistryClient(global.GlobalConfig.Consul.Address, global.GlobalConfig.Consul.Port)
	err = registerClient.Register(global.GlobalConfig.MainServer.Address, global.GlobalConfig.MainServer.Port, "video-api", "video-api")
	if err != nil {
		zap.L().Error("Consul服务注册失败！错误信息：" + err.Error())
	}
	zap.L().Info("Consul服务注册成功！")

	// 运行主程序
	r.Run(":" + global.GlobalConfig.MainServer.Port)
}
