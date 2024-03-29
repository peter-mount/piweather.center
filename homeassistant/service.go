package homeassistant

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/util/config"
	mq "github.com/peter-mount/piweather.center/util/mq/amqp"
	"github.com/peter-mount/piweather.center/util/mq/mqtt"
	"path"
	"strings"
)

func init() {
	kernel.RegisterAPI((*Service)(nil), &service{})
}

type Service interface {
	StoreReading(ctx context.Context) error
}

// Service implements the HomeAssistant Integration
type service struct {
	ConfigManager config.Manager `kernel:"inject"`
	Amqp          mq.Pool        `kernel:"inject"`
	Mqtt          mqtt.Pool      `kernel:"inject"`
	Config        *HomeAssistant
}

func (s *service) Stop() {
	s.Config.close()
}

func (s *service) Start() error {
	ha := &HomeAssistant{}

	if err := s.ConfigManager.ReadYamlOptional("homeassistant.yaml", ha); err != nil {
		return err
	}

	// If disabled do nothing
	if ha.Disabled {
		return nil
	}

	switch {
	case ha.Amqp != "" && ha.AmqpPublisher != nil:
		broker := s.Amqp.GetMQ(ha.Amqp)
		if broker == nil {
			return fmt.Errorf("amqp broker %q undefined", ha.Amqp)
		}
		ha.amqp = broker
		if err := ha.AmqpPublisher.Bind(broker); err != nil {
			return err
		}

	case ha.Mqtt != "" && ha.MqttPublisher != nil:
		broker := s.Mqtt.GetMQ(ha.Mqtt)
		if broker == nil {
			return fmt.Errorf("mqtt broker %q undefined", ha.Mqtt)
		}
		ha.mqtt = broker
		if err := ha.MqttPublisher.Bind(broker); err != nil {
			return err
		}

	default:
		return fmt.Errorf("no message broker configured")
	}

	s.Config = ha

	return s.SendConfiguration()
}

func (s *service) SendConfiguration() error {
	for _, sensors := range s.Config.Sensors {
		for name, entity := range sensors.Entities {
			if entity.Name == "" {
				entity.Name = name
			}

			if entity.SensorType == "" {
				entity.SensorType = "sensor"
			}

			// Add the optional Device entry
			entity.Device = sensors.Device

			entity.ObjectId = strings.Join([]string{
				"piweather_center",
				sensors.NodeId,
				sensors.ObjectIdPrefix,
				entity.Name,
			}, "-")

			entity.UniqueID = strings.Join([]string{
				"piweather_center",
				sensors.NodeId,
				sensors.ObjectIdPrefix,
				entity.Name,
			}, "-")

			topicPrefix := path.Join(
				s.Config.DiscoveryPrefix,
				entity.SensorType,
				sensors.NodeId,
				entity.Name,
			)
			configTopic := path.Join(topicPrefix, "config")

			if err := s.publish(configTopic, entity); err != nil {
				return fmt.Errorf("failed to publish %q: %v", configTopic, err)
			}
		}
	}

	return nil
}

func (s *service) publish(topic string, msg any) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if log.IsVerbose() {
		log.Println(string(b))
	}

	switch {
	case s.Config.AmqpPublisher != nil:
		return s.Config.AmqpPublisher.Publish(
			strings.ReplaceAll(topic, "/", "."),
			b,
		)

	default:
		return nil
	}
}
