package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Queue struct {
	Broker  string `yaml:"broker"`
	Topic   string `yaml:"topic"`
	mq      *MQ
	handler MessageHandler
}

func (s *MQ) publishHandler(_ mqtt.Client, msg mqtt.Message) {
	s.Log("%s: received topic %q %s", msg.Topic(), msg.Payload())

	s.getSubscription(msg.Topic()).Do(msg)
}

func (s *MQ) getSubscription(topic string) MessageHandler {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.subscriptions[topic]
}

func (s *MQ) addSubscription(topic string, h MessageHandler) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriptions[topic] = h
}

func (s *MQ) Subscribe(topic string, h MessageHandler) error {
	s.Log("Subscribing to %q", topic)
	s.addSubscription(topic, h)
	return s.wait(s.client.Subscribe(topic, 1, nil), "subscribe")
}

func (q *Queue) Bind(broker *MQ) error {
	if q.mq != nil {
		return fmt.Errorf("MQTT:%s: Queue %s already bound", broker.name, q.Topic)
	}

	q.mq = broker

	return broker.Subscribe(q.Topic, q.handleMessage)
}

func (q *Queue) handleMessage(msg mqtt.Message) {
	q.handler.Do(msg)
}

func (q *Queue) AddHandler(h MessageHandler) {
	q.handler = q.handler.Then(h)
}
