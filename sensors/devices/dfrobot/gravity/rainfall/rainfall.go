package rainfall

import (
	smbus2 "github.com/peter-mount/piweather.center/sensors/bus/smbus"
	"time"
)

const (
	// The address of the input register for high 16-bit cumulative rainfall since working started.
	SEN0575_INPUT_REG_CUMULATIVE_RAINFALL_H = 0x0009
	// The address of the input register for raw data (low 16-bit) in memory.
	SEN0575_INPUT_REG_RAW_DATA_L = 0x000A
	// The address of the input register for raw data (high 16-bit) in memory.
	SEN0575_INPUT_REG_RAW_DATA_H = 0x000B
	// The address of the input register for system working time in memory.
	SEN0575_INPUT_REG_SYS_TIME = 0x000C
	// Set the time to calculate cumulative rainfall.
	SEN0575_HOLDING_REG_RAW_RAIN_HOUR = 0x0006
	// Set the base rainfall value.
	SEN0575_HOLDING_REG_RAW_BASE_RAINFALL = 0x0007

	// The address of the input register for PID in memory.
	SEN0575_I2C_REG_PID = 0x00
	// The address of the input register for VID in memory.
	SEN0575_I2C_REG_VID = 0x02
	// sen0575I2cRegVersion The address of the input register for firmware version in memory.
	sen0575I2cRegVersion = 0x0A
	// SEN0575 Stores the cumulative rainfall within the set time
	sen0575I2cRegTimeRainfall = 0x0C
	// sen0575I2cRegCumulativeRainfall Stores the cumulative rainfall since starting work
	sen0575I2cRegCumulativeRainfall = 0x10
	// SEN0575 Stores the low 16 bits of the raw data
	SEN0575_I2C_REG_RAW_DATA = 0x14
	// sen0575I2cRegSysTime Stores the system time
	sen0575I2cRegSysTime = 0x18

	// Sets the time for calculating the cumulative rainfall
	sen0575I2cRegRawRainHour = 0x26
	// Sets the base value of accumulated rainfall
	SEN0575_I2C_REG_RAW_BASE_RAINFALL = 0x28

	I2C_MODE  = 0
	UART_MODE = 1

	// i2cAddr of the device
	i2cAddr = 0x1d
)

type DFRGravityRainFall struct {
	manager *smbus2.Manager `kernel:"inject"`
	device  smbus2.SMBus
	version uint16
}

func (s *DFRGravityRainFall) Start() error {
	return s.manager.UseDevice(1, i2cAddr, func(bus smbus2.SMBus) error {
		ver, err := s.GetFirmwareVersion(bus)
		if err == nil {
			s.version = ver
		}
		return err
	})
}

func (s *DFRGravityRainFall) ReadSensor() (interface{}, error) {
	rec := newRainFall(s.version)

	err := s.manager.UseDevice(1, i2cAddr, func(bus smbus2.SMBus) error {
		uptime, err := s.GetSensorWorkingTime(bus)
		if err == nil {
			rec.Device.Uptime = int(uptime.Seconds())

			rec.Record.Total, err = s.GetCumulativeRainFall(bus)
		}

		if err == nil {
			rec.Record.Hour, err = s.GetRainFall(bus, 1)
		}

		if err == nil {
			rec.Record.Day, err = s.GetRainFall(bus, 24)
		}

		if err == nil {
			rec.Record.BucketCount, err = s.GetBucketCount(bus)
		}

		return err
	})

	return rec, err
}

// GetFirmwareVersion returns the firmware version of the SEN0575
func (s *DFRGravityRainFall) GetFirmwareVersion(bus smbus2.SMBus) (uint16, error) {
	return bus.ReadRegisterUint16(sen0575I2cRegVersion)
}

// GetSensorWorkingTime returns the uptime of the SEN0575
func (s *DFRGravityRainFall) GetSensorWorkingTime(bus smbus2.SMBus) (time.Duration, error) {
	ret, err := bus.ReadRegisterUint16(sen0575I2cRegSysTime)
	if err != nil {
		return 0, err
	}

	return time.Minute * time.Duration(ret), nil
}

func (s *DFRGravityRainFall) GetCumulativeRainFall(bus smbus2.SMBus) (float64, error) {
	ret, err := bus.ReadRegisterUint32(sen0575I2cRegCumulativeRainfall)
	if err != nil {
		return 0, err
	}

	return float64(ret) / 10000.0, nil
}

func (s *DFRGravityRainFall) GetRainFall(bus smbus2.SMBus, hours uint8) (float64, error) {
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

func (s *DFRGravityRainFall) GetBucketCount(bus smbus2.SMBus) (uint32, error) {
	return bus.ReadRegisterUint32(SEN0575_I2C_REG_RAW_DATA)
}
