package smbus

import "github.com/go-daq/smbus"

func (d *smBus) WriteRegister(register uint8, buf []byte) error {
	c, err := smbus.Open(d.bus.bus, d.i2cAddr)
	if err != nil {
		return err
	}
	defer c.Close()

	return c.WriteBlockData(d.i2cAddr, register, buf)
}

func (d *smBus) WriteRegisterUint8(register, value uint8) error {
	return d.WriteRegister(register, []byte{value})
}
