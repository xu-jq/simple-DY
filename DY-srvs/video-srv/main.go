/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-04 11:22:55
 * @FilePath: /simple-DY/DY-srvs/video-srv/main.go
 * @Description: 主程序
 */
package main

import (
	"log"
	"os"
	"os/signal"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/initialize"
	"sync"
	"syscall"
	"time"

	"net/http"
	_ "net/http/pprof"

	"go.uber.org/zap"
)

func main() {

	debug := true // 线下环境为True，线上环境为False

	// pprof性能测试
	// go tool pprof -http=:6061 http://localhost:6060/debug/pprof/profile
	if debug {
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
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

	// 初始化数据库
	global.DB, err = initialize.InitDb()
	if err != nil {
		zap.L().Error("连接数据库失败！错误信息：" + err.Error())
		return
	}
	zap.L().Info("数据库初始化成功！")

	go func() {
		// 初始化服务
		zap.L().Info("开始进行服务初始化！")
		initialize.InitHandler()
	}()

	// 初始化退出的条件变量
	global.GRPCExitSignal = sync.NewCond(&sync.Mutex{})

	// 接收主程序退出的信号
	exitsignal := make(chan os.Signal, 1)
	signal.Notify(exitsignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitsignal
	zap.L().Info("接收到退出信号......")

	// 通知其他进程退出
	global.GRPCExitSignal.Broadcast()
	zap.L().Info("video-srv正在退出......")

	// 等待退出
	time.Sleep(3 * time.Second)
}
