/*
 * @Date: 2023-01-19 15:53:00
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 15:16:21
 * @FilePath: /simple-DY/DY-api/video-web/initialize/config.go
 * @Description: 全局配置初始化
 */
package initialize

import (
	"simple-DY/DY-api/video-web/config"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(debug bool) (config.Config, error) {

	// 根据线上线下环境切换配置文件
	filename := "config-pro.yaml"
	if debug {
		filename = "config-debug.yaml"
	}
	zap.L().Info("配置文件：" + filename)

	// viper配置
	viper.AddConfigPath("./")     //设置读取的文件路径
	viper.SetConfigName(filename) //设置读取的文件名
	viper.SetConfigType("yaml")   //设置文件的类型
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Info("配置文件更改！" + e.Name)
		// 打印配置信息
		LogConfigOutput(viper.AllKeys())
	})
	viper.WatchConfig()
	Config := &config.Config{}
	//尝试进行配置读取
	if err := viper.ReadInConfig(); err != nil {
		return *Config, err
	}

	// 将配置写入结构体
	viper.Unmarshal(&Config)

	// 打印配置信息
	LogConfigOutput(viper.AllKeys())

	return *Config, nil
}

// 在日志中输出当前的配置信息
func LogConfigOutput(allkey []string) {
	for _, key := range allkey {
		zap.L().Info(key + ": " + viper.GetString(key))
	}
}
