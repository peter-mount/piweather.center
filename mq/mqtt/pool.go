package mqtt

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/util/config"
)

func init() {
	kernel.RegisterAPI((*Pool)(nil), &pool{})
}

type Pool interface {
	GetMQ(string) *MQ
}

type pool struct {
	ConfigManager config.Manager `kernel:"inject"`
	Worker        task.Queue     `kernel:"worker"`
	Brokers       *map[string]*MQ
}

func (p *pool) GetMQ(n string) *MQ {
	return (*p.Brokers)[n]
}

func (p *pool) loadConfig() error {
	m := make(map[string]*MQ)
	p.Brokers = &m
	return p.ConfigManager.ReadYamlOptional("mqtt.yaml", p.Brokers)
}

func (p *pool) Start() error {
	if err := p.loadConfig(); err != nil {
		return err
	}

	for name, mq := range *p.Brokers {
		if mq.Url == "" {
			return fmt.Errorf("MQTT:%s:mqtt broker url required", name)
		}

		mq.name = name

		if err := mq.connect(); err != nil {
			return err
		}
	}

	return nil
}
