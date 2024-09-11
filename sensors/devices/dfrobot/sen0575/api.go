package sen0575

import (
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
)

const (
	// The address of the input register for PID in memory.
	sen0575I2cRegPid = 0x00
	// The address of the input register for VID in memory.
	sen0575I2cRegVid = 0x02
	// sen0575I2cRegVersion The address of the input register for firmware version in memory.
	sen0575I2cRegVersion = 0x0A
	// sen0575 Stores the cumulative rainfall within the set time
	sen0575I2cRegTimeRainfall = 0x0C
	// sen0575I2cRegCumulativeRainfall Stores the cumulative rainfall since starting work
	sen0575I2cRegCumulativeRainfall = 0x10
	// sen0575 Stores the low 16 bits of the raw data
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

// GetFirmwareVersion returns the firmware version of the sen0575
func (s *sen0575) GetFirmwareVersion(bus i2c.I2C) (string, error) {
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

// GetSensorWorkingTime returns the uptime of the sen0575
func (s *sen0575) GetSensorWorkingTime(bus i2c.I2C) (uint16, error) {
	return bus.ReadRegisterUint16(sen0575I2cRegSysTime)
}

func (s *sen0575) GetCumulativeRainFall(bus i2c.I2C) (value.Value, error) {
	ret, err := bus.ReadRegisterUint32(sen0575I2cRegCumulativeRainfall)
	if err != nil {
		return measurement.MilliMeters.Value(0), err
	}

	return measurement.MilliMeters.Value(float64(ret) / 10000.0), nil
}

func (s *sen0575) GetRainFall(bus i2c.I2C, hours uint8) (value.Value, error) {
	if hours < 1 || hours > 24 {
		return measurement.MilliMeters.Value(0), hourRangeError
	}

	err := bus.WriteRegisterUint8(sen0575I2cRegRawRainHour, hours)
	if err != nil {
		return measurement.MilliMeters.Value(0), err
	}

	ret, err := bus.ReadRegisterUint32(sen0575I2cRegTimeRainfall)
	if err != nil {
		return measurement.MilliMeters.Value(0), err
	}

	return measurement.MilliMeters.Value(float64(ret) / 10000.0), nil
}

func (s *sen0575) GetBucketCount(bus i2c.I2C) (value.Value, error) {
	c, err := bus.ReadRegisterUint32(sen0575I2cRegRawData)
	return value.Integer.Value(float64(c)), err
}

// SetRainAccumulatedValue sets the factor to multiply bucket count to get the required
// rain values in mm.
//
// The spec states the bucket's resolution is 0.28mm, but I found that one unit uses 0.274.
// So if the readings are incorrect then use: err = s.SetRainAccumulatedValue(c, 0.2794)
// to reset it.
//
// You can also use this to reset it if using a new bucket.
func (s *sen0575) SetRainAccumulatedValue(bus i2c.I2C, val float64) error {
	return bus.WriteRegisterUint16(sen0575I2cRegRawBaseRainfall, uint16(val*10000.0))
}
