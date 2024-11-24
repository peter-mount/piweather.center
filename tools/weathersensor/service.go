package weathersensor

import (
	"context"
	"flag"
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/sensors"
	"github.com/peter-mount/piweather.center/config/util"
	sensors2 "github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/store/broker"
	"time"
)

type Service struct {
	ListDevices    *bool                 `kernel:"flag,list-devices,List Devices"`
	Daemon         *kernel.Daemon        `kernel:"inject"`
	Cron           *cron.CronService     `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
}

func (s *Service) Start() error {
	if *s.ListDevices {
		s.Daemon.ClearDaemon()
		return s.listDevices()
	}

	var src *sensors2.Sensors
	for _, arg := range flag.Args() {

		b, err := sensors.NewParser().
			ParseFile(arg)
		if err == nil {
			src, err = src.Merge(b)
		}
		if err != nil {
			return err
		}
	}

	if err := sensors.NewBuilder[any]().
		Sensor(s.sensor).
		Build().
		Sensors(src); err != nil {
		return err
	}

	s.Daemon.SetDaemon()
	return nil
}

func (s *Service) sensor(v sensors2.SensorVisitor[any], sensor *sensors2.Sensor) error {
	var err error
	switch {
	case sensor.I2C != nil:
		err = s.i2cSensor(v, sensor)
	case sensor.Serial != nil:
		err = s.serialSensor(v, sensor)
	default:
		err = participle.Errorf(sensor.Pos, "invalid device bus for %q", sensor.ID)
	}
	if err == nil {
		err = util.VisitorStop
	}
	return err
}

// PollDevice will configure a task that will poll the given instance based on a cron definition.
// Any errors returned by the device when it's polled will be reported in the log.
func (s *Service) PollDevice(dev device.Device, instance device.Instance, publisher publisher.Publisher, cronDef string) error {
	_, err := s.Cron.AddTask(cronDef, func(_ context.Context) error {
		err := instance.RunDevice(publisher)
		if err != nil {
			log.Printf("device %q error %s",
				dev.Info().ID,
				err.Error())
		}
		return nil
	})
	return err
}

// RunDevice will call the instance in a separate goroutine.
// Any error returned by the device will be logged, and it will retry the device after a short delay.
func (s *Service) RunDevice(dev device.Device, instance device.Instance, publisher publisher.Publisher) {
	go func() {
		for {
			err := instance.RunDevice(publisher)
			if err != nil {
				log.Printf("device %q error %s",
					dev.Info().ID,
					err.Error())
				time.Sleep(time.Second)
			}
		}
	}()
}
