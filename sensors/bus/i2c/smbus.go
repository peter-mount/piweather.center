package i2c

import "unsafe"

type i2cCmd struct {
	rw  uint8
	cmd uint8
	len uint32
	ptr unsafe.Pointer
}

func (c *i2cDevice) smbusCommand(rw, reg uint8, len uint32, ptr unsafe.Pointer) error {
	cmd := i2cCmd{rw: rw, cmd: reg, len: len, ptr: ptr}
	cmdPtr := unsafe.Pointer(&cmd)
	return ioctl(c.f.Fd(), i2cSMBus, uintptr(cmdPtr))
}
