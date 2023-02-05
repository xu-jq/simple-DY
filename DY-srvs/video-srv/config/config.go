/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:26:17
 * @FilePath: /simple-DY/DY-srvs/video-srv/config/config.go
 * @Description: 配置文件结构体
 */
package config

type Config struct {
	Address      Address      // 地址
	GRPC         GRPC         // GRPC相关
	MySQL        MySQL        // 数据库
	OSS          OSS          // 静态资源存储
	StaticBackup StaticBackup // 静态资源备份
	JWT          JWT          // 鉴权
	Time         Time         // 时间相关
	RabbitMQ     RabbitMQ     // 消息队列
	Consul       Consul       // Consul
}

// 地址相关
type Address struct {
	Out string // 本机外网地址
	In  string // 本机内网地址
}

// GRPC相关
type GRPC struct {
	Address     string // GRPC服务地址
	Port        string
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

// 数据库
type MySQL struct {
	Address  string // MySQL地址
	Port     string // MySQL端口
	UserName string // MySQL用户名
	Password string // MySQL密码
	DataBase string // MySQL数据库

}

// 静态资源存储
type OSS struct {
	Address     string // 外链地址
	VideoPath   string // 视频路径
	VideoSuffix string // 视频后缀
	ImagePath   string // 图片路径
	ImageSuffix string // 图片后缀
	AccessKey   string // AccessKey
	SecretKey   string // SecretKey
	Region      string // OSS地区
	Bucket      string // Bucket
}

// 静态资源备份
type StaticBackup struct {
	StaticPath  string // 备份路径
	VideoPath   string // 视频路径
	VideoSuffix string // 视频后缀
	ImagePath   string // 图片路径
	ImageSuffix string // 图片后缀
}

// 鉴权
type JWT struct {
	Secret           string // JWT密钥
	TokenExpiresTime int64  // Token失效时间（秒）
}

// 时间相关
type Time struct {
	TimeFormat string // 时间格式化的格式
}

// 消息队列
type RabbitMQ struct {
	Address     string // RabbitMQ地址
	Port        string // RabbitMQ端口
	UserName    string // RabbitMQ用户名
	Password    string // RabbitMQ密码
	VirtualHost string // RabbitMQ VirtualHost
}

// Consul
type Consul struct {
	Address string // 地址
	Port    string // 端口
}
