package main

import (
	"Go-redis-rabbitmq-auth/server"
	"log"
)

func main() {
	err := server.StartServer()

	if err != nil {
		log.Fatalf(err.Error())
	}
}
