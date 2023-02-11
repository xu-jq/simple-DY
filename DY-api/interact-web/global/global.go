package global

import (
	ut "github.com/go-playground/universal-translator"
	"simple-DY/DY-api/interact-web/config"
	"simple-DY/DY-api/interact-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	InteractSrvClient proto.InteractServiceClient

	NacosConfig config.NacosConfig
)
