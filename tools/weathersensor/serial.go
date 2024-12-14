//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package weathersensor

import (
	"fmt"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/device"
	"go.bug.st/serial"
)

func (s *Service) serialSensor(v station.Visitor[*state], sensor *station.Serial) error {
	st := v.Get()

	dev, err := device.LookupSerialDevice(sensor.Driver)
	if err != nil {
		return errors.Errorf(sensor.Pos, "device %q for %q not found", sensor.Driver, st.sensor.Target)
	}

	mode := &serial.Mode{
		BaudRate: sensor.BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	instance := dev.NewInstance(sensor.Port, mode)

	publisher := s.publisher(st.sensor)

	switch dev.Info().PollMode {
	case bus.PollReading:
		err = s.PollDevice(dev, instance, publisher, st.sensor.Poll.Definition())

		s.addSensor(
			"serial",
			st.station.Name,
			st.sensor.Target.OriginalName,
			"poll "+st.sensor.Poll.Definition(),
			sensor.Port,
			fmt.Sprintf("%d", sensor.BaudRate))

	case bus.PushReading:
		s.RunDevice(dev, instance, publisher)

		s.addSensor(
			"serial",
			st.station.Name,
			st.sensor.Target.OriginalName,
			"push",
			sensor.Port,
			fmt.Sprintf("%d", sensor.BaudRate))
	}

	if err == nil {
		s.sensorCount++
	}

	return errors.Error(sensor.Pos, err)
}
