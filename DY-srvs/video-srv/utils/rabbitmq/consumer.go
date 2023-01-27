/*
 * @Date: 2023-01-26 21:58:55
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-27 10:38:17
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/rabbitmq/consumer.go
 * @Description:
 */
package rabbitmq

import (
	"encoding/json"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"simple-DY/DY-srvs/video-srv/utils/oss"
	"sync"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// simple模式的消费端
func ConsumeSimple() {
	defer global.Wg.Done()

	// 连接消息队列
	conn, err := amqp.Dial("amqp://" + global.GlobalConfig.RabbitMQ.UserName + ":" + global.GlobalConfig.RabbitMQ.Password + "@" + global.GlobalConfig.RabbitMQ.Address + ":" + global.GlobalConfig.RabbitMQ.Port + "/" + global.GlobalConfig.RabbitMQ.VirtualHost)
	if err != nil {
		zap.L().Error("消费者连接消息队列失败！错误信息：" + err.Error())
		return
	}
	defer conn.Close()

	// 获取通道
	ch, err := conn.Channel()
	if err != nil {
		zap.L().Error("消费者获取通道失败！错误信息：" + err.Error())
		return
	}

	// 队列声明
	q, err := ch.QueueDeclare(
		"OSS",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("消费者声明队列失败！错误信息：" + err.Error())
		return
	}

	//消费者接收消息，msgs为只读通道
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		zap.L().Error("消费者接收消息失败！错误信息：" + err.Error())
		return
	}

	zap.L().Info("消费者正在监听消息队列......")

	for v := range msgs {
		var info Message
		zap.L().Info("消费者接收到的消息：" + string(v.Body))

		// 将字符串反解析为结构体
		json.Unmarshal(v.Body, &info)

		// 并行上传视频文件和图片文件
		success := make([]bool, 2)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			// 上传视频文件
			zap.L().Info("视频文件开始上传！路径：" + info.VideoOSSFileName)
			err = oss.UploadFileToQiniuOSS(info.VideoStaticFileName, info.VideoOSSFileName)
			if err != nil {
				zap.L().Error("无法上传视频文件！错误信息：" + err.Error())
				return
			}
			zap.L().Info("视频文件上传成功！路径：" + info.VideoOSSFileName)
			success[0] = true
		}()
		go func() {
			defer wg.Done()
			// 上传图片文件
			zap.L().Info("图片文件开始上传！路径：" + info.ImageOSSFileName)
			err = oss.UploadFileToQiniuOSS(info.ImageStaticFileName, info.ImageOSSFileName)
			if err != nil {
				zap.L().Error("无法上传图片文件！错误信息：" + err.Error())
				return
			}
			zap.L().Info("图片文件上传成功！路径：" + info.ImageOSSFileName)
			success[1] = true
		}()
		wg.Wait()

		// 判断两个文件是否全部上传成功，如果全部上传成功则写入数据库
		if success[0] && success[1] {
			// 向数据库中插入数据
			dao.InsertVideo(info.AuthorId, info.FileName, info.Time, info.Title)
			zap.L().Info("数据库写入成功！")
		} else {
			zap.L().Error("文件上传出错！不能写入数据库！")
		}

	}
}
