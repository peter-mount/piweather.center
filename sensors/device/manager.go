package device

import (
	"errors"
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"sync"
)

var (
	mutex          sync.Mutex
	deviceNotFound = errors.New("device not found")
)

// Instance of a device
type Instance interface {
	// Init must be called once an instance has been created to initialise the device.
	Init() error
	// NewReading returns a blank Reading for this device
	NewReading() *reading.Reading
	// ReadSensor takes measurements from the device
	// when PollMode == PollReading
	ReadSensor() (*reading.Reading, error)
	// RunDevice takes measurements from the device in realtime.
	// This is only called when PollMode == PushReading
	RunDevice(publisher.Publisher) error
}

// RegisterDevice registers a new Device handler.
// This will panic if the ID has already been registered.
// IDs are case-insensitive.
func RegisterDevice(device Device) {
	if device == nil {
		panic(errors.New("device is nil"))
	}

	info := device.Info()

	switch info.BusType {
	case bus.BusI2C:
		registerI2CDevice(device)
	case bus.BusSerial:
		registerSerialDevice(device)
	default:
		panic(fmt.Errorf("unknown bus type: %s", info.BusType.Label()))
	}
}
