package bh1750

import (
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	device2 "github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

const (
	// delay after commands to allow the device to perform actions.
	// for reads the datasheet says 120ms for H & H2 modes, 16ms for L mode
	delay = 200 * time.Millisecond

	// H mode, 1lx resolution
	hModeOneTime = byte(0x20)
	// H2 mode, 0.5lx resolution
	hMode2OneTime = byte(0x21)

	// lux value which, if less than this we switch to H2 mode
	lowLightLimit = 100.0
)

type bh1750 struct {
	device2.BasicI2CDevice
	buf []byte
	// If true then use H2 mode which has 0.5lx resolution. Used in low light levels
	mode bool
}

func (s *bh1750) ReadSensor() (*reading.Reading, error) {
	ret := s.NewReading()
	return ret, s.readSensor(ret)
}

func (s *bh1750) RunDevice(p publisher.Publisher) error {
	rec, err := s.ReadSensor()
	if err == nil {
		err = p.Do(rec)
	}
	return err
}

func (s *bh1750) readSensor(ret *reading.Reading) error {
	return s.UseDevice(func(bus i2c.I2C) error {

		lux, err := s.readLight(bus)
		if err != nil {
			return err
		}

		if value.IsPositive(lux) {
			ret.SetFloat64("lux", measurement.Lux, lux)
		}

		return nil
	})
}

func (s *bh1750) readLight(bus i2c.I2C) (float64, error) {
	mode := hModeOneTime
	if s.mode {
		mode = hMode2OneTime
	}
	err := bus.WriteByte(mode)
	if err != nil {
		return 0.0, err
	}
	time.Sleep(delay)

	n, err := bus.Read(s.buf)
	if err != nil {
		return 0.0, err
	}
	if n == len(s.buf) {
		lux := float64((int(s.buf[0])<<8)|(int(s.buf[1]))) / 1.2

		// Switch to H2 mode (0.5lx resolution) when light levels are low
		s.mode = value.LessThanEqual(lux, lowLightLimit)

		return lux, nil
	}

	return 0.0, nil
}
