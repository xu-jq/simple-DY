/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 18:08:57
 * @FilePath: /simple-DY/DY-srvs/video-srv/config/config.go
 * @Description: 配置文件结构体
 */
package config

type Config struct {
	MainServerAddress           string // 服务器ip地址
	MainServerPort              string // 启动端口号
	GRPCServerAddress           string // GRPC服务器地址
	GRPCServerFeedPort          string // GRPC服务端口号
	GRPCServerPublishActionPort string // GRPC服务端口号
	GRPCServerPublishListPort   string // GRPC服务端口号
	GRPCServerUserInfoPort      string // GRPC服务端口号
	GRPCServerUserLoginPort     string // GRPC服务端口号
	GRPCServerUserRegisterPort  string // GRPC服务端口号
	MySQLAddress                string // MySQL服务器地址
	MySQLPort                   string // MySQL端口号
	MySQLUserName               string // MySQL用户名
	MySQLPassword               string // MySQL密码
	MySQLDataBase               string // MySQL数据库
	NginxAddress                string // 静态资源服务器地址
	NginxPort                   string // 静态资源服务器端口号
	StaticPath                  string // 静态资源地址
	VideoPath                   string // 视频存放地址
	ImagePath                   string // 图片存放地址
	JwtSecret                   string // jwt密钥
	TimeString                  string // 时间格式化格式
	TokenExpiresTime            int64  // token过期时间
}
