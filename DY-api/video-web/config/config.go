/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 16:12:13
 * @FilePath: /simple-DY/DY-api/video-web/config/config.go
 * @Description: 配置文件结构体
 */

package config

type Config struct {
	MainServerAddress string // 服务器ip地址
	MainServerPort    string // 启动端口号
	NginxPort         string // 静态资源服务器端口号
	StaticPath        string // 静态资源地址
	VideoPath         string // 视频存放地址
	ImagePath         string // 图片存放地址
}
