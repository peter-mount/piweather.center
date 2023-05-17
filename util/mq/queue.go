package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// Queue represents the configuration for declaring a queue in RabbitMQ
type Queue struct {
	Broker     string    `json:"broker" xml:"broker,attr" yaml:"broker"`
	Name       string    `json:"name" xml:"name,attr" yaml:"name"`
	Binding    []Binding `json:"binding" xml:"binding" yaml:"binding"`
	Durable    bool      `json:"durable" xml:"durable,attr,omitempty" yaml:"durable,omitempty"`
	AutoDelete bool      `json:"autoDelete" xml:"autoDelete,attr" yaml:"autoDelete,omitempty"`
	channel    *amqp.Channel
	mq         *MQ
}

// Binding defines a binding for a Queue
type Binding struct {
	Topic string `json:"topic" xml:"topic,attr" yaml:"topic"`
	Key   string `json:"key" xml:"key,attr" yaml:"key"`
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
