//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package device

import (
	"github.com/peter-mount/piweather.center/sensors/bus"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"go.bug.st/serial"
	"time"
)

type SerialDevice interface {
	Device
	NewInstance(portName string, mode *serial.Mode) Instance
}

func LookupSerialDevice(name string) (SerialDevice, error) {
	dev := lookupDevice(bus.BusSerial, name)
	if dev == nil {
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
	return nil, deviceNotImplemented
}

func (b *BasicSerialDevice) RunDevice(_ publisher.Publisher) error {
	return deviceNotImplemented
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
