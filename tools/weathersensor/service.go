package weathersensor

import (
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/sensors/devices/dfrobot/sen0575"
	"time"
)

type Service struct {
	DFRGravityRainFall *rainfall.Sen0575 `kernel:"inject"`
	ListDevices        *bool             `kernel:"flag,list-devices,List Devices"`
}

func (s *Service) Start() error {
	if *s.ListDevices {
		return s.listDevices()
	}

	return s.testSensor()
}

func (s *Service) testSensor() error {
	for {
		rec, err := s.ReadSensor()
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

func (s *Service) ReadSensor() (interface{}, error) {
	return s.DFRGravityRainFall.ReadSensor()
}
