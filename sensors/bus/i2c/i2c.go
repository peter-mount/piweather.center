package i2c

import (
	"fmt"
	"os"
)

// I2C provides an interface to an I2C device providing raw or SMBus protocols
type I2C interface {
	// Read a []byte from the I2C device
	Read([]byte) (int, error)

	// Write a []byte to the I2C device
	Write([]byte) (int, error)

	// WriteByte writes a single byte to the I2C device
	WriteByte(byte) error

	// ReadRegister reads an SMBus register from the device
	ReadRegister(reg uint8, buf []byte) error

	// ReadRegisterUint8 reads an SMBus register from the device as a Uint8
	ReadRegisterUint8(register uint8) (uint8, error)

	// ReadRegisterUint16 reads an SMBus register from the device as a Uint16
	ReadRegisterUint16(register uint8) (uint16, error)

	// ReadRegisterUint32 reads an SMBus register from the device as a Uint32
	ReadRegisterUint32(register uint8) (uint32, error)

	// ReadRegisterUint64 reads an SMBus register from the device as a Uint64
	ReadRegisterUint64(register uint8) (uint64, error)

	// WriteRegister writes to a SMBus register on the device
	WriteRegister(reg uint8, buf []byte) error

	// WriteRegisterUint8 writes an uint8 to a SMBus register on the device
	WriteRegisterUint8(register, value uint8) error

	// WriteRegisterUint16 writes an uint16 to a SMBus register on the device
	WriteRegisterUint16(register uint8, value uint16) error

	// WriteRegisterUint32 writes an uint32 to a SMBus register on the device
	WriteRegisterUint32(register uint8, value uint32) error

	// WriteRegisterUint64 writes an uint64 to a SMBus register on the device
	WriteRegisterUint64(register uint8, value uint64) error
}

type i2cDevice struct {
	f *os.File
}

// Task handles accessing the specific device
type Task func(I2C) error

// Then returns a new Task which will invoke both tasks in sequence
func (a Task) Then(b Task) Task {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(i2c I2C) error {
		err := a(i2c)
		if err == nil {
			err = b(i2c)
		}
		return err
	}
}

// Of returns a Task formed of the provided tasks.
// Tasks are invoked in the sequence provided.
// If no Tasks are provided then nil is returned, which is the nop-Task.
func Of(task Task, tasks ...Task) Task {
	for _, t := range tasks {
		task = task.Then(t)
	}
	return task
}

func (a Task) execute(bus int, addr uint8) error {
	if a == nil {
		return nil
	}

	f, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%d", bus), os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := ioctlAddr(f.Fd(), addr); err != nil {
		return err
	}

	return a(&i2cDevice{f: f})
}

// UseI2C will call a Task against a specific I2C device.
//
// If task is nil this does nothing.
//
// Note: This function provides direct access to the I2C bus with no locking,
// so multiple goroutines can access the device at the same time, which might cause
// concurrency issues with the device.
//
// You are recommended to use UseI2CConcurrent instead as that ensures only one Task can access
// a specific device at any one time within this process.
//
// Accessing different devices on the same bus in different goroutines is supported
// as the i2c kernel module ensures that IO on each bus is thread safe.
func UseI2C(bus int, addr uint8, task Task) error {
	if task == nil {
		return nil
	}
	return task.execute(bus, addr)
}
