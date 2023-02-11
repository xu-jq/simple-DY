package initalize

import (
	"github.com/go-redis/redis"
	"simple-DY/DY-srvs/interact-srv/global"
)

func InitRedis() {
	RDB := global.RDB
	RDB = redis.NewClient(&redis.Options{
		Addr:     global.ServerConfig.RedisInfo.Host, // 指定
		Password: global.ServerConfig.RedisInfo.Password,
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
