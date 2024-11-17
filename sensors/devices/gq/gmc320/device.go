package gmc320

import (
	"github.com/peter-mount/piweather.center/sensors/bus"
	device2 "github.com/peter-mount/piweather.center/sensors/device"
	"go.bug.st/serial"
)

func init() {
	device2.RegisterDevice(&device{})
}

type device struct{}

func (d *device) Info() device2.DeviceInfo {
	return device2.DeviceInfo{
		ID:           "GMC320",
		Manufacturer: "GQ Electronics LLC",
		Model:        "GMC320",
		Description:  "Geiger Counter",
		BusType:      bus.BusSerial,
		PollMode:     bus.PushReading,
	}
}

func (d *device) NewInstance(portName string, mode *serial.Mode) device2.Instance {
	return &gmc320{
		BasicSerialDevice: device2.NewBasicSerialDevice(d, portName, mode),
	}
}
