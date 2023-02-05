/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:17:23
 * @FilePath: /simple-DY/DY-api/video-web/initialize/srv_conn.go
 * @Description: 初始化客户端GRPC连接
 */
package initialize

import (
	"simple-DY/DY-api/video-web/global"
	pb "simple-DY/DY-api/video-web/proto"

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
	if name == "video-srv" {
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
	conn, err = initConn("video-srv")
	if err != nil {
		zap.L().Error("video-srv初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("video-srv初始化连接成功！")
	global.VideoServiceClient = pb.NewVideoServiceClient(conn)

	// wang hui的服务初始化

	conn, err = initConn("social-srv")
	if err != nil {
		zap.L().Error("social-srv初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("social-srv初始化连接成功！")
	global.SocialServiceClient = pb.NewSocialServiceClient(conn)

	// xu junqi的服务初始化

	conn, err = initConn("interact-srv")
	if err != nil {
		zap.L().Error("interact-srv初始化连接失败！错误信息：" + err.Error())
	}
	zap.L().Info("interact-srv初始化连接成功！")
	global.InteractServiceClient = pb.NewInteractServiceClient(conn)
}
