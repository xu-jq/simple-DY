/*
 * @Date: 2023-01-26 21:58:55
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-27 10:31:16
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/rabbitmq/producer.go
 * @Description:
 */
package rabbitmq

import (
	"encoding/json"
	"simple-DY/DY-srvs/video-srv/global"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// Simple模式的生产端
func PublishSimple(info Message) (err error) {

	// 连接消息队列
	conn, err := amqp.Dial("amqp://" + global.GlobalConfig.RabbitMQ.UserName + ":" + global.GlobalConfig.RabbitMQ.Password + "@" + global.GlobalConfig.RabbitMQ.Address + ":" + global.GlobalConfig.RabbitMQ.Port + "/" + global.GlobalConfig.RabbitMQ.VirtualHost)
	if err != nil {
		zap.L().Error("生产者连接消息队列失败！错误信息：" + err.Error())
		return
	}
	defer conn.Close()

	// 获取通道
	ch, err := conn.Channel()
	if err != nil {
		zap.L().Error("生产者获取通道失败！错误信息：" + err.Error())
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
		zap.L().Error("生产者声明队列失败！错误信息：" + err.Error())
		return
	}

	mes, err := json.Marshal(info)
	if err != nil {
		zap.L().Error("生产者结构体序列化为JSON失败！错误信息：" + err.Error())
		return
	}

	// 发送消息
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        mes, //消息的内容
		},
	)
	if err != nil {
		zap.L().Error("生产者发送消息失败！错误信息：" + err.Error())
	}
	zap.L().Info("生产者发送消息成功！")
	return
}
