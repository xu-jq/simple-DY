/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-03 22:40:48
 * @FilePath: /simple-DY/DY-api/video-web/main.go
 * @Description: 主程序
 */
package main

import (
	"log"
	"os"
	"os/signal"
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/initialize"
	"simple-DY/DY-api/video-web/utils/consul"
	"syscall"

	"net/http"
	_ "net/http/pprof"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func main() {

	debug := true // 线下环境为True，线上环境为False

	// pprof性能测试
	// go tool pprof -http=:7071 http://localhost:7070/debug/pprof/profile
	if debug {
		go func() {
			log.Println(http.ListenAndServe(":7070", nil))
		}()
	}

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
	serviceid := uuid.NewV4().String()
	err = registerClient.Register(global.GlobalConfig.MainServer.Address, global.GlobalConfig.MainServer.Port, "video-api", serviceid, []string{"api", "video"})
	defer registerClient.DeRegister(serviceid)
	if err != nil {
		zap.L().Error("Consul服务注册失败！错误信息：" + err.Error())
	}
	zap.L().Info("Consul服务注册成功！")

	go func() {
		// 运行主程序
		r.Run(":" + global.GlobalConfig.MainServer.Port)
	}()

	exitsignal := make(chan os.Signal, 1)
	signal.Notify(exitsignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitsignal
	zap.L().Info("video-api正在退出......")
}
