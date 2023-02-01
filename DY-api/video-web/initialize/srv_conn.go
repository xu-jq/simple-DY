/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-01 22:44:45
 * @FilePath: /simple-DY/DY-api/video-web/initialize/srv_conn.go
 * @Description: 初始化客户端GRPC连接
 */
package initialize

import (
	"simple-DY/DY-api/video-web/global"
	pb "simple-DY/DY-api/video-web/proto"
	"simple-DY/DY-api/video-web/utils/consul"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn
var err error

// 封装重复部分为一个辅助函数
func initConn(name string) (conn *grpc.ClientConn, err error) {
	size := global.GlobalConfig.GRPC.GRPCMsgSize.CommonMB
	// 上传文件需要将传输限制开大
	if name == "PublishAction" {
		size = global.GlobalConfig.GRPC.GRPCMsgSize.LargeMB
	}
	conn, err = grpc.Dial(
		"consul://"+global.GlobalConfig.Consul.Address+":"+global.GlobalConfig.Consul.Port+"/"+name+"?wait=14s",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*size),
			grpc.MaxCallSendMsgSize(1024*1024*size),
		),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	return
}

// 初始化GRPC连接，注册到Consul中
func InitSrvConn() {
	conn, err = initConn("Feed")
	if err != nil {
		zap.L().Error("Feed初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("Feed初始化连接成功！")
	global.FeedSrvClient = pb.NewFeedClient(conn)

	conn, err = initConn("PublishAction")
	if err != nil {
		zap.L().Error("PublishAction初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("PublishAction初始化连接成功！")
	global.PublishActionSrvClient = pb.NewPublishActionClient(conn)

	conn, err = initConn("PublishList")
	if err != nil {
		zap.L().Error("PublishList初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("PublishList初始化连接成功！")
	global.PublishListSrvClient = pb.NewPublishListClient(conn)

	conn, err = initConn("UserInfo")
	if err != nil {
		zap.L().Error("UserInfo初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("UserInfo初始化连接成功！")
	global.UserInfoSrvClient = pb.NewUserInfoClient(conn)

	conn, err = initConn("UserLogin")
	if err != nil {
		zap.L().Error("UserLogin初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("UserLogin初始化连接成功！")
	global.UserLoginSrvClient = pb.NewUserLoginClient(conn)

	conn, err = initConn("UserRegister")
	if err != nil {
		zap.L().Error("UserRegister初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("UserRegister初始化连接成功！")
	global.UserRegisterSrvClient = pb.NewUserRegisterClient(conn)

	// 服务注册
	registerClient := consul.NewRegistryClient(global.GlobalConfig.Consul.Address, global.GlobalConfig.Consul.Port)
	err = registerClient.Register(global.GlobalConfig.MainServer.Address, global.GlobalConfig.MainServer.Port, "video-api", "video-api")
	if err != nil {
		zap.L().Error("Consul服务注册失败！错误信息：" + err.Error())
	}
	zap.L().Info("Consul服务注册成功！")
}
