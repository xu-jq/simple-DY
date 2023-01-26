/*
 * @Date: 2023-01-25 22:47:19
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 17:37:38
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/dao/videodao.go
 * @Description: videos表操作
 */
package dao

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	"strconv"
	"time"

	"go.uber.org/zap"
)

/**
 * @description: 插入数据到Videos表中
 * @param {int64} authorid
 * @param {string} filename
 * @param {int64} publishtime
 * @param {string} title
 * @return {*}
 */
func InsertVideo(authorid int64, filename string, publishtime int64, title string) {
	// 构建插入的结构体
	videoInfo := models.Videos{
		AuthorId:    authorid,
		FileName:    filename,
		PublishTime: publishtime,
		Title:       title,
	}
	// 插入到Videos表中
	global.DB.Create(&videoInfo)
}

/**
 * @description: 根据用户的id获取全部的投稿视频
 * @param {int64} id
 * @return {*}
 */
func GetAuthorVideos(id int64) []map[string]interface{} {
	result := []map[string]interface{}{}

	// 查询作者的视频
	global.DB.Model(&models.Videos{}).Where("author_id = " + strconv.FormatInt(id, 10)).Order("publish_time DESC").Find(&result)

	return result
}

/**
 * @description: 根据时间和数量获取视频流
 * @param {int64} inputTime
 * @param {int} num
 * @return {*}
 */
func GetFeedVideos(inputTime int64, num int) (result []map[string]interface{}, latestTimeStamp map[string]interface{}) {
	// 获取请求的时间戳
	timeStamp := inputTime / 1000
	zap.L().Info("此请求的时间是：" + time.Unix(timeStamp, 0).Format(global.GlobalConfig.Time.TimeFormat))

	// 查询前30个视频
	videoQuery := global.DB.Model(&models.Videos{}).Where("publish_time < " + strconv.FormatInt(timeStamp, 10)).Order("publish_time DESC").Limit(num).Find(&result)

	latestTimeStamp = make(map[string]interface{})
	// 查询前30个视频的最早时间
	global.DB.Table("(?) as u", videoQuery).Select("publish_time as t").Order("publish_time ASC").Limit(1).Find(&latestTimeStamp)

	// 数据库中没有更早的视频，就直接使用当前的时间戳替换
	if len(latestTimeStamp) == 0 {
		latestTimeStamp["t"] = time.Now().Unix()
	}
	return
}
