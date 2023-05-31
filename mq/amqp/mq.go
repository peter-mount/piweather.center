package amqp

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/task"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type MQ struct {
	// Url to connect to
	Url string `json:"url" xml:"url,attr" yaml:"url"`
	// Exchange for publishing, defaults to amq.topic
	Exchange string `json:"exchange,omitempty" xml:"exchange,omitempty" yaml:"exchange,omitempty"`
	// ConnectionName that appears in the management plugin
	ConnectionName string `json:"connectionName,omitempty" xml:"connectionName,omitempty" yaml:"connectionName,omitempty"`
	// HeartBeat in seconds. Defaults to 10
	HeartBeat int `json:"heartBeat,omitempty" xml:"heartBeat,attr,omitempty" yaml:"heartBeat,omitempty"`
	// Product name in the management plugin (optional)
	Product string `json:"product" xml:"product" yaml:"product"`
	// Version in the management plugin (optional)
	Version string `json:"version" xml:"version" yaml:"version"`
	// ===== Internal
	name       string
	connection *amqp.Connection // amqp connection
	channel    *amqp.Channel    // amqp channel
	worker     task.Queue
}

func (s *MQ) exchange() string {
	if s.Exchange == "" {
		return "amq.topic"
	}
	return s.Exchange
}

func (s *MQ) Log(f string, a ...interface{}) {
	log.Printf("AMQP:"+s.name+":"+f, a...)
}

// Connect connects to the RabbitMQ instance that's been configured.
func (s *MQ) Connect() error {
	if s.connection != nil {
		return nil
	}

	s.Log("connecting")

	var heartBeat = s.HeartBeat
	if heartBeat == 0 {
		heartBeat = 10
	}

	// Use the user provided client name
	if connection, err := amqp.DialConfig(s.Url, amqp.Config{
		Heartbeat: time.Duration(heartBeat) * time.Second,
		Properties: amqp.Table{
			"product":         s.Product,
			"version":         s.Version,
			"connection_name": s.ConnectionName,
		},
		Locale: "en_US",
	}); err != nil {
		return err
	} else {
		s.connection = connection
	}

	if channel, err := s.NewChannel(); err != nil {
		return err
	} else {
		s.channel = channel
	}

	s.Log("connected")

	return s.channel.ExchangeDeclare(s.exchange(), "topic", true, false, false, false, nil)
}

func (s *MQ) NewChannel() (*amqp.Channel, error) {
	if channel, err := s.connection.Channel(); err != nil {
		return nil, err
	} else {
		return channel, nil
	}
}

func (s *MQ) AttachPublisher(pub *Publisher) error {
	if pub.mq == nil {
		ch, err := s.NewChannel()
		if err != nil {
			return err
		}
		pub.channel = ch
		pub.mq = s
	}
	return nil
}

func (s *MQ) Bind(queue *Queue) error {
	if queue.mq != nil {
		return fmt.Errorf("queue %q already bound", queue.Name)
	}

	channel, err := s.NewChannel()
	if err != nil {
		return err
	}
	queue.channel = channel

	// Create our queue
	_, err = s.QueueDeclare(channel, queue.Name, queue.Durable, queue.AutoDelete, false, false, nil)
	if err != nil {
		return err
	}

	for _, binding := range queue.Binding {
		topic := binding.Topic
		if topic == "" {
			topic = "amq.topic"
		}

		s.Log("bind %q %q:%q", queue.Name, binding.Topic, binding.Key)

		err = s.QueueBind(channel, queue.Name, binding.Key, topic, false, nil)
		if err != nil {
			return err
		}
	}

	queue.mq = s
	return nil
}

func (s *MQ) Consume(queue *Queue, tag string, autoAck bool, task Task) error {
	if queue.mq == nil {
		if err := s.Bind(queue); err != nil {
			return err
		}
	}

	s.Log("consuming %q", queue.Name)

	ch, err := s.rawConsume(queue.channel, queue.Name, tag, autoAck, false, false, false, nil)
	if err != nil {
		return err
	}

	t := task
	if !autoAck {
		t = queue.nakTask(t)
	}
	go queue.process(ch, t)

	return nil
}

// Publish a message
func (s *MQ) Publish(routingKey string, msg []byte) error {
	return s.channel.Publish(
		s.exchange(),
		routingKey,
		false,
		false,
		amqp.Publishing{
			Body: msg,
		})
}

// QueueDeclare declares a queue
func (s *MQ) QueueDeclare(channel *amqp.Channel, name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

// QueueBind binds a queue to an exchange & routingKey
func (s *MQ) QueueBind(channel *amqp.Channel, name, key, exchange string, noWait bool, args amqp.Table) error {
	return channel.QueueBind(name, key, exchange, noWait, args)
}

// Consume adds a consumer to the queue and returns a GO channel
func (s *MQ) rawConsume(channel *amqp.Channel, queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}
