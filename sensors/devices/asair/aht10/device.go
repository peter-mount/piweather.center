package aht10

import (
	"github.com/peter-mount/piweather.center/sensors/bus"
	device2 "github.com/peter-mount/piweather.center/sensors/device"
)

func init() {
	device2.RegisterDevice(&device{})
}

// device implementation for SEN0575 over I2C
type device struct {
}

func (d *device) Info() device2.DeviceInfo {
	return device2.DeviceInfo{
		ID:           "AHT10",
		Manufacturer: "ASAIR",
		Model:        "AHT10",
		Description:  "Temperature & Humidity sensor",
		BusType:      bus.BusI2C,
		PollMode:     bus.PollReading,
	}
}

func (d *device) NewInstance(bus int, addr uint8) device2.Instance {
	return &aht10{
		BasicI2CDevice: device2.NewBasicI2CDevice(d, bus, addr),
		buf:            make([]byte, 6),
	}
}
