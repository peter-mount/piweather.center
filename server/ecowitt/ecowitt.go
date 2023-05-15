package ecowitt

import (
	"context"
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/util/graphite"
	"github.com/peter-mount/piweather.center/util/mq"
	"strings"
	"time"
)

type Server struct {
	MQ       *mq.MQ            `kernel:"inject"`
	Queue    *mq.Queue         `kernel:"config,ecowitt2mqttQueue"`
	Graphite graphite.Graphite `kernel:"inject"`
}

func (s *Server) Start() error {

	err := s.MQ.ConsumeTask(s.Queue, "graphite", mq.Guard(s.consume))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) consume(ctx context.Context) error {
	body := mq.Delivery(ctx)

	data := make(map[string]interface{})
	err := json.Unmarshal(body.Body, &data)
	if err != nil {
		log.Println(err)
		return err
	}

	// Timestamp in UTC of message
	ts, ok := data["dateutc"].(string)
	if !ok {
		return nil
	}
	// The old format was "YYYY-MM-DD HH:MM:SS" whilst the new format is correct
	// So convert the old to new format for timestamp
	ts = strings.ReplaceAll(ts, " ", "T")
	if !strings.HasSuffix(ts, "Z") {
		ts = ts + "Z"
	}

	t, err := time.Parse("2006-01-02T15:04:05Z", ts)
	if err != nil {
		return err
	}

	return s.submitReadings(t, body.RoutingKey, data)
}

func (s *Server) submitReadings(t time.Time, routingKey string, data map[string]interface{}) error {
	for k, v := range data {
		err := s.submitReading(t, routingKey, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) submitReading(t time.Time, routingKey, k string, v interface{}) error {
	// Convert a bool type to a 1 or 0 numeric type
	if b, ok := v.(bool); ok {
		if b {
			v = 1
		} else {
			v = 0
		}
	}

	return s.Graphite.Publish(t, routingKey+"."+k, v)
}
