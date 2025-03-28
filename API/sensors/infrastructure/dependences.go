package infrastructure

import "API/sensors/infrastructure/adapters"

var rabbitmq *adapters.RabbitMQ
var postgres *adapters.PostgreSQL

func GoDependences() {
	rabbitmq = adapters.NewRabbitMQ()
	postgres = adapters.NewPostgreSQL()
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbitmq
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}
