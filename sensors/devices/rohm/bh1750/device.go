package bh1750

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
		ID:           "BH1750",
		Manufacturer: "ROHM Semiconductor",
		Model:        "BH1750",
		Description:  "Ambient Light sensor",
		BusType:      bus.BusI2C,
		PollMode:     bus.PollReading,
	}
}

func (d *device) NewInstance(bus int, addr uint8) device2.Instance {
	return &bh1750{
		BasicI2CDevice: device2.NewBasicI2CDevice(d, bus, addr),
		buf:            make([]byte, 2),
	}
}
