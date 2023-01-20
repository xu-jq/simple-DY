/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 19:25:38
 * @FilePath: /simple-DY/DY-srvs/video-srv/initialize/db.go
 * @Description: 数据库初始化连接
 */
package initialize

import (
	"simple-DY/DY-srvs/video-srv/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() {
	dsn := global.GlobalConfig.MySQLUserName + ":" + global.GlobalConfig.MySQLPassword + "@tcp(" + global.GlobalConfig.MySQLAddress + ":" + global.GlobalConfig.MySQLPort + ")/" + global.GlobalConfig.MySQLDataBase + "?charset=utf8&parseTime=True&loc=Local"
	// dsn := "dymysql:gxnw21XxRhY@tcp(121.37.98.68:3306)/simpledy?charset=utf8&parseTime=True&loc=Local"
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("连接数据库失败！错误信息为：" + err.Error())
	}
}
