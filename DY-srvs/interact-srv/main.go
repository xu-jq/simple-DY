package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"simple-DY/DY-srvs/interact-srv/global"
	"simple-DY/DY-srvs/interact-srv/initalize"
	"simple-DY/DY-srvs/interact-srv/proto"
	"simple-DY/DY-srvs/interact-srv/utils"
	"simple-DY/DY-srvs/interact-srv/utils/register/consul"
	"syscall"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 8888, "端口号")

	//初始化
	initalize.InitLogger()
	initalize.InitConfig()
	initalize.InitDB()
	initalize.InitRedis()
	initalize.InitSocialSrvConn()
	initalize.InitVideoSrvConn()

	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port: ", *Port)

	server := grpc.NewServer()
	proto.RegisterInteractServiceServer(server, &proto.UnimplementedInteractServiceServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//启动服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = register_client.Register(global.ServerConfig.Host, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	if err != nil {
		panic(err)
	}
	zap.S().Debugf("启动服务器，端口：%d", *Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = register_client.DeRegister(serviceId)
	if err != nil {
		zap.S().Info("注销失败：", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
