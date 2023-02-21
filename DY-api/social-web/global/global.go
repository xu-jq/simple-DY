// Package global
/*

 */
package global

import (
	ut "github.com/go-playground/universal-translator"
	"simple-DY/DY-srvs/social-srv/config"
	"simple-DY/DY-srvs/social-srv/proto"
)

var (
	Trans           ut.Translator
	ServerConfig    *config.ServerConfig = &config.ServerConfig{}
	SocialSrvClient proto.SocialServiceClient
)
