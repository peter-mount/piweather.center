package smbus

import (
	"encoding/binary"
	"github.com/go-daq/smbus"
)

func (d *smBus) ReadRegister(register uint8, length int) ([]byte, error) {
	c, err := smbus.Open(d.bus.bus, d.i2cAddr)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	buf := make([]byte, length)

	err = c.ReadBlockData(d.i2cAddr, register, buf)

	return buf, err
}

func (d *smBus) ReadRegisterUint16(register uint8) (uint16, error) {
	buf, err := d.ReadRegister(register, 2)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(buf), nil
}

func (d *smBus) ReadRegisterUint32(register uint8) (uint32, error) {
	buf, err := d.ReadRegister(register, 4)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buf), nil
}
