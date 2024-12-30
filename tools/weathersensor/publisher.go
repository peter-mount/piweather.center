package weathersensor

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/publisher"
)

func (s *Service) publisher(stationId, sensorId string, sensor *station.Sensor) publisher.Publisher {

	pubBuilder := publisher.NewBuilder().
		SetId(stationId + "." + sensorId).
		FilterEmpty()

	for _, p := range sensor.Publisher {
		switch {
		case p.Log:
			pubBuilder = pubBuilder.Log()
		case p.DB:
			pubBuilder = pubBuilder.DB(s.DatabaseBroker)
		}
	}

	return pubBuilder.Build()
}
