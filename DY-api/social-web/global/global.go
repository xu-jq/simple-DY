// Package global
/*

 */
package global

import (
	ut "github.com/go-playground/universal-translator"
	"simple-DY/DY-api/social-web/config"
	"simple-DY/DY-api/social-web/proto"
)

var (
	Trans           ut.Translator
	ServerConfig    *config.ServerConfig = &config.ServerConfig{}
	SocialSrvClient proto.SocialServiceClient
)
