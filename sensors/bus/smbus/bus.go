package smbus

import "sync"

type i2cBus struct {
	bus     int
	devices map[uint8]*smBus
	mutex   sync.Mutex
}

func (b *i2cBus) handle(device *smBus, task Task) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
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
}

type Task func(SMBus) error
