/*
 * @Date: 2023-01-20 19:05:40
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:25:33
 * @FilePath: /simple-DY/DY-srvs/video-srv/initialize/handler.go
 * @Description: 初始化服务协程
 */
package initialize

import (
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/handler"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/consul"
	"simple-DY/DY-srvs/video-srv/utils/rabbitmq"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitHandler() {

	global.Wg.Add(2)

	// Video服务
	go func() {
		s := grpc.NewServer()
		pb.RegisterVideoServiceServer(s, &handler.Videoserver{})
		service(s, global.GlobalConfig.GRPC.Port, "video-srv")
	}()

	go rabbitmq.ConsumeSimple()

	global.Wg.Wait()
}

func service(s *grpc.Server, port string, name string) {
	defer global.Wg.Done()

	lis, err := net.Listen("tcp", global.GlobalConfig.Address.In+":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	zap.L().Info(name + "服务监听地址：" + lis.Addr().String())

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//服务注册
	register_client := consul.NewRegistryClient(global.GlobalConfig.Consul.Address, global.GlobalConfig.Consul.Port)
	serviceid := uuid.NewV4().String()
	register_client.Register(global.GlobalConfig.Address.Out, port, name, serviceid, []string{"srv", "video"})
	defer register_client.DeRegister(serviceid)

	go func() {
		if err := s.Serve(lis); err != nil {
			zap.L().Error("无法提供服务！错误信息：" + err.Error())
		}
	}()

	// 等待主程序的退出信号
	global.GRPCExitSignal.L.Lock()
	defer global.GRPCExitSignal.L.Unlock()
	global.GRPCExitSignal.Wait()

	// 停止服务
	s.GracefulStop()
	zap.L().Info(name + "退出成功！")
}
