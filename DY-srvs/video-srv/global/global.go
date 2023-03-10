/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-04 11:18:09
 * @FilePath: /simple-DY/DY-srvs/video-srv/global/global.go
 * @Description: ćšć±ćé
 */
package global

import (
	"simple-DY/DY-srvs/video-srv/config"
	"sync"

	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	GlobalConfig   config.Config
	Wg             sync.WaitGroup
	GRPCExitSignal *sync.Cond
)
