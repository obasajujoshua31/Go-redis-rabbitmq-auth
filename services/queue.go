package services

import (
	"Go-redis-rabbitmq-auth/config"
	"encoding/json"

	"github.com/streadway/amqp"
)

const (
	emailConfirmation = "email-confirmation"
	duration          = true
	autoDelete        = false
	exclusive         = false
	nowait            = false
	preFetchCount     = 1
	preFetchSize      = 0
	global            = false
	noLocal           = false
	autoAck           = false
	consumer          = ""
	mandatory = false
	immediate = false
)

type Message struct {
	Subject     string
	ToEmail     string
	PlainText   string
	FromName    string
	FromEmail   string
	ToName      string
	HTMLContent string
}

type QueueChannel struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

//CreateConnection ...
func CreateConnection(config config.AppConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitMQURL)

	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	return channel, err

}

func CreateQueue(channel *amqp.Channel) (QueueChannel, error) {
	queue, err := channel.QueueDeclare(emailConfirmation, duration, autoDelete, exclusive, nowait, nil)

	if err != nil {
		return QueueChannel{}, err
	}

	err = channel.Qos(preFetchCount, preFetchSize, global)

	if err != nil {
		return QueueChannel{}, err
	}

	return QueueChannel{
		Channel: channel,
		Queue:   queue,
	}, nil
}

//PublishMessageToQueue ...
func (ch *QueueChannel) PublishMessageToQueue(message Message) error {
	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	err = ch.Channel.Publish("", ch.Queue.Name, mandatory, immediate, amqp.Publishing{
	  DeliveryMode: amqp.Persistent,
	  ContentType: "text/plain",
	  Body: body,
	})
	if err != nil {
		return err
	}
	return nil
}
