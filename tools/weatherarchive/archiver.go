package weatherarchive

import (
	"context"
	"errors"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/tools/weatheringress/payload"
	"github.com/peter-mount/piweather.center/util/mq/amqp"
	"github.com/rabbitmq/amqp091-go"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Archiver is a generic service which can accept raw data from Ingress or RabbitMQ and
// archive them into either log files, the console or post to RabbitMQ for remote archiving.
//
// Note: When used with RabbitMQ, you can either publish to or consume from but not both.
type Archiver struct {
	StoreDir    *string        `kernel:"flag,archive-dir,Archive directory"`
	LogMessages *bool          `kernel:"flag,archive-log,Dump messages to stdout"`
	Publish     *bool          `kernel:"flag,archive-publish,Publish to broker"`
	QueueName   *string        `kernel:"flag,archive-queue,Queue to consume from broker"`
	Broker      *Broker        `kernel:"inject"`
	Worker      task.Queue     `kernel:"worker"`
	Daemon      *kernel.Daemon `kernel:"inject"`
	mqQueue     *amqp.Queue
	task        task.Task
	mutex       sync.Mutex
}

func (s *Archiver) PostInit() error {
	// We cannot consume and publish at the same time
	if *s.QueueName != "" && *s.Publish {
		return errors.New("-archive-publish and -archive-queue are mutually exclusive")
	}

	if *s.LogMessages {
		s.task = s.task.Then(s.archiveReadingLog)
	}

	if *s.StoreDir != "" {
		s.task = s.task.Then(s.archiveReadingDisk)
	}

	if *s.Publish {
		s.task = s.task.Then(s.archiveReadingPublish)
	}

	// If we have a task then enforce us into daemon mode unless we have a webserver
	if s.task != nil && !s.Daemon.IsWebserver() {
		s.Daemon.SetDaemon()
	}

	return nil
}

func (s *Archiver) Start() error {
	// Archiver is disabled if no task
	if s.task == nil {
		return nil
	}

	if *s.QueueName != "" {

		s.mqQueue = &amqp.Queue{
			Name:       *s.QueueName,
			Durable:    true,
			AutoDelete: false,
		}

		err := s.Broker.ConsumeKeys(s.mqQueue, "archive", s.archiveAmqp, "archive.#")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Archiver) archiveFileName(name string, t time.Time) string {
	return ArchiveFileName(*s.StoreDir, name, t)
}

func ArchiveFileName(dir, name string, t time.Time) string {
	p := filepath.Join(strings.Split(name, ".")...)
	return filepath.Join(dir, p, t.UTC().Format("2006/01/02")+".log")
}

func (s *Archiver) Archive(ctx context.Context) error {
	s.archivePayload(payload.GetPayload(ctx))
	return nil
}

func (s *Archiver) archiveAmqp(delivery amqp091.Delivery) error {
	p, err := payload.FromLog(string(delivery.Body))
	if err == nil {
		s.archivePayload(p)
	}
	return nil
}

func (s *Archiver) archivePayload(rec *payload.Payload) {
	if s.task != nil && rec != nil {
		s.Worker.AddTask(task.Of(s.task).WithValue("record", rec))
	}
}

func (s *Archiver) archiveReadingLog(ctx context.Context) error {
	rec := ctx.Value("record").(*payload.Payload)
	log.Printf("Archive:%v", string(rec.Msg()))
	return nil
}

func (s *Archiver) archiveReadingPublish(ctx context.Context) error {
	rec := ctx.Value("record").(*payload.Payload)
	_ = s.Broker.Publish("archive."+strings.ToLower(rec.Id()), rec.Msg())
	return nil
}

func (s *Archiver) archiveReadingDisk(ctx context.Context) error {
	rec := ctx.Value("record").(*payload.Payload)
	_ = s.appendReading(s.archiveFileName(rec.Id(), rec.Time()), rec.ToLog())
	return nil
}

func (s *Archiver) appendReading(fileName string, msg string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(msg)

	return nil
}
