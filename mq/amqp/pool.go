package amqp

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/util/task"
	common "github.com/peter-mount/piweather.center"
	"strings"
)

func init() {
	kernel.RegisterAPI((*Pool)(nil), &pool{})
}

type Pool interface {
	GetMQ(string) *MQ
}

type pool struct {
	Brokers *map[string]*MQ `kernel:"config,amqp"`
	Worker  task.Queue      `kernel:"worker"`
}

func (p *pool) GetMQ(n string) *MQ {
	return (*p.Brokers)[n]
}

func (p *pool) Start() error {
	s := strings.SplitN(common.Version, " ", 2)
	appName := strings.Join([]string{"piweather.center", s[0]}, " ")
	appVersion := strings.Trim(s[1], "()")

	for name, mq := range *p.Brokers {
		if mq.Url == "" || !strings.HasPrefix(mq.Url, "amqp://") {
			return fmt.Errorf("AMQP:%s:amqp broker url required", name)
		}

		mq.name = name

		if mq.ConnectionName == "" {
			mq.ConnectionName = name
		}
		if mq.Product == "" {
			mq.Product, mq.Version = appName, appVersion
		}
		if mq.Version == "" {
			mq.Version = appVersion
		}
	}

	return nil
}
