package weathersensor

import (
	_ "github.com/peter-mount/piweather.center/sensors/devices"
)

type Service struct {
	ListDevices *bool `kernel:"flag,list-devices,List Devices"`
}

func (s *Service) Start() error {
	if *s.ListDevices {
		return s.listDevices()
	}

	return nil //s.testSensor()
}

/*
func (s *Service) testSensor() error {
	// Lookup device
	dev, err := device.LookupSerialDevice("gmc320")
	if err != nil {
		return err
	}

	// create instance
	instance := dev.NewInstance("/dev/ttyUSB0", &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: 0,
	})

	pub := publisher.NewBuilder().
		SetId("test.office.geiger").
		FilterEmpty().
		Log().
		Build()

	return instance.RunDevice(pub)
}
*/

/*
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

		time.Sleep(  time.Second)
	}
}
*/
