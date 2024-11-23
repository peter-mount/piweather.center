//go:build !(aix || plan9 || solaris || windows)

package weathersensor

import (
	"github.com/peter-mount/go-script/errors"
	sensors2 "github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/device"
)

func (s *Service) i2cSensor(_ sensors2.SensorVisitor[any], sensor *sensors2.Sensor) error {
	dev, err := device.LookupI2CDevice(sensor.Device)
	if err != nil {
		return errors.Errorf(sensor.Pos, "device %q for %q not found", sensor.Device, sensor.ID)
	}

	instance := dev.NewInstance(sensor.I2C.Bus, uint8(sensor.I2C.Device))

	err = instance.Init()
	if err != nil {
		return errors.Errorf(sensor.Pos, "failed to init instance %q: %v", sensor.ID, err)
	}

	publisher := s.publisher(sensor)

	switch dev.Info().PollMode {
	case bus.PollReading:
		if sensor.Poll == nil || sensor.Poll.Definition == "" {
			return errors.Errorf(sensor.Pos, "i2c device %q requires poll period defining", sensor.Device)
		}

		err = s.PollDevice(dev, instance, publisher, sensor.Poll.Definition)

	case bus.PushReading:
		return errors.Errorf(sensor.Pos, "push readings not supported for i2c")
	}

	return err
}
