/*
 * @Date: 2023-01-26 14:34:04
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 17:41:47
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/backup/backup.go
 * @Description: 备份文件的操作
 */
package backup

import (
	"os"
	"simple-DY/DY-srvs/video-srv/global"

	uuid "github.com/satori/go.uuid"
)

/**
 * @description: 根据id找到视频和图片文件应该存储的位置
 * @param {string} id
 * @return {string} filename 文件标题
 * @return {string} videoStaticFileName 视频文件路径
 * @return {string} imageStaticFileName 图片文件路径
 */
func GenerateFilePath(id string) (fileName, videoStaticFileName, imageStaticFileName string, err error) {
	// 用户文件夹路径
	userPath := global.GlobalConfig.StaticBackup.StaticPath + id
	// 用户视频与图片的路径
	videoStaticPath := userPath + global.GlobalConfig.StaticBackup.VideoPath
	imageStaticPath := userPath + global.GlobalConfig.StaticBackup.ImagePath

	_, err = os.Stat(userPath)
	if os.IsNotExist(err) {
		videoerr := os.MkdirAll(videoStaticPath, 0666)
		imageerr := os.Mkdir(imageStaticPath, 0666)
		if videoerr != nil || imageerr != nil {
			err = videoerr
			return
		}
	} else if err != nil {
		return
	}

	// 生成文件名称
	fileName = uuid.NewV4().String()

	// 组装完整的文件名称
	videoStaticFileName = videoStaticPath + fileName + global.GlobalConfig.StaticBackup.VideoSuffix
	imageStaticFileName = imageStaticPath + fileName + global.GlobalConfig.StaticBackup.ImageSuffix
	err = nil
	return
}
