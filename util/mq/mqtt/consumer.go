package mqtt

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/peter-mount/go-kernel/v2/util/task"
)

const (
	DeliveryKey = "mqtt.Message"
)

// Delivery returns the mqtt.Message message from a context when used by a worker task
func Delivery(ctx context.Context) mqtt.Message {
	return ctx.Value(DeliveryKey).(mqtt.Message)
}

// Message returns the message content from a context when used by a worker task
func Message(ctx context.Context) []byte {
	return Delivery(ctx).Payload()
}

func ContextTask(f task.Task, ctx context.Context) MessageHandler {
	return func(msg mqtt.Message) {
		_ = f.WithValue(DeliveryKey, msg).Do(ctx)
	}
}
