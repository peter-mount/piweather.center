package i2c

import (
	"encoding/binary"
	"unsafe"
)

// Read reads data from the remote i2c device into p.
func (c *i2cDevice) Read(p []byte) (int, error) {
	return c.f.Read(p)
}

// ReadRegister reads len(buf) data into the byte slice, from the designated register.
func (c *i2cDevice) ReadRegister(reg uint8, buf []byte) error {
	if len(buf) > int(i2cSMBusBlockMax) {
		return errSmbusBlockSize
	}

	data := make([]byte, len(buf)+1, i2cSMBusBlockMax+2)
	data[0] = byte(len(buf))
	err := c.smbusCommand(i2cSMBusRead, reg, i2cSMBusI2CBlockData, unsafe.Pointer(&data[0]))
	if err != nil {
		return err
	}

	copy(buf[:], data[1:len(buf)+1])

	return nil
}

func (c *i2cDevice) ReadRegisterUint8(register uint8) (uint8, error) {
	buf := make([]byte, 1)
	err := c.ReadRegister(register, buf)
	if err != nil {
		return 0, err
	}

	return buf[0], nil
}

func (c *i2cDevice) ReadRegisterUint16(register uint8) (uint16, error) {
	buf := make([]byte, 2)
	err := c.ReadRegister(register, buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(buf), nil
}

func (c *i2cDevice) ReadRegisterUint32(register uint8) (uint32, error) {
	buf := make([]byte, 4)
	err := c.ReadRegister(register, buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buf), nil
}

func (c *i2cDevice) ReadRegisterUint64(register uint8) (uint64, error) {
	buf := make([]byte, 8)
	err := c.ReadRegister(register, buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(buf), nil
}
