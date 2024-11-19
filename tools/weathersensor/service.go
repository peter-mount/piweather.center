package weathersensor

import (
	"flag"
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/config/sensors"
	"github.com/peter-mount/piweather.center/config/util"
	sensors2 "github.com/peter-mount/piweather.center/config/util/sensors"
	_ "github.com/peter-mount/piweather.center/sensors/devices"
)

type Service struct {
	ListDevices *bool             `kernel:"flag,list-devices,List Devices"`
	Daemon      *kernel.Daemon    `kernel:"inject"`
	Cron        *cron.CronService `kernel:"inject"`
}

func (s *Service) Start() error {
	if *s.ListDevices {
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
