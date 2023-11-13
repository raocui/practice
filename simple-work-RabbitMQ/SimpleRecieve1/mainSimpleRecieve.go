package main

import "_/eva-RabbitMQ/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"queue1")
	rabbitmq.ConsumeSimple()
}
