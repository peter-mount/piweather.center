package weatheregress

import (
	"flag"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
)

type Egress struct {
	Amqp           amqp.Pool             `kernel:"inject"`
	Mqtt           mqtt.Pool             `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Daemon         *kernel.Daemon        `kernel:"inject"`
	QueueName      *string               `kernel:"flag,metric-queue,DB queue name,database.calc"`
	mqQueue        *amqp.Queue
	script         *lang.Script
	processor      lang.Visitor
}

func (s *Egress) Start() error {
	p := lang.NewParser()
	script, err := p.ParseFiles(flag.Args()...)
	if err != nil {
		return err
	}
	s.script = script

	s.initProcessor()

	s.mqQueue = &amqp.Queue{
		Name:       *s.QueueName,
		Durable:    true,
		AutoDelete: false,
	}

	err = s.DatabaseBroker.ConsumeKeys(s.mqQueue, "egress", s.processMetricUpdate, "metric.#")
	if err != nil {
		return err
	}

	log.Println(version.Version)

	// Mark the application as a daemon
	s.Daemon.SetDaemon()

	return nil
}
