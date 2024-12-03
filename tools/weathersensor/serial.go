//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package weathersensor

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/device"
	"go.bug.st/serial"
)

func (s *Service) serialSensor(v station.Visitor[*state], sensor *station.Serial) error {
	st := v.Get()

	dev, err := device.LookupSerialDevice(st.sensor.Device)
	if err != nil {
		return errors.Errorf(sensor.Pos, "device %q for %q not found", st.sensor.Device, st.sensor.Target)
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
		if st.sensor.Poll == nil || st.sensor.Poll.Definition == "" {
			return errors.Errorf(sensor.Pos, "serial device %q requires poll period defining", st.sensor.Device)
		}

		err = s.PollDevice(dev, instance, publisher, st.sensor.Poll.Definition)

	case bus.PushReading:
		s.RunDevice(dev, instance, publisher)
	}

	return err
}
