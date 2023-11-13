package main

import (
	"_/eva-RabbitMQ/RabbitMQ"
	"fmt"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"queue1")
	ticker := time.NewTicker(1 * time.Second)
	for i := 0; i < 100; i++ {
		<-ticker.C
		rabbitmq.PublishSimple(fmt.Sprintf("Hello kuteng%d", i))
	}
	ticker.Stop()

	fmt.Println("发送成功！")
}
