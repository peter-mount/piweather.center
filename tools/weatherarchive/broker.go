package weatherarchive

import (
	"fmt"
	"github.com/peter-mount/go-build/version"
	amqp2 "github.com/peter-mount/piweather.center/util/mq/amqp"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Broker struct {
	Broker    *string    `kernel:"flag,archive-broker,Send archives to rabbitmq"`
	Exchange  *string    `kernel:"flag,archive-exchange,Exchange to use,amq.topic"`
	Amqp      amqp2.Pool `kernel:"inject"`
	mutex     sync.Mutex
	mq        *amqp2.MQ
	queues    []*amqp2.Queue
	publisher *amqp2.Publisher
	appName   string
}

func (s *Broker) IsActive() bool {
	return *s.Broker != ""
}

func (s *Broker) RoutingKey(name string) string {
	return "archive." + strings.ReplaceAll(strings.ToLower(name), "/", ".")
}

func (s *Broker) Start() error {
	if !s.IsActive() {
		return nil
	}

	s.mq = s.Amqp.GetMQ(*s.Broker)
	if s.mq == nil {
		return fmt.Errorf("broker %q undefined", *s.Broker)
	}

	s.appName = filepath.Base(os.Args[0])
	s.mq.ConnectionName = s.appName + " Archiver"
	s.mq.Version = version.Version

	if s.mq.Exchange == "" {
		s.mq.Exchange = *s.Exchange
	}
	if s.mq.Exchange == "" {
		s.mq.Exchange = "amq.topic"
	}

	return nil
}

func (s *Broker) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.publisher != nil {
		s.publisher.Stop()
		s.publisher = nil
	}

	for _, queue := range s.queues {
		queue.Stop()
	}
	s.queues = nil
}

func (s *Broker) tag(tag string) string {
	return strings.TrimSpace(strings.Join([]string{s.appName, "archiver", tag}, " "))
}

func (s *Broker) Consume(queue *amqp2.Queue, tag string, task amqp2.Task) error {
	if !s.IsActive() {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := queue.Bind(s.mq)
	if err == nil {
		err = queue.Start(s.tag(tag), false, task)
	}
	if err == nil {
		s.queues = append(s.queues, queue)
	}
	return err
}

func (s *Broker) ConsumeKeys(queue *amqp2.Queue, tag string, task amqp2.Task, keys ...string) error {
	for _, key := range keys {
		queue.AddBinding(amqp2.Binding{
			Topic: s.mq.Exchange,
			Key:   key,
		})
	}

	return s.Consume(queue, tag, task)
}

func (s *Broker) Publish(key string, msg []byte) error {
	if !s.IsActive() {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.connectPublisher(); err != nil {
		return err
	}
	return s.publisher.Publish(key, msg)
}

func (s *Broker) connectPublisher() error {

	if s.publisher != nil {
		return nil
	}

	pub := &amqp2.Publisher{
		Exchange:  s.mq.Exchange,
		Mandatory: false,
		Immediate: false,
		Tag:       s.tag("publisher"),
	}

	if err := pub.Bind(s.mq); err != nil {
		return err
	}

	s.publisher = pub
	return nil
}
