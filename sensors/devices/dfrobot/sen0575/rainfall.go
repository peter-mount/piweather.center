package rainfall

import (
	"context"
	"fmt"
	"github.com/peter-mount/piweather.center/sensors"
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	"time"
)

const (
	// The address of the input register for PID in memory.
	sen0575I2cRegPid = 0x00
	// The address of the input register for VID in memory.
	sen0575I2cRegVid = 0x02
	// sen0575I2cRegVersion The address of the input register for firmware version in memory.
	sen0575I2cRegVersion = 0x0A
	// SEN0575 Stores the cumulative rainfall within the set time
	sen0575I2cRegTimeRainfall = 0x0C
	// sen0575I2cRegCumulativeRainfall Stores the cumulative rainfall since starting work
	sen0575I2cRegCumulativeRainfall = 0x10
	// SEN0575 Stores the low 16 bits of the raw data
	sen0575I2cRegRawData = 0x14
	// sen0575I2cRegSysTime Stores the system time
	sen0575I2cRegSysTime = 0x18

	// Sets the time for calculating the cumulative rainfall
	sen0575I2cRegRawRainHour = 0x26
	// Sets the base value of accumulated rainfall
	sen0575I2cRegRawBaseRainfall = 0x28

	// i2cAddr of the device
	i2cAddr = 0x1d
)

type Sen0575 struct {
	lastReading     time.Time // time of last reading
	lastBucketCount uint32    // total bucket tips from last reading
}

func (s *Sen0575) task(ctx context.Context) error {
	return nil
}

func (s *Sen0575) Start() error {
	/*	return i2c.UseI2CConcurrent(1, i2cAddr, func(c i2c.I2C) error {
			ver, err := s.GetFirmwareVersion(c)
			if err == nil {
				s.version = ver
			}

			// Use to recalibrate the bucket
			//rr = s.SetRainAccumulatedValue(c, 0.2794)

			return err
		})
	*/

	// Take initial reading so we can extract the raw tip count between subsequent readings
	_, err := s.ReadSensor()
	return err
}

func (s *Sen0575) ReadSensor() (interface{}, error) {
	//rec := newRainFall(s.version)
	rec := Record{}
	ret := sensors.NewReading(sensors.LookupDevice("sen0575"))
	ret.Readings = &rec

	err := i2c.UseI2CConcurrent(1, i2cAddr, func(bus i2c.I2C) error {
		var err error

		now := time.Now()

		//rec.Device.Uptime, err = s.GetSensorWorkingTime(bus)
		//if err == nil {
		rec.Total, err = s.GetCumulativeRainFall(bus)
		//}

		if err == nil {
			rec.Hour, err = s.GetRainFall(bus, 1)
		}

		if err == nil {
			rec.Day, err = s.GetRainFall(bus, 24)
		}

		if err == nil {
			rec.BucketCount, err = s.GetBucketCount(bus)
		}

		if err == nil {
			// Duration since last reading, accounting for initialisation
			if !s.lastReading.IsZero() {
				rec.Duration = uint32(now.Sub(s.lastReading).Seconds())
			}
			s.lastReading = now

			// Tips since last reading.
			rec.Tips = rec.BucketCount - s.lastBucketCount
			s.lastBucketCount = rec.BucketCount
		}

		return err
	})

	return ret, err
}

// GetFirmwareVersion returns the firmware version of the SEN0575
func (s *Sen0575) GetFirmwareVersion(bus i2c.I2C) (string, error) {
	v, err := bus.ReadRegisterUint16(sen0575I2cRegVersion)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d.%d.%d.%d",
		(v>>12)&0x0f,
		(v>>8)&0x0f,
		(v>>4)&0x0f,
		v&0x0f,
	), nil
}

// GetSensorWorkingTime returns the uptime of the SEN0575
func (s *Sen0575) GetSensorWorkingTime(bus i2c.I2C) (uint16, error) {
	return bus.ReadRegisterUint16(sen0575I2cRegSysTime)
}

func (s *Sen0575) GetCumulativeRainFall(bus i2c.I2C) (float64, error) {
	ret, err := bus.ReadRegisterUint32(sen0575I2cRegCumulativeRainfall)
	if err != nil {
		return 0, err
	}

	return float64(ret) / 10000.0, nil
}

func (s *Sen0575) GetRainFall(bus i2c.I2C, hours uint8) (float64, error) {
	if hours < 1 || hours > 24 {
		return 0, hourRangeError
	}

	err := bus.WriteRegisterUint8(sen0575I2cRegRawRainHour, hours)
	if err != nil {
		return 0, err
	}

	ret, err := bus.ReadRegisterUint32(sen0575I2cRegTimeRainfall)
	if err != nil {
		return 0, err
	}

	return float64(ret) / 10000.0, nil
}

func (s *Sen0575) GetBucketCount(bus i2c.I2C) (uint32, error) {
	return bus.ReadRegisterUint32(sen0575I2cRegRawData)
}

// SetRainAccumulatedValue sets the factor to multiply bucket count to get the required
// rain values in mm.
//
// The spec states the bucket's resolution is 0.28mm, but I found that one unit uses 0.274.
// So if the readings are incorrect then use: err = s.SetRainAccumulatedValue(c, 0.2794)
// to reset it.
//
// You can also use this to reset it if using a new bucket.
func (s *Sen0575) SetRainAccumulatedValue(bus i2c.I2C, val float64) error {
	return bus.WriteRegisterUint16(sen0575I2cRegRawBaseRainfall, uint16(val*10000.0))
}
