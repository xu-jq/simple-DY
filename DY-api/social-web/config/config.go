// Package config
/*

 */
package config

type SocialSrvInfo struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name" json:"name"`
	Host          string        `mapstructure:"host" json:"host"`
	Tags          []string      `mapstructure:"tags" json:"tags"`
	Port          int           `mapstructure:"port" json:"port"`
	SocialSrvInfo SocialSrvInfo `mapstructure:"social_srv" json:"social_srv"`
	JWTInfo       JWTConfig     `mapstructure:"jwt" json:"jwt"`
	RedisInfo     RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo    ConsulConfig  `mapstructure:"consul" json:"consul"`
}
