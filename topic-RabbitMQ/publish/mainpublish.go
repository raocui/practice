package main

import (
	"_/topic-RabbitMQ/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.one")
	kutengTwo := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.two")
	for i := 0; i <= 100; i++ {
		kutengOne.PublishTopic("Hello kuteng topic one!" + strconv.Itoa(i))
		kutengTwo.PublishTopic("Hello kuteng topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
