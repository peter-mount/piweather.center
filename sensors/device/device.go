package device

import (
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"time"
)

// Device defines the functions all devices have to implement
type Device interface {
	// Info returns the DeviceInfo for this device
	Info() DeviceInfo
}

// DeviceInfo holds metadata about the device.
type DeviceInfo struct {
	// ID of the device
	ID string
	// Manufacturer of the device
	Manufacturer string
	// Model of the device
	Model string
	// Description of what the device does
	Description string
	// BusType of the device
	BusType bus.BusType
	// PollMode of the device
	PollMode bus.PollMode
}

func newReading(dev Device) *reading.Reading {
	return &reading.Reading{
		Time:   time.Now().UTC(), // default to now but usually the device will override it
		Device: dev.Info().Model,
	}
}
