package mq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strings"
	"time"
)

// Publisher represents the configuration for how to send messages to RabbitMQ
type Publisher struct {
	Exchange  string            `yaml:"exchange"`  // Exchange to submit to
	Mandatory bool              `yaml:"mandatory"` // Publish mode
	Immediate bool              `yaml:"immediate"` // Publish mode
	Replace   map[string]string `yaml:"replace"`   // Replace prefix table
	Ignore    []string          `yaml:"ignore"`    // Ignore prefixes
	Debug     bool              `yaml:"debug"`     // Debug mode
	Disabled  bool              `yaml:"disabled"`  // Publish disabled
	channel   *amqp.Channel
	mq        *MQ
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

	return p.channel.Publish(
		p.Exchange,
		key,
		p.Mandatory,
		p.Immediate,
		amqp.Publishing{
			Headers:   headers,
			Timestamp: timestamp,
			Body:      body,
		},
	)
}

// EncodeKey converts any spaces or / in the DeliveryKey to . so they are more compatible with
// graphite. The output will also in lower case.
func (p *Publisher) EncodeKey(key string) string {
	key = EncodeKey(key)
	for k, v := range p.Replace {
		if strings.HasPrefix(key, k+".") {
			a := v + key[len(k):]
			key = a
		}
	}
	return key
}

func EncodeKey(key string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(key), " ", "."), "/", ".")
}
