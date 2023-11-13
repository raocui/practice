package main

import "_/routing-RabbitMQ/RabbitMQ"

func main() {
	kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	kutengtwo.RecieveRouting()
}
