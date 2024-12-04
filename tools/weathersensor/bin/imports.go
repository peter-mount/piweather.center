//go:build !(aix || plan9 || solaris)

package main

// Import sensors only on platforms that support at least one of serial or i2c
import (
	_ "github.com/peter-mount/piweather.center/sensors/devices"
)
