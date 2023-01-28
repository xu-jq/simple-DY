package global

import (
	"gorm.io/gorm"
	"simple-DY/DY-srvs/social-srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)
