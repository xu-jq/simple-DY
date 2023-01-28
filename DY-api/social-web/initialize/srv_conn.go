package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 必须要导入
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"simple-DY/DY-api/social-web/global"
	"simple-DY/DY-api/social-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	socialConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.SocialSrvInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【社交服务失败】")
	}
	socialSrvClient := proto.NewSocialServiceClient(socialConn)
	global.SocialSrvClient = socialSrvClient
}
