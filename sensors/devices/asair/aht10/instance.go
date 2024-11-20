package aht10

import (
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	device2 "github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"time"
)

const (
	// delay after commands to allow the aht10 to perform actions.
	// for reads the datasheet says over 75ms
	delay = 100 * time.Millisecond

	// Initialize device
	initDevice = 0xE1

	// Trigger measurement
	triggerMeasurement = 0xAC
)

type aht10 struct {
	device2.BasicI2CDevice
	buf []byte
}

func (s *aht10) ReadSensor() (*reading.Reading, error) {
	ret := s.NewReading()
	return ret, s.readSensor(ret)
}

func (s *aht10) RunDevice(p publisher.Publisher) error {
	rec, err := s.ReadSensor()
	if err == nil {
		err = p.Do(rec)
	}
	return err
}

func (s *aht10) readSensor(ret *reading.Reading) error {
	return s.UseDevice(func(bus i2c.I2C) error {
		err := bus.WriteByte(triggerMeasurement)
		if err != nil {
			return err
		}
		time.Sleep(delay)

		n, err := bus.Read(s.buf)
		if err != nil {
			return err
		}
		if n == len(s.buf) {
			curTemp := float64((int(s.buf[3]&0xf)<<16)|(int(s.buf[4])<<8)|int(s.buf[5]))*200.0/(1<<20) - 50
			curHumi := float64((int(s.buf[1])<<12)|(int(s.buf[2])<<4)|(int(s.buf[3]&0xf0)>>4)) * 100.0 / (1 << 20)

			ret.SetFloat64("temp", measurement.Celsius, curTemp)
			ret.SetFloat64("humidity", measurement.RelativeHumidity, curHumi)
		}

		return nil
	})
}
