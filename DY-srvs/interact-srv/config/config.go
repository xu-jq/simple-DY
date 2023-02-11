package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type VideoSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type SocialSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Name          string          `mapstructure:"name" json:"name"`
	Tags          []string        `mapstructure:"tags" json:"tags"`
	Host          string          `mapstructure:"host" json:"host"`
	MysqlInfo     MysqlConfig     `mapstructure:"mysql" json:"mysql"`
	ConsulInfo    ConsulConfig    `mapstructure:"consul" json:"consul"`
	SocialSrvInfo SocialSrvConfig `mapstructure:"social_srv" json:"social_srv"`
	VideoSrvInfo  VideoSrvConfig  `mapstructure:"video_srv" json:"video_srv"`
	RedisInfo     RedisConfig     `mapstructure:"redis" json:"redis"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
