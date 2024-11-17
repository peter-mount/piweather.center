//go:build !(aix || plan9 || solaris || windows)

package device

import (
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
)

// I2CDevice represents a Device that operates over an I2C bus
type I2CDevice interface {
	Device
	// NewInstance returns a new Instance bound to a specific I2C bus and address on that bus
	NewInstance(int, uint8) Instance
}

// LookupI2CDevice returns the named I2CDevice. This will fail with an error if no device has been registered
func LookupI2CDevice(name string) (I2CDevice, error) {
	dev := lookupDevice(bus.BusI2C, name)
	if dev == nil {
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
	return nil, deviceNotImplemented
}

func (b BasicI2CDevice) RunDevice(_ publisher.Publisher) error {
	return deviceNotImplemented
}

// UseDevice will execute the Task against this device.
func (b BasicI2CDevice) UseDevice(task i2c.Task) error {
	return i2c.UseI2CConcurrent(b.bus, b.addr, task)
}
