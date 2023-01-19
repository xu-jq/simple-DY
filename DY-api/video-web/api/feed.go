/*
 * @Date: 2023-01-19 14:08:05
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 14:34:42
 * @FilePath: /simple-DY/DY-api/video-web/api/feed.go
 * @Description: 1.1 视频流接口
 */
package api

import (
	"fmt"
	"log"
	"net/http"
	"simple-DY/DY-api/video-web/models"

	"github.com/gin-gonic/gin"
)

// 1.1 视频流接口 /douyin/feed/
func Feed(c *gin.Context) {
	var feedRequest models.FeedRequest
	err := c.ShouldBindJSON(&feedRequest)
	if err == nil {
		log.Fatal("feedRequest error")
	}
	video1 := models.Video{
		Id: 1,
		Author: models.User{
			Id:            1,
			Name:          "1",
			FollowCount:   1,
			FollowerCount: 1,
			IsFollow:      true,
		},
		PlayUrl:       "http://121.37.98.68:81/videos/example1.mp4",
		CoverUrl:      "http://121.37.98.68:81/images/example1.jpg",
		FavoriteCount: 1,
		CommentCount:  1,
		IsFavorite:    true,
		Title:         "example1",
	}
	video2 := models.Video{
		Id: 2,
		Author: models.User{
			Id:            1,
			Name:          "1",
			FollowCount:   1,
			FollowerCount: 1,
			IsFollow:      true,
		},
		PlayUrl:       "http://121.37.98.68:81/videos/example2.mp4",
		CoverUrl:      "http://121.37.98.68:81/images/example2.jpg",
		FavoriteCount: 1,
		CommentCount:  1,
		IsFavorite:    true,
		Title:         "example2",
	}
	video3 := models.Video{
		Id: 3,
		Author: models.User{
			Id:            1,
			Name:          "1",
			FollowCount:   1,
			FollowerCount: 1,
			IsFollow:      true,
		},
		PlayUrl:       "http://121.37.98.68:81/videos/example3.mp4",
		CoverUrl:      "http://121.37.98.68:81/images/example3.jpg",
		FavoriteCount: 1,
		CommentCount:  1,
		IsFavorite:    true,
		Title:         "example3",
	}
	c.JSON(http.StatusOK, models.FeedResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: 0,
			StatusMsg:  "获取视频流成功",
		},
		NextTime:  1674048618617,
		VideoList: []models.Video{video1, video2, video3},
	})

	fmt.Println("success!")

}
