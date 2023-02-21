/**
* @Author Wang Hui
* @Description
* @Date
**/
package mq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"simple-DY/DY-api/social-web/global"
)

func NewClient() *amqp.Connection {
	amqpConfig := global.ServerConfig.AmqpConfig
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		amqpConfig.User, amqpConfig.Password, amqpConfig.Host, amqpConfig.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		zap.S().Error("MQ NewClient error:", err)
	}
	return conn
}
func sendMsg(v interface{}) {
	conn := NewClient()
	ch, err := conn.Channel()
	if err != nil {
		zap.S().Error("Failed to open a channel:", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"chat_msg", // 队列名称
		false,      // 是否持久化
		false,      // 是否自动删除
		false,      // 是否独占
		false,      // 是否等待服务器响应
		nil,        // 额外参数
	)
	if err != nil {
		zap.S().Error("Failed to declare a queue:", err)
		return
	}
	data := []byte{}
	err = json.Unmarshal(data, v)
	if err != nil {
		zap.S().Error("sendMsg json.Unmarshal:", err)
		return
	}
	err = ch.Publish(
		"",     // 交换机名称
		q.Name, // 路由键
		false,  // 是否强制发送到队列
		false,  // 是否等待服务器响应
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		zap.S().Error("Failed to publish a message:", err)
		return
	}
	log.Printf(" [x] Sent %s", v)
}

func rcvMsg(v interface{}) {
	conn := NewClient()
	ch, err := conn.Channel()
	if err != nil {
		zap.S().Error("Failed to open a channel:", err)
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"chat_msg", // 队列名称
		false,      // 是否持久化
		false,      // 是否自动删除
		false,      // 是否独占
		false,      // 是否等待服务器响应
		nil,        // 额外参数
	)
	if err != nil {
		zap.S().Error("Failed to declare a queue:", err)
		return
	}
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者名称
		true,   // 是否自动应答
		false,  // 是否独占
		false,  // 是否等待服务器响应
		false,  // 额外参数
		nil,
	)
	if err != nil {
		zap.S().Error("Failed to register a consumer:", err)
		return
	}
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
