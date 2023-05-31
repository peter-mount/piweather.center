package amqp

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// Queue represents the configuration for declaring a queue in RabbitMQ
type Queue struct {
	Broker     string           `json:"broker" xml:"broker,attr" yaml:"broker"`
	Name       string           `json:"name" xml:"name,attr" yaml:"name"`
	Binding    []Binding        `json:"binding" xml:"binding" yaml:"binding"`
	Durable    bool             `json:"durable" xml:"durable,attr,omitempty" yaml:"durable,omitempty"`
	AutoDelete bool             `json:"autoDelete" xml:"autoDelete,attr" yaml:"autoDelete,omitempty"`
	mq         *MQ              // Broker definition
	connection *amqp.Connection // amqp connection
	notify     chan *amqp.Error // Connection error notifications
	channel    *amqp.Channel    // Active channel
}

// Binding defines a binding for a Queue
type Binding struct {
	Topic string `json:"topic" xml:"topic,attr" yaml:"topic"`
	Key   string `json:"key" xml:"key,attr" yaml:"key"`
}

func (q *Queue) Bind(broker *MQ) error {
	if q.mq != nil {
		return fmt.Errorf("Queue %s already bound to a broker", q.Name)
	}
	q.mq = broker
	return nil
}

func (q *Queue) Start(tag string, autoAck bool, task Task) error {
	log.Printf("Queue Start %q", tag)
	conn, err := q.mq.connect(tag)
	if err != nil {
		return err
	}
	q.connection = conn

	q.notify = conn.NotifyClose(make(chan *amqp.Error))

	ch, err := conn.Channel()
	if err != nil {
		q.Stop()
		return err
	}
	q.channel = ch

	_, err = ch.QueueDeclare(q.Name, q.Durable, q.AutoDelete, false, false, nil)
	if err != nil {
		return err
	}

	for _, binding := range q.Binding {
		topic := binding.Topic
		if topic == "" {
			topic = "amq.topic"
		}

		log.Printf("bind %q %q:%q", q.Name, binding.Topic, binding.Key)

		err = ch.QueueBind(q.Name, binding.Key, topic, false, nil)
		if err != nil {
			return err
		}
	}

	msgs, err := ch.Consume(q.Name, tag, autoAck, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case notification := <-q.notify:
				if notification == nil {
					// We are stopping so just close the connection
					q.Stop()
				} else {
					// Try to reconnect
					q.logError(q.reconnect(notification, tag, autoAck, task))
				}
				// Stop this goroutine
				return

			case msg := <-msgs:
				if autoAck {
					q.logError(task(msg))
				} else {
					if err := task(msg); err != nil {
						q.logError(err)
						q.logError(msg.Nack(false, true))
					} else {
						q.logError(msg.Ack(false))
					}
				}
			}
		}
	}()

	return nil
}

func (q *Queue) reconnect(notification *amqp.Error, tag string, autoAck bool, task Task) error {
	log.Printf("Reconnecting queue %q for %d %q", q.Name, notification.Code, notification.Reason)
	q.Stop()
	time.Sleep(5 * time.Second)
	return q.Start(tag, autoAck, task)
}

func (q *Queue) Stop() {
	if q.connection != nil {
		_ = q.connection.Close()
		q.connection = nil
	}
}

func (q *Queue) logError(err error) {
	if err != nil {
		log.Printf("error, queue=%q, error=%v", q.Name, err)
	}
}
