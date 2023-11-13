package main

import (
	"_/routing-RabbitMQ/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	for i := 0; i <= 100; i++ {
		kutengone.PublishRouting("Hello kuteng one!" + strconv.Itoa(i))
		kutengtwo.PublishRouting("Hello kuteng Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
