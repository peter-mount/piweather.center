//go:build !(aix || plan9 || solaris || windows)

package weathersensor

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/device"
)

func (s *Service) i2cSensor(v station.Visitor[*state], sensor *station.I2C) error {
	st := v.Get()

	dev, err := device.LookupI2CDevice(st.sensor.Device)
	if err != nil {
		return errors.Errorf(sensor.Pos, "device %q for %q not found", st.sensor.Device, st.sensor.Target)
	}

	instance := dev.NewInstance(sensor.Bus, uint8(sensor.Device))

	err = instance.Init()
	if err != nil {
		return errors.Errorf(sensor.Pos, "failed to init instance %q: %v", st.sensor.Target, err)
	}

	publisher := s.publisher(st.sensor)

	switch dev.Info().PollMode {
	case bus.PollReading:
		if st.sensor.Poll == nil || st.sensor.Poll.Definition == "" {
			return errors.Errorf(sensor.Pos, "i2c device %q requires poll period defining", sensor.Device)
		}

		err = s.PollDevice(dev, instance, publisher, st.sensor.Poll.Definition)

	case bus.PushReading:
		return errors.Errorf(sensor.Pos, "push readings not supported for i2c")
	}

	return err
}
