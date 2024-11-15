//go:build !(aix || plan9 || solaris || windows)

package i2c

import "syscall"

func ioctl(fd, cmd, arg uintptr) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func ioctlAddr(fd uintptr, addr uint8) error {
	return ioctl(fd, i2cSlave, uintptr(addr))
}
