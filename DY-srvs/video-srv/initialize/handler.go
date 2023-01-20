/*
 * @Date: 2023-01-20 19:05:40
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 19:15:43
 * @FilePath: /simple-DY/DY-srvs/video-srv/initialize/handler.go
 * @Description: 服务协程
 */
package initialize

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/handler"
)

func InitHandler() {
	global.Wg.Add(1)
	go handler.FeedService()
	global.Wg.Wait()
}
