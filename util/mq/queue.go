package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// Queue represents the configuration for declaring a queue in RabbitMQ
type Queue struct {
	Name       string    `yaml:"name"`
	Binding    []Binding `yaml:"binding"`
	Durable    bool      `yaml:"durable"`
	AutoDelete bool      `yaml:"autoDelete"`
	channel    *amqp.Channel
	mq         *MQ
}

// Binding defines a binding for a Queue
type Binding struct {
	Topic string `yaml:"topic"`
	Key   string `yaml:"key"`
}

func (q *Queue) process(ch <-chan amqp.Delivery, f Task) {
	for {
		msg := <-ch
		q.logError(f(msg))
	}
}

func (q *Queue) logError(err error) {
	if err != nil {
		log.Printf("error, queue=%q, error=%v", q.Name, err)
	}
}
