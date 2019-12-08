package server

import (
	"Go-redis-rabbitmq-auth/config"
	"Go-redis-rabbitmq-auth/services"
)

func StartServer() error {

	config, err := config.LoadEnv()
	if err != nil {
		return err
	}

	conn, err := services.CreateConnection(config)

	if err != nil {
		return err
	}

	defer conn.Close()
	channel, err := services.CreateChannel(conn)

	if err != nil {
		return err
	}
	defer channel.Close()
	

	queue, err := services.CreateQueue(channel)

	if err != nil {
		return err
	}

	message := services.Message{
		Subject:     "This is the subject",
		ToEmail:     "obasajujoshua31@gmail.com",
		PlainText:   "This is my message",
		FromName:    "Joshua Obasaju",
		ToName:      "Joshua Obasaju",
		FromEmail:   "obasajujoshua31@gmail.com",
		HTMLContent: `<h3 style="color:green" >This is another message</h3>`,
	}

	err = queue.PublishMessageToQueue(message)

	if err != nil {
		return err
	}

	return nil

}
