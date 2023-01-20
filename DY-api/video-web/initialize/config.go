/*
 * @Date: 2023-01-19 15:53:00
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 17:47:13
 * @FilePath: /simple-DY/DY-api/video-web/initialize/config.go
 * @Description: 全局配置初始化
 */
package initialize

import (
	"simple-DY/DY-api/video-web/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(debug bool) config.Config {

	// 根据线上线下环境切换配置文件
	filename := "config-pro.yaml"
	if debug {
		filename = "config-debug.yaml"
	}
	zap.L().Info("配置文件为" + filename)

	viper.AddConfigPath("./")     //设置读取的文件路径
	viper.SetConfigName(filename) //设置读取的文件名
	viper.SetConfigType("yaml")   //设置文件的类型
	Config := &config.Config{}
	//尝试进行配置读取
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error("配置读取失败！错误信息：" + err.Error())
	} else {
		viper.Unmarshal(&Config)

		// 打印配置信息
		for _, key := range viper.AllKeys() {
			zap.L().Info(key + ": " + viper.GetString(key))
		}
	}
	return *Config
}
