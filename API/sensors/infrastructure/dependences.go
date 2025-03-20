package infrastructure

import "API/sensors/infrastructure/adapters"

var rabbitmq *adapters.RabbitMQ

func GoDependences() {
	rabbitmq = adapters.NewRabbitMQ()
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbitmq
}
