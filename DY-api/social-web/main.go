package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"simple-DY/DY-api/social-web/global"
	"simple-DY/DY-api/social-web/initialize"
	"simple-DY/DY-api/social-web/utils/register/consul"
	"syscall"
)

func main() {
	// 1. 初始化操作
	// 初始化配置文件
	initialize.InitCfg()
	// 初始化日志
	initialize.InitLogger()
	// 初始化路由
	Router := initialize.InitRouter()
	// 初始化微服务连接
	initialize.InitSrvConn()

	// 2. 服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	// 3. 服务启动
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}

	// 4. 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
