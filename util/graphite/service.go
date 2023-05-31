package graphite

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	mq2 "github.com/peter-mount/piweather.center/mq/amqp"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Graphite handles receiving events from rabbit and logging the responses to graphite
// via RabbitMQ.
type Graphite interface {
	Publish(t time.Time, k string, v interface{}) error
}

func init() {
	kernel.RegisterAPI((*Graphite)(nil), &graphite{})
}

type graphite struct {
	MQ        *mq2.MQ        `kernel:"inject"`
	Publisher *mq2.Publisher `kernel:"config,graphitePublisher"`
}

func (m *graphite) Start() error {
	err := m.MQ.AttachPublisher(m.Publisher)
	return err
}

func (m *graphite) Publish(t time.Time, k string, v interface{}) error {
	var val string

	if v != nil {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int:
			val = strconv.Itoa(v.(int))
		case reflect.Float64:
			val = fmt.Sprintf("%.3f", v.(float64))
		case reflect.String:
			switch strings.ToLower(v.(string)) {
			case "on":
				val = "1"
			case "off":
				val = "0"
			case "true":
				val = "1"
			case "false":
				val = "0"
			}
		case reflect.Bool:
			if v.(bool) {
				val = "1"
			} else {
				val = "0"
			}
		}
	}

	if val != "" {
		ts := t.UTC()
		key := m.Publisher.EncodeKey(k)
		msg := fmt.Sprintf("%s %s %d", key, val, ts.Unix())

		return m.Publisher.Post(key, []byte(msg), nil, ts)
	}

	return nil
}
