//go:build !(aix || plan9 || solaris || windows)

package devices

/*
	import each I2C device we will support

	I2C is not supported on aix, plan9, solaris or windows
	due to syscall.SYS_IOCTL being undefined
*/
import (
	_ "github.com/peter-mount/piweather.center/sensors/devices/asair/aht10"
	_ "github.com/peter-mount/piweather.center/sensors/devices/dfrobot/sen0575"
	_ "github.com/peter-mount/piweather.center/sensors/devices/rohm/bh1750"
)
