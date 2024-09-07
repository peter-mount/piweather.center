package sensors

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"strings"
	"sync"
)

var (
	mutex   sync.Mutex
	devices map[string]Device
)

func init() {
	devices = make(map[string]Device)
}

type Device interface {
	Info() DeviceInfo
	NewTask() task.Task
}

type DeviceInfo struct {
	ID           string
	Manufacturer string
	Model        string
	Description  string
	BusType      BusType
}

// RegisterDevice registers a new Device handler.
// This will panic if the ID has already been registered.
// IDs are case-insensitive.
func RegisterDevice(device Device) {
	mutex.Lock()
	defer mutex.Unlock()

	name := strings.ToLower(device.Info().ID)

	if name == "" {
		panic(errors.New("device name cannot be empty"))
	}

	if _, exists := devices[name]; exists {
		panic(fmt.Errorf("device with name %s already exists", name))
	}

	devices[name] = device
}

func LookupDevice(name string) Device {
	mutex.Lock()
	defer mutex.Unlock()
	return devices[strings.ToLower(name)]
}
