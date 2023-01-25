/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 13:28:49
 * @FilePath: /simple-DY/DY-srvs/video-srv/global/global.go
 * @Description: 全局变量
 */
package global

import (
	"simple-DY/DY-srvs/video-srv/config"
	"sync"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	GlobalConfig config.Config
	Wg           sync.WaitGroup
)
