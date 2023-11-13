package main

import "_/topic-RabbitMQ/RabbitMQ"

func main() {
	kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "#")
	kutengOne.RecieveTopic()
}
