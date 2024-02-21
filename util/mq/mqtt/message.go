package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MessageHandler func(mqtt.Message)

func (a MessageHandler) Then(b MessageHandler) MessageHandler {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(m mqtt.Message) {
		a(m)
		b(m)
	}
}

func (a MessageHandler) Do(m mqtt.Message) {
	if a != nil {
		a(m)
	}
}
