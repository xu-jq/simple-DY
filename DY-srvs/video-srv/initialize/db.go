/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 14:58:25
 * @FilePath: /simple-DY/DY-srvs/video-srv/initialize/db.go
 * @Description: 数据库初始化连接
 */
package initialize

import (
	"simple-DY/DY-srvs/video-srv/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	dsn := global.GlobalConfig.MySQL.UserName + ":" + global.GlobalConfig.MySQL.Password + "@tcp(" + global.GlobalConfig.MySQL.Address + ":" + global.GlobalConfig.MySQL.Port + ")/" + global.GlobalConfig.MySQL.DataBase + "?charset=utf8&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
