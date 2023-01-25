/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 15:45:43
 * @FilePath: /simple-DY/DY-api/video-web/config/config.go
 * @Description: 配置文件结构体
 */
package config

type Config struct {
	MainServer MainServer // 主服务器
	GRPC       GRPC       // GRPC相关
}

// 主服务器
type MainServer struct {
	Address string // 地址
	Port    string // 端口
}

// GRPC相关
type GRPC struct {
	Address           string      // GRPC服务地址
	FeedPort          string      // Feed服务端口号
	PublishActionPort string      // PublishAction服务端口号
	PublishListPort   string      // PublishList服务端口号
	UserInfoPort      string      // UserInfo服务端口号
	UserLoginPort     string      // UserLogin服务端口号
	UserRegisterPort  string      // UserRegister服务端口号
	GRPCMsgSize       GRPCMsgSize // GRPC消息传递大小限制
	GRPCTimeOut       GRPCTimeOut // GRPC超时时间
}

// GRPC消息传递大小限制
type GRPCMsgSize struct {
	CommonMB int // 普通消息（4MB）
	LargeMB  int // 文件字节流（32MB）
}

// GRPC超时时间
type GRPCTimeOut struct {
	CommonSecond int // 普通超时时间（1秒）
	FileSecond   int // 文件超时时间（10秒）
}
