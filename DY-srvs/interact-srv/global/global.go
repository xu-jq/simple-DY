package global

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"simple-DY/DY-srvs/interact-srv/config"
	"simple-DY/DY-srvs/interact-srv/proto"
	"time"
)

var (
	DB              *gorm.DB
	RDB             *redis.Client
	ServerConfig    config.ServerConfig
	NacosConfig     config.NacosConfig
	SocialSrvClient proto.SocialServiceClient
	VideoSrvClient  proto.VideoServiceClient
)

func init() {
	dsn := "dymysql:gxnw21XxRhY@tcp(121.37.98.68:3306)/simpledy?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
