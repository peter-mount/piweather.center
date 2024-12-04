//go:build !(aix || plan9 || solaris || windows)

package device

import (
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"strings"
)

var (
	i2cDevices map[string]Device
)

func init() {
	i2cDevices = make(map[string]Device)
}

func registerI2CDevice(device Device) {
	name := strings.ToLower(device.Info().ID)
	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := i2cDevices[name]; exists {
		panic(fmt.Errorf("i2c %s device with name %q already exists", name))
	}
	i2cDevices[name] = device
}

func listI2CDevices() []DeviceInfo {
	mutex.Lock()
	defer mutex.Unlock()
	var r []DeviceInfo
	for _, i2cDevice := range i2cDevices {
		r = append(r, i2cDevice.Info())
	}
	return r
}

// I2CDevice represents a Device that operates over an I2C bus
type I2CDevice interface {
	Device
	// NewInstance returns a new Instance bound to a specific I2C bus and address on that bus
	NewInstance(int, uint8) Instance
}

// LookupI2CDevice returns the named I2CDevice. This will fail with an error if no device has been registered
func LookupI2CDevice(name string) (I2CDevice, error) {
	n := strings.ToLower(name)
	mutex.Lock()
	defer mutex.Unlock()
	dev, exists := i2cDevices[n]
	if !exists {
		return nil, deviceNotFound
	}
	// If this fails then RegisterDevice failed when checking the interface
	return dev.(I2CDevice), nil
}

// BasicI2CDevice contains the common details required by any I2CDevice.
type BasicI2CDevice struct {
	device I2CDevice
	bus    int
	addr   uint8
}

// NewBasicI2CDevice returns a populated BasicI2CDevice instance which should be used by an Instance returned by I2CDevice.NewInstance
func NewBasicI2CDevice(device I2CDevice, bus int, addr uint8) BasicI2CDevice {
	return BasicI2CDevice{device: device, bus: bus, addr: addr}
}

// Init the device. The default implementation does nothing.
// Override this if you need to initialise the device first.
func (b BasicI2CDevice) Init() error {
	return nil
}

// NewReading creates a new Reading ready for use in taking measurements
func (b BasicI2CDevice) NewReading() *reading.Reading {
	return newReading(b.device)
}

func (b BasicI2CDevice) ReadSensor() (*reading.Reading, error) {
	return nil, fmt.Errorf("device %q does not implement ReadSensor()", b.device.Info().ID)
}

func (b BasicI2CDevice) RunDevice(_ publisher.Publisher) error {
	return fmt.Errorf("device %q must implement RunDevice due to limitations in go", b.device.Info().ID)
}

// UseDevice will execute the Task against this device.
func (b BasicI2CDevice) UseDevice(task i2c.Task) error {
	return i2c.UseI2CConcurrent(b.bus, b.addr, task)
}
