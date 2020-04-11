package main

import "fmt"

func main() {
	rabbitmq := NewRabbitMQSimple("" + "test")
	rabbitmq.PublishSimple("Hello test222!")
	fmt.Println("发送成功！")
}
