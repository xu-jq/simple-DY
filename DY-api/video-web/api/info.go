/*
 * @Date: 2023-02-02 18:44:40
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 19:22:44
 * @FilePath: /simple-DY/DY-api/video-web/api/info.go
 * @Description: 工具api
 */
package api

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func userService(c *gin.Context, idString string) (id, followCount, followerCount int64, name, statusMsg string, statusCode int32, isFollow bool) {

	var wgUser sync.WaitGroup
	// 并行查询服务
	wgUser.Add(3)
	// douyinUser
	go func() {
		defer wgUser.Done()
		responseUserInfo, err := douyinUser(idString)
		if err != nil {
			zap.L().Error("douyinUser GRPC失败！错误信息：" + err.Error())
			return
		}
		id = responseUserInfo.User.GetId()
		name = responseUserInfo.User.GetName()
		statusCode = responseUserInfo.GetStatusCode()
		statusMsg = responseUserInfo.GetStatusMsg()
	}()
	// douyinGetFollowList
	go func() {
		defer wgUser.Done()
		responseGetFollowList, err := douyinGetFollowList(idString)
		if err != nil {
			zap.L().Error("douyinGetFollowList GRPC失败！错误信息：" + err.Error())
			return
		}
		followCount = int64(len(responseGetFollowList.UserList))
	}()
	// douyinFollowerList
	go func() {
		defer wgUser.Done()
		responseFollowerList, err := douyinFollowerList(idString)
		if err != nil {
			zap.L().Error("douyinFollowerList GRPC失败！错误信息：" + err.Error())
			return
		}
		followerCount = int64(len(responseFollowerList.UserList))

		// 查询token是否关注了user，要去user的粉丝列表里面找
		myId, success := c.Get("TokenId")
		if !success {
			zap.L().Error("无法获取用户的TokenId！")
			return
		}
		// 遍历列表进行查找
		for _, user := range responseFollowerList.UserList {
			if user.Id == myId {
				isFollow = true
			}
		}
	}()

	wgUser.Wait()
	zap.L().Info("User并行查询返回成功！")
	return
}

func videoService(c *gin.Context, idString string) (favoriteCount, commentCount int64, isFavorite bool) {
	var wgVideo sync.WaitGroup
	// 并行查询服务
	wgVideo.Add(2)
	go func() {
		defer wgVideo.Done()
		myId, success := c.Get("TokenId")
		if !success {
			zap.L().Error("无法获取用户的TokenId！")
			return
		}
		fmt.Println(myId)
	}()
	go func() {
		defer wgVideo.Done()
	}()
	wgVideo.Wait()
	return
}