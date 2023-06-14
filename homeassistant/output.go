package homeassistant

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	store2 "github.com/peter-mount/piweather.center/store"
	"path/filepath"
	"strings"
)

func (s *service) StoreReading(ctx context.Context) error {
	if s.Config != nil {
		return s.storeReading(ctx)
	}
	return nil
}

func (s *service) storeReading(ctx context.Context) error {

	store := store2.StoreFromContext(ctx)
	if store == nil {
		return nil
	}

	// map of state_topic maps to generate
	topics := make(map[string]map[string]interface{})

	//values := value.MapFromContext(ctx)

	for _, sensor := range s.Config.Sensors {
		for _, entity := range sensor.Entities {
			if entity.SensorId != "" && entity.StateTopic != "" {
				reading := store.GetReading(entity.SensorId)
				if reading != nil && reading.Value.IsValid() {

					// TODO add conversion here

					// Store using last value past .
					topic := topics[entity.StateTopic]
					if topic == nil {
						topic = make(map[string]interface{})
					}

					field := strings.ReplaceAll(filepath.Ext(entity.SensorId), ".", "")
					field = strings.ToLower(field)
					topic[field] = reading.Value.Float()

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
