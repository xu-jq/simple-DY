/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 15:37:13
 * @FilePath: /simple-DY/DY-api/video-web/config/config.go
 * @Description: 配置文件结构体
 */
package config

type Config struct {
	MainServer MainServer // 主服务器
	GRPC       GRPC       // GRPC相关
	Consul     Consul     // Consul相关
	JWT        JWT        // JWT鉴权
	Time       Time       // 时间相关
}

// 主服务器
type MainServer struct {
	Address string // 地址
	Port    string // 端口
}

// GRPC相关
type GRPC struct {
	GRPCMsgSize GRPCMsgSize // GRPC消息传递大小限制
	GRPCTimeOut GRPCTimeOut // GRPC超时时间
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

// Consul
type Consul struct {
	Address string // 地址
	Port    string // 端口
}

// 鉴权
type JWT struct {
	Secret           string // JWT密钥
	TokenExpiresTime int64  // Token失效时间（秒）
}

// 时间相关
type Time struct {
	TimeFormat string // 时间格式化格式
}
