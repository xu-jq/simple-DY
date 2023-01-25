/*
 * @Date: 2023-01-24 21:52:37
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 15:18:20
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/oss/upload.go
 * @Description: 上传文件到七牛云
 */

package oss

import (
	"context"
	"simple-DY/DY-srvs/video-srv/global"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

func UploadFileToQiniuOSS(filePath string, name string) (err error) {
	putPolicy := storage.PutPolicy{
		Scope: global.GlobalConfig.OSS.Bucket,
	}
	mac := qbox.NewMac(global.GlobalConfig.OSS.AccessKey, global.GlobalConfig.OSS.SecretKey)

	ret := storage.PutRet{}

	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        &storage.ZoneHuadong, // 空间对应的机房
		UseHTTPS:      true,                 // 是否使用https域名
		UseCdnDomains: false,                // 上传是否使用CDN上传加速
	}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)

	err = formUploader.PutFile(context.Background(), &ret, upToken, name, filePath, &storage.PutExtra{})

	if err != nil {
		zap.L().Error("上传文件失败！错误信息：" + err.Error())
		return
	}
	zap.L().Info("上传文件成功！")
	return
}
