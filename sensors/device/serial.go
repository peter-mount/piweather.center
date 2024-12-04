//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package device

import (
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"go.bug.st/serial"
	"strings"
	"time"
)

var (
	serialDevices map[string]Device
)

func init() {
	serialDevices = make(map[string]Device)
}

func registerSerialDevice(device Device) {
	name := strings.ToLower(device.Info().ID)
	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := serialDevices[name]; exists {
		panic(fmt.Errorf("serial %s device with name %q already exists", name))
	}
	serialDevices[name] = device
}

func listSerialDevices() []DeviceInfo {
	mutex.Lock()
	defer mutex.Unlock()
	var r []DeviceInfo
	for _, device := range serialDevices {
		r = append(r, device.Info())
	}
	return r
}

type SerialDevice interface {
	Device
	NewInstance(portName string, mode *serial.Mode) Instance
}

func LookupSerialDevice(name string) (SerialDevice, error) {
	n := strings.ToLower(name)
	mutex.Lock()
	defer mutex.Unlock()
	dev, exists := serialDevices[n]
	if !exists {
		return nil, deviceNotFound
	}
	// If this fails then RegisterDevice failed when checking the interface
	return dev.(SerialDevice), nil
}

type BasicSerialDevice struct {
	device   SerialDevice
	portName string
	mode     *serial.Mode
	port     serial.Port
}

func NewBasicSerialDevice(device SerialDevice, portName string, mode *serial.Mode) BasicSerialDevice {
	return BasicSerialDevice{
		device:   device,
		portName: portName,
		mode:     mode,
	}
}

func (b *BasicSerialDevice) Open() error {
	if b.port == nil {
		p, err := serial.Open(b.portName, b.mode)
		if err != nil {
			return err
		}
		b.port = p
	}
	return nil
}

func (b *BasicSerialDevice) Close() error {
	if b.port != nil {
		defer func() {
			b.port = nil
		}()
		return b.port.Close()
	}
	return nil
}

func (b *BasicSerialDevice) Run(f func() error) error {
	if err := b.Open(); err != nil {
		return err
	}
	defer b.Close()
	return f()
}

func (b *BasicSerialDevice) Init() error {
	return nil
}

// NewReading creates a new Reading ready for use in taking measurements
func (b *BasicSerialDevice) NewReading() *reading.Reading {
	return newReading(b.device)
}

func (b *BasicSerialDevice) ReadSensor() (*reading.Reading, error) {
	return nil, fmt.Errorf("device %q does not implement ReadSensor()", b.device.Info().ID)
}

func (b *BasicSerialDevice) RunDevice(_ publisher.Publisher) error {
	return fmt.Errorf("device %q does not implement RunDevice()", b.device.Info().ID)
}

func (b *BasicSerialDevice) Read(d []byte) (int, error) {
	return b.port.Read(d)
}

func (b *BasicSerialDevice) Write(d []byte) (int, error) {
	return b.port.Write(d)
}

func (b *BasicSerialDevice) Drain() error {
	return b.port.Drain()
}

func (b *BasicSerialDevice) ResetInputBuffer() error {
	return b.port.ResetInputBuffer()
}

func (b *BasicSerialDevice) ResetOutputBuffer() error {
	return b.port.ResetOutputBuffer()
}

// SetReadTimeout sets the timeout for the Read operation or use serial.NoTimeout
// to disable read timeout.
func (b *BasicSerialDevice) SetReadTimeout(t time.Duration) error {
	return b.port.SetReadTimeout(t)
}
