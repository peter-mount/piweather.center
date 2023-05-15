package mq

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/rabbitmq"
	"github.com/peter-mount/go-kernel/v2/util/task"
)

type MQ struct {
	mq     *rabbitmq.RabbitMQ `kernel:"config,mq"`
	worker task.Queue         `kernel:"worker"`
}

func (m *MQ) Start() error {
	return m.mq.Connect()
}

func (m *MQ) AttachPublisher(pub *Publisher) error {
	if pub.mq == nil {
		ch, err := m.mq.NewChannel()
		if err != nil {
			return err
		}
		pub.channel = ch
		pub.mq = m
	}
	return nil
}

func (m *MQ) Bind(queue *Queue) error {
	if queue.mq != nil {
		return fmt.Errorf("queue %q already bound", queue.Name)
	}

	channel, err := m.mq.NewChannel()
	if err != nil {
		return err
	}
	queue.channel = channel

	// Create our queue
	_, err = m.mq.QueueDeclare(channel, queue.Name, queue.Durable, queue.AutoDelete, false, false, nil)
	if err != nil {
		return err
	}

	for _, binding := range queue.Binding {
		topic := binding.Topic
		if topic == "" {
			topic = "amq.topic"
		}

		err = m.mq.QueueBind(channel, queue.Name, binding.Key, topic, false, nil)
		if err != nil {
			return err
		}
	}

	queue.mq = m
	return nil
}

func (m *MQ) Consume(queue *Queue, tag string, autoAck bool, task Task) error {
	if queue.mq == nil {
		if err := m.Bind(queue); err != nil {
			return err
		}
	}

	ch, err := m.mq.Consume(queue.channel, queue.Name, tag, autoAck, false, false, false, nil)
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
