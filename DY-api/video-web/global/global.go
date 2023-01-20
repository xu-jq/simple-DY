/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 16:53:26
 * @FilePath: /simple-DY/DY-api/video-web/global/global.go
 * @Description: 全局变量
 */
package global

import (
	"simple-DY/DY-api/video-web/config"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator
	GlobalConfig config.Config
)
