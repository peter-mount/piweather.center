package i2c

import (
	"encoding/binary"
	"io"
	"time"
	"unsafe"
)

// Write sends buf to the remote i2c device.
// The interpretation of the message is implementation dependant.
func (c *i2cDevice) Write(buf []byte) (int, error) {
	return c.f.Write(buf)
}

// WriteByte sends a single byte to the remote i2c device.
// The interpretation of the message is implementation dependant.
func (c *i2cDevice) WriteByte(b byte) error {
	var buf [1]byte
	buf[0] = b
	n, err := c.f.Write(buf[:])
	if err == nil && n == 0 {
		err = io.ErrShortWrite
	}
	return err
}

// WriteRegister writes the buf byte slice to a designated register.
func (c *i2cDevice) WriteRegister(reg uint8, buf []byte) error {
	if len(buf) > int(i2cSMBusBlockMax) {
		return errSmbusBlockSize
	}

	data := make([]byte, 1+len(buf), i2cSMBusBlockMax+2)
	data[0] = byte(len(buf))
	copy(data[1:], buf)

	err := c.smbusCommand(i2cSMBusWrite, reg, i2cSMBusI2CBlockData, unsafe.Pointer(&data[0]))
	if err == nil {
		time.Sleep(50 * time.Millisecond)
	}
	return err
}

func (c *i2cDevice) WriteRegisterUint8(register, value uint8) error {
	return c.WriteRegister(register, []byte{value})
}

func (c *i2cDevice) WriteRegisterUint16(register uint8, value uint16) error {
	buf := binary.LittleEndian.AppendUint16(nil, value)
	return c.WriteRegister(register, buf[:])
}

func (c *i2cDevice) WriteRegisterUint32(register uint8, value uint32) error {
	buf := binary.LittleEndian.AppendUint32(nil, value)
	return c.WriteRegister(register, buf[:])
}

func (c *i2cDevice) WriteRegisterUint64(register uint8, value uint64) error {
	buf := binary.LittleEndian.AppendUint64(nil, value)
	return c.WriteRegister(register, buf[:])
}
