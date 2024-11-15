package gmc320

import (
	"github.com/peter-mount/piweather.center/sensors"
	"go.bug.st/serial"
)

func init() {
	sensors.RegisterDevice(&device{})
}

type device struct{}

func (d *device) Info() sensors.DeviceInfo {
	return sensors.DeviceInfo{
		ID:           "GMC320",
		Manufacturer: "GQ Electronics LLC",
		Model:        "GMC320",
		Description:  "Geiger Counter",
		BusType:      sensors.BusSerial,
		PollMode:     sensors.PushReading,
	}
}

func (d *device) NewInstance(portName string, mode *serial.Mode) sensors.Instance {
	return &gmc320{
		BasicSerialDevice: sensors.NewBasicSerialDevice(d, portName, mode),
	}
}
