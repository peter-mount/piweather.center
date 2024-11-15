package sen0575

import (
	"github.com/peter-mount/piweather.center/sensors"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

func init() {
	sensors.RegisterDevice(&device{})
}

// device implementation for SEN0575 over I2C
type device struct {
	lastReading     time.Time   // time of last reading
	lastBucketCount value.Value // total bucket tips from last reading
}

func (d *device) Info() sensors.DeviceInfo {
	return sensors.DeviceInfo{
		ID:           "Sen0575",
		Manufacturer: "DFRobot",
		Model:        "SEN0575",
		Description:  "Rain Fall Detector",
		BusType:      sensors.BusI2C,
		PollMode:     sensors.PollReading,
	}
}

func (d *device) NewInstance(bus int, addr uint8) sensors.Instance {
	return &sen0575{
		BasicI2CDevice: sensors.NewBasicI2CDevice(d, bus, addr),
	}
}
