package station

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/peter-mount/piweather.center/util/unit"
	amqp "github.com/rabbitmq/amqp091-go"
	"gopkg.in/yaml.v2"
	"strings"
	"time"
)

type Payload struct {
	time time.Time
	msg  []byte
	data map[string]interface{}
}

func (p *Payload) Time() time.Time {
	return p.time
}

func (p *Payload) Msg() []byte {
	return p.msg
}

func (p *Payload) Data() map[string]interface{} {
	return p.data
}

func (p *Payload) Get(path string) (interface{}, bool) {
	m := p.data
	keys := strings.Split(path, ".")
	l := len(keys) - 1
	for i, k := range keys {
		v, ok := m[k]
		if !ok {
			return nil, false
		}
		if i == l {
			return v, true
		}

		if nm, ok := v.(map[string]interface{}); ok {
			m = nm
		} else {
			return nil, false
		}
	}
	return nil, false
}

func (s *Sensors) FromAMQP(msg amqp.Delivery) (*Payload, error) {
	return s.FromBytes(msg.Body)
}

func (s *Sensors) FromBytes(msg []byte) (*Payload, error) {
	// Payload needs a copy of msg in case the provider reuses that slice.
	// Time defaults to now in UTC before being overridden by the payload
	// as that allows for those messages without a time
	p := &Payload{
		time: time.Now().UTC(),
		msg:  make([]byte, len(msg)),
		data: make(map[string]interface{}),
	}
	copy(p.msg, msg)

	var err error
	switch s.Format {
	case "json", "JSON", "":
		err = json.Unmarshal(p.msg, &p.data)

	case "xml", "XML":
		err = xml.Unmarshal(p.msg, &p.data)

	case "yaml", "yml", "YAML", "YML":
		err = yaml.Unmarshal(p.msg, &p.data)

	default:
		err = fmt.Errorf("unsupported format %q", s.Format)
	}
	if err != nil {
		return nil, err
	}

	if s.Timestamp != "" {
		if ts, ok := p.Get(s.Timestamp); ok {
			if st, ok := ts.(string); ok {
				if t := unit.ParseTime(st); !t.IsZero() {
					p.time = t
				}
			}
		}
	}

	return p, nil
}
