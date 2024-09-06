package i2c

import (
	"time"
	"unsafe"
)

type i2cCmd struct {
	rw  uint8
	cmd uint8
	len uint32
	ptr unsafe.Pointer
}

func (c *i2cDevice) smbusCommand(rw, reg uint8, len uint32, ptr unsafe.Pointer) error {
	cmd := i2cCmd{rw: rw, cmd: reg, len: len, ptr: ptr}
	cmdPtr := unsafe.Pointer(&cmd)
	err := ioctl(c.f.Fd(), i2cSMBus, uintptr(cmdPtr))

	// FIXME Investigate this hack as, without a 10ms delay after the command has been invoked, corruption or errors occur. With the delay it's stable.
	if err == nil {
		time.Sleep(10 * time.Millisecond)
	}

	return err
}
