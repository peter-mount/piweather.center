package weathersensor

import (
	"context"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	station2 "github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/util/config"
	"path/filepath"
	"time"
)

type Service struct {
	Daemon         *kernel.Daemon        `kernel:"inject"`
	Cron           *cron.CronService     `kernel:"inject"`
	Config         config.Manager        `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Stations       *station.Stations     `kernel:"inject"`
	dashDir        string
}

const (
	dashDir    = "stations"
	fileSuffix = ".sensor"
)

func (s *Service) Start() error {
	s.dashDir = filepath.Join(s.Config.EtcDir(), dashDir)

	// Load existing dashboards
	stations, err := s.Stations.LoadDirectory(s.dashDir, fileSuffix, station.SensorOption)
	if err != nil {
		return err
	}

	if err := station2.NewBuilder[*state]().
		I2C(s.i2cSensor).
		Sensor(s.sensor).
		Serial(s.serialSensor).
		Station(s.station).
		Build().
		Set(&state{service: s}).
		Stations(stations); err != nil {
		return err
	}

	s.Daemon.SetDaemon()

	log.Println(version.Version)
	return nil
}

type state struct {
	service *Service
	station *station2.Station
	sensor  *station2.Sensor
}

func (s *Service) station(v station2.Visitor[*state], d *station2.Station) error {
	v.Get().station = d
	return nil
}

func (s *Service) sensor(v station2.Visitor[*state], d *station2.Sensor) error {
	v.Get().sensor = d
	return nil
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
