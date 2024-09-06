package smbus

import (
	"github.com/go-daq/smbus"
	"sync"
)

type i2cBus struct {
	mutex   sync.Mutex
	bus     int              // the i2c bus in use
	devices map[uint8]*smBus // registered i2c addresses on this bus
}

func (b *i2cBus) handle(device *smBus, task Task) (err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	device.conn, err = smbus.Open(b.bus, device.i2cAddr)
	if err != nil {
		return
	}
	defer func() {
		_ = device.conn.Close()
		device.conn = nil
	}()

	return task(device)
}

type SMBus interface {
	ReadRegister(register uint8, length int) ([]byte, error)
	ReadRegisterUint16(register uint8) (uint16, error)
	ReadRegisterUint32(register uint8) (uint32, error)
	WriteRegister(register uint8, buf []byte) error
	WriteRegisterUint8(register, value uint8) error
}

type smBus struct {
	bus     *i2cBus
	i2cAddr uint8
	conn    *smbus.Conn
}

type Task func(SMBus) error
