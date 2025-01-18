package weathersensor

import (
	"fmt"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/bus/rtl433"
	"strings"
)

func (s *Service) rtl433(v station.Visitor[*state], sensor *station.Rtl433) error {
	st := v.Get()

	freq := strings.TrimSuffix(fmt.Sprintf("%.3f", sensor.Frequency), ".000")

	if freq == "0" {
		freq = "433"
	}

	s.rtl433Sensors[freq] = append(s.rtl433Sensors[freq], st.sensor)

	s.addSensor("rtl433", st.station.Name, st.sensor.Target.OriginalName, "", "", freq+"M")

	// Copy st.sensor as we need this value and st is transient
	parentSensor := st.sensor
	publisherId := st.station.Name + "." + parentSensor.Target.OriginalName
	s.httpPublisher[publisherId] = s.publisher(parentSensor)

	s.RTL433.AddListener(freq+"M", &rtl433.Listener{
		Model:   sensor.Model,
		Id:      sensor.Id,
		SubType: sensor.SubType,
		Handler: func(message *rtl433.Message) {
			_ = s.processPayload(publisherId, message.Payload, parentSensor)
		},
	})

	s.sensorCount++

	return nil
}
