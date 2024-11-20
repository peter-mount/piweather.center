package device

import (
	"errors"
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"strings"
	"sync"
)

var (
	mutex          sync.Mutex
	devices        map[bus.BusType]map[string]Device
	deviceNotFound = errors.New("device not found")
)

func init() {
	devices = make(map[bus.BusType]map[string]Device)
	devices[bus.BusI2C] = make(map[string]Device)
	devices[bus.BusSPI] = make(map[string]Device)
	devices[bus.BusSerial] = make(map[string]Device)
}

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
	//if info.BusType == busUndefined {
	//	panic(fmt.Errorf("bus type undefined for %q", info.ID))
	//}
	//if info.PollMode == pollUndefined {
	//	panic(fmt.Errorf("poll mode undefined for %q", info.ID))
	//}

	name := strings.ToLower(info.ID)
	if name == "" {
		panic(errors.New("device name cannot be empty"))
	}

	mutex.Lock()
	defer mutex.Unlock()

	busType := info.BusType

	m := devices[busType]
	if m == nil {
		panic(fmt.Errorf("device %q has invalid BusType %d", name, busType))
	}

	if _, exists := m[name]; exists {
		panic(fmt.Errorf("%s device with name %q already exists", busType.Label(), name))
	}

	m[name] = device
}

func lookupDevice(bus bus.BusType, name string) Device {
	mutex.Lock()
	defer mutex.Unlock()

	m := devices[bus]
	if m == nil {
		return nil
	}

	return m[strings.ToLower(name)]
}
