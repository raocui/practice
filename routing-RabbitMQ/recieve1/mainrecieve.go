package main

import "_/routing-RabbitMQ/RabbitMQ"

func main() {
	kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengone.RecieveRouting()
}
