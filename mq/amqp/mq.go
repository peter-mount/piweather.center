package amqp

import (
	"github.com/peter-mount/go-kernel/v2/log"
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
	name string
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
func (s *MQ) connect(name string) (*amqp.Connection, error) {
	if name == "" {
		name = s.ConnectionName
	}
	if name == "" {
		name = s.name
	}
	s.Log("connecting %q", name)

	var heartBeat = s.HeartBeat
	if heartBeat == 0 {
		heartBeat = 10
	}

	// Use the user provided client name
	connection, err := amqp.DialConfig(s.Url, amqp.Config{
		Heartbeat: time.Duration(heartBeat) * time.Second,
		Properties: amqp.Table{
			"product":         s.Product,
			"version":         s.Version,
			"connection_name": name,
		},
		Locale: "en_US",
	})
	if err != nil {
		return nil, err
	}

	s.Log("connected")
	return connection, nil
}
