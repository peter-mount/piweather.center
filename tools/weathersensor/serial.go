//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package weathersensor

import (
	"context"
	"github.com/peter-mount/go-script/errors"
	sensors2 "github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/device"
	"go.bug.st/serial"
)

func (s *Service) serialSensor(_ sensors2.SensorVisitor[any], sensor *sensors2.Sensor) error {
	// Lookup device
	dev, err := device.LookupSerialDevice(sensor.Device)
	if err != nil {
		return errors.Errorf(sensor.Pos, "device %q for %q not found", sensor.Device, sensor.ID)
	}

	// create instance
	def := sensor.Serial

	mode := &serial.Mode{
		BaudRate: def.BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	switch def.DataBits {
	case 5, 6, 7, 8:
		mode.DataBits = def.DataBits
	default:
		mode.DataBits = 8
	}

	switch def.Parity {
	case "odd":
		mode.Parity = serial.OddParity
	case "even":
		mode.Parity = serial.EvenParity
	case "none", "no":
		mode.Parity = serial.NoParity
	default:
		mode.Parity = serial.NoParity
	}

	switch def.StopBits {
	case "1":
		mode.StopBits = serial.OneStopBit
	case "1.5":
		mode.StopBits = serial.OnePointFiveStopBits
	case "2":
		mode.StopBits = serial.TwoStopBits
	}

	instance := dev.NewInstance(def.Port, mode)

	publisher := s.publisher(sensor)

	switch dev.Info().PollMode {
	case bus.PollReading:
		if sensor.Poll == nil || sensor.Poll.Definition == "" {
			return errors.Errorf(sensor.Pos, "serial device %q requires poll period defining", sensor.Device)
		}

		_, err = s.Cron.AddTask(sensor.Poll.Definition, func(_ context.Context) error {
			return instance.RunDevice(publisher)
		})

	case bus.PushReading:
		go instance.RunDevice(publisher)
	}

	return err
}
