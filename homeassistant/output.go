package homeassistant

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/weather/value"
	"path/filepath"
	"strings"
)

func (s *service) StoreReading(ctx context.Context) error {
	if s.Config != nil {
		return s.storeReading(ctx)
	}
	log.Println("***SKIP***")
	return nil
}

func (s *service) storeReading(ctx context.Context) error {

	// map of state_topic maps to generate
	topics := make(map[string]map[string]interface{})

	values := value.MapFromContext(ctx)

	for _, sensor := range s.Config.Sensors {
		for _, entity := range sensor.Entities {
			if entity.SensorId != "" && entity.StateTopic != "" {
				val := values.Get(entity.SensorId)
				if val.IsValid() {
					// TODO add conversion here

					// Store using last value past .
					topic := topics[entity.StateTopic]
					if topic == nil {
						topic = make(map[string]interface{})
					}

					field := strings.ReplaceAll(filepath.Ext(entity.SensorId), ".", "")
					field = strings.ToLower(field)
					topic[field] = val.Float()

					topics[entity.StateTopic] = topic
				}

			}
		}
	}

	for topic, values := range topics {
		if len(values) > 0 {
			err := s.publish(topic, values)
			if err != nil {
				log.Printf("Error publishing to %q: %v", topic, err)
			}
		}
	}

	return nil
}
