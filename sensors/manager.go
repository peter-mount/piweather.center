package sensors

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	mutex   sync.Mutex
	devices map[BusType]map[string]Device
)

func init() {
	devices = make(map[BusType]map[string]Device)
	devices[BusI2C] = make(map[string]Device)
	devices[BusSPI] = make(map[string]Device)
	devices[BusSerial] = make(map[string]Device)
}

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
	BusType BusType
	// PollMode of the device
	PollMode PollMode
}

// Instance of a device
type Instance interface {
	// Init must be called once an instance has been created to initialise the device.
	Init() error
	// NewReading returns a blank Reading for this device
	NewReading() *Reading
	// ReadSensor takes measurements from the device
	ReadSensor() (*Reading, error)
}

// RegisterDevice registers a new Device handler.
// This will panic if the ID has already been registered.
// IDs are case-insensitive.
func RegisterDevice(device Device) {
	if device == nil {
		panic(errors.New("device is nil"))
	}
	mutex.Lock()
	defer mutex.Unlock()

	name := strings.ToLower(device.Info().ID)

	if name == "" {
		panic(errors.New("device name cannot be empty"))
	}

	bus := device.Info().BusType

	m := devices[bus]
	if m == nil {
		panic(fmt.Errorf("device %q has invalid BusType %d", name, bus))
	}

	if _, exists := m[name]; exists {
		panic(fmt.Errorf("%s device with name %q already exists", bus.Label(), name))
	}

	m[name] = device
}

func lookupDevice(bus BusType, name string) Device {
	mutex.Lock()
	defer mutex.Unlock()

	m := devices[bus]
	if m == nil {
		return nil
	}

	return m[strings.ToLower(name)]
}
