package weathersensor

import (
	"github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/sensors/publisher"
)

func (s *Service) publisher(sensor *sensors.Sensor) publisher.Publisher {

	pubBuilder := publisher.NewBuilder().
		SetId(sensor.ID)

	for _, p := range sensor.Publisher {
		switch {
		case p.Log:
			pubBuilder = pubBuilder.Log()
		case p.FilterEmpty:
			pubBuilder = pubBuilder.FilterEmpty()
		}
	}

	return pubBuilder.Build()
}
