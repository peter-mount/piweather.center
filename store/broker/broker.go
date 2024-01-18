package broker

import (
	"fmt"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/rabbitmq/amqp091-go"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func init() {
	kernel.RegisterAPI((*DatabaseBroker)(nil), &broker{})
}

// DatabaseBroker is used by the db server, ingress and egress to manage
// the RabbitMQ connection between them
type DatabaseBroker interface {
	Exchange() string
	Consume(queue *amqp.Queue, tag string, task amqp.Task) error
	ConsumeKeys(queue *amqp.Queue, tag string, task amqp.Task, keys ...string) error
	PublishMetric(metric api.Metric) error
	amqp.PublishAPI
}

const (
	brokerName = "database"
)

type broker struct {
	Amqp      amqp.Pool `kernel:"inject"`
	mutex     sync.Mutex
	mq        *amqp.MQ
	queues    []*amqp.Queue
	publisher *amqp.Publisher
	appName   string
}

func (s *broker) Start() error {
	s.appName = filepath.Base(os.Args[0])

	s.mq = s.Amqp.GetMQ(brokerName)
	if s.mq == nil {
		return fmt.Errorf("broker %q undefined", brokerName)
	}

	s.mq.ConnectionName = s.appName + " Database"
	s.mq.Version = version.Version

	if s.mq.Exchange == "" {
		s.mq.Exchange = "amq.topic"
	}

	return nil
}

func (s *broker) Stop() {
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

func (s *broker) Exchange() string {
	return s.mq.Exchange
}

func (s *broker) tag(tag string) string {
	return strings.TrimSpace(strings.Join([]string{s.appName, "database", tag}, " "))
}

func (s *broker) Consume(queue *amqp.Queue, tag string, task amqp.Task) error {
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

func (s *broker) ConsumeKeys(queue *amqp.Queue, tag string, task amqp.Task, keys ...string) error {
	for _, key := range keys {
		queue.AddBinding(amqp.Binding{
			Topic: s.Exchange(),
			Key:   key,
		})
	}

	return s.Consume(queue, tag, task)
}

func (s *broker) Publish(key string, msg []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.connectPublisher(); err != nil {
		return err
	}
	return s.publisher.Publish(key, msg)
}

func (s *broker) PublishJSON(key string, payload interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.connectPublisher(); err != nil {
		return err
	}
	return s.publisher.PublishJSON(key, payload)
}

func (s *broker) PublishApi(key string, msg interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.connectPublisher(); err != nil {
		return err
	}
	return s.publisher.PublishApi(key, msg)
}

func (s *broker) Post(key string, body []byte, headers amqp091.Table, timestamp time.Time) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.connectPublisher(); err != nil {
		return err
	}
	return s.publisher.Post(key, body, headers, timestamp)
}

func (s *broker) PublishMetric(metric api.Metric) error {
	return s.PublishJSON("metric."+amqp.EncodeKey(metric.Metric), metric)
}

func (s *broker) connectPublisher() error {

	if s.publisher != nil {
		return nil
	}

	pub := &amqp.Publisher{
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
