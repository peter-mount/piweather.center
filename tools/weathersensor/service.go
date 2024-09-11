package weathersensor

import (
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/sensors"
	_ "github.com/peter-mount/piweather.center/sensors/devices/dfrobot/sen0575"
	"time"
)

type Service struct {
	ListDevices *bool `kernel:"flag,list-devices,List Devices"`
}

func (s *Service) Start() error {
	if *s.ListDevices {
		return s.listDevices()
	}

	return s.testSensor()
}

func (s *Service) testSensor() error {
	// Lookup device
	dev, err := sensors.LookupI2CDevice("sen0575")
	if err != nil {
		return err
	}

	// create instance
	instance := dev.NewInstance(1, 0x1d)

	// Initialise the instance
	err = instance.Init()
	if err != nil {
		return err
	}

	for {
		rec, err := instance.ReadSensor()
		if err != nil {
			log.Println(err)
		} else {
			b, err := json.Marshal(&rec)
			if err != nil {
				return err
			}

			log.Println(string(b))
		}

		time.Sleep( /*5 */ time.Second)
	}
}