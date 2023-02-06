// Package initialize /**
package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"simple-DY/DY-api/social-web/global"
)

// GetEnvInfo 获取当前系统变量
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)

}

// InitCfg 初始化配置文件，加载配置文件的内容到全局变量中
func InitCfg() {
	// 获取"TIKTOK_DEBUG"判断线上线下环境，使用不同的配置文件。
	debug := GetEnvInfo("TIKTOK_DEBUG")
	cfgPrefix := "config"
	cfgFileName := fmt.Sprintf("%s-pro.yaml", cfgPrefix)
	if debug {
		cfgFileName = fmt.Sprintf("%s-debug.yaml", cfgPrefix)
	}
	// 通过viper读取配置
	// 1.New()
	v := viper.New()
	v.SetConfigFile(cfgFileName)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 2.unmarshal到全局变量中
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：%v", global.ServerConfig)

	// 3.viper动态监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
	})
}
