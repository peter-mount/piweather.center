package amqp

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/util/task"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	DeliveryKey = "amqp.Delivery"
)

// Delivery returns the amqp.Delivery message from a context when used by a worker task
func Delivery(ctx context.Context) amqp.Delivery {
	return ctx.Value(DeliveryKey).(amqp.Delivery)
}

// Message returns the message content from a context when used by a worker task
func Message(ctx context.Context) []byte {
	return Delivery(ctx).Body
}

func ContextTask(f task.Task, ctx context.Context) Task {
	return func(msg amqp.Delivery) error {
		return f.WithValue(DeliveryKey, msg).Do(ctx)
	}
}

/*
func (s *MQ) ConsumeTask(queue *Queue, tag string, f task.Task) error {
	return s.Consume(queue, tag,
		true,
		func(msg amqp.Delivery) error {
			s.worker.AddTask(f.WithValue(DeliveryKey, msg))
			return nil
		})
}

func (s *MQ) ConsumePriorityTask(queue *Queue, tag string, priority int, f task.Task) error {
	return s.Consume(queue, tag,
		true,
		func(msg amqp.Delivery) error {
			s.worker.AddPriorityTask(priority, f.WithValue(DeliveryKey, msg))
			return nil
		})
}

func (s *MQ) AddPriorityTask(priority int, f task.Task) Task {
	return func(msg amqp.Delivery) error {
		s.worker.AddPriorityTask(priority, f.WithValue(DeliveryKey, msg))
		return nil
	}
}

func Guard(b task.Task) task.Task {
	return func(ctx context.Context) error {
		if err := b.Do(ctx); err != nil {
			msg := Delivery(ctx)
			log.Printf("Error on %q\n\n%s\n\n%v", msg.RoutingKey, msg.Body, err)
			//return err
		}
		return nil
	}
}
*/
