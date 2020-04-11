package main

func main() {
	rabbitmq := NewRabbitMQSimple("" + "test")
	rabbitmq.ConsumeSimple()
}
