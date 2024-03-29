package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
	"time"
)

type PublishAPI interface {
	// Publish sends a raw message with the specified routingKey
	Publish(key string, msg []byte) error

	// PublishJSON sends a JSON message with the specified routingKey
	PublishJSON(key string, payload interface{}) error

	// PublishApi sends the payload using the supplied routing key.
	// []byte and string are sent as-is otherwise the message is marshaled into JSON before sending.
	PublishApi(key string, msg interface{}) error

	// Post is the underlying function used by the Publish functions.
	// It sends the actual message to RabbitMQ.
	Post(key string, body []byte, headers amqp.Table, timestamp time.Time) error
}

// Publisher represents the configuration for how to send messages to RabbitMQ
type Publisher struct {
	Exchange   string            `yaml:"exchange"`  // Exchange to submit to
	Mandatory  bool              `yaml:"mandatory"` // Publish mode
	Immediate  bool              `yaml:"immediate"` // Publish mode
	Replace    map[string]string `yaml:"replace"`   // Replace prefix table
	Ignore     []string          `yaml:"ignore"`    // Ignore prefixes
	Debug      bool              `yaml:"debug"`     // Debug mode
	Disabled   bool              `yaml:"disabled"`  // Publish disabled
	Tag        string            `yaml:"tag"`       // Tag for connection
	connection *amqp.Connection  // amqp connection
	channel    *amqp.Channel
	mq         *MQ
}

func (p *Publisher) Bind(broker *MQ) error {
	if p.mq != nil {
		return fmt.Errorf("publisher already bound to a broker")
	}
	p.mq = broker

	if p.Exchange == "" {
		p.Exchange = "amq.topic"
	}

	return nil
}

func (p *Publisher) Stop() {
	if p.connection != nil {
		_ = p.connection.Close()
		p.connection = nil
	}
}

// Publish sends a raw message with the specified routingKey
func (p *Publisher) Publish(key string, msg []byte) error {
	return p.Post(key, msg, nil, time.Now())
}

// PublishJSON sends a JSON message with the specified routingKey
func (p *Publisher) PublishJSON(key string, payload interface{}) error {
	msg, err := json.Marshal(payload)
	if err == nil {
		err = p.Publish(key, msg)
	}
	return err
}

// PublishApi sends the payload using the supplied routing key.
// []byte and string are sent as-is otherwise the message is marshaled into JSON before sending.
func (p *Publisher) PublishApi(key string, msg interface{}) error {
	var data []byte

	if b, ok := msg.([]byte); ok {
		data = b
	} else if s, ok := msg.(string); ok {
		data = []byte(s)
	} else {
		b, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		data = b
	}

	return p.Publish(key, data)
}

// Post is the underlying function used by the Publish functions.
// It sends the actual message to RabbitMQ.
func (p *Publisher) Post(key string, body []byte, headers amqp.Table, timestamp time.Time) error {

	key = p.EncodeKey(key)

	// Check for ignored entries
	for _, v := range p.Ignore {
		if strings.HasPrefix(key, v+".") {
			if p.Debug {
				log.Printf("Ignoring %q:%q %s", p.Exchange, key, body)
			}
			return nil
		}
	}

	if p.Debug {
		log.Printf("Post %q:%q %s", p.Exchange, key, body)
	}

	if p.Disabled {
		return nil
	}

	if p.connection == nil || p.connection.IsClosed() || p.channel == nil || p.channel.IsClosed() {
		p.Stop()
		log.Printf("Publisher Start %q", p.Tag)
		conn, err := p.mq.connect(p.Tag)
		if err != nil {
			return err
		}
		p.connection = conn

		ch, err := conn.Channel()
		if err != nil {
			p.Stop()
			return err
		}
		p.channel = ch
	}

	return p.channel.PublishWithContext(
		context.Background(),
		p.Exchange,
		key,
		p.Mandatory,
		p.Immediate,
		amqp.Publishing{
			Headers:         headers,         // Message headers
			Timestamp:       timestamp,       // Timestamp of message
			ContentType:     "text/json",     // MIME content type
			ContentEncoding: "",              // MIME content encoding
			Priority:        0,               // Normal priority
			AppId:           p.Tag,           // Tag of publisher
			DeliveryMode:    amqp.Persistent, // Required to persist messages on broker restart
			Body:            body,            // Message body
		},
	)
}

// EncodeKey converts any spaces or / in the DeliveryKey to . so they are more compatible with
// graphite. The output will also in lower case.
func (p *Publisher) EncodeKey(key string) string {
	key = EncodeKey(key)
	if p.Replace != nil {
		for k, v := range p.Replace {
			if strings.HasPrefix(key, k+".") {
				a := v + key[len(k):]
				key = a
			}
		}
	}
	return key
}

func EncodeKey(key string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(key), " ", "."), "/", ".")
}
