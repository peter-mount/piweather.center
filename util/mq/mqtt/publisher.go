package mqtt

import "fmt"

type Publisher struct {
	Topic    string `yaml:"topic"`              // Topic to publish to
	Retained bool   `yaml:"retained,omitempty"` // true if message to be retained
	mq       *MQ
}

func (p *Publisher) Bind(broker *MQ) error {
	if p.mq != nil {
		return fmt.Errorf("publisher already bound to a broker")
	}
	p.mq = broker

	return nil
}

func (p *Publisher) Publish(message []byte) error {
	return p.mq.publish(p.Topic, p.Retained, message)
}

func (s *MQ) publish(topic string, retained bool, message []byte) error {
	return s.wait(s.client.Publish(topic, 0, retained, message), "publish")
}
