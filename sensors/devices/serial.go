//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package devices

/*
	Import each Serial device we will support

	Serial is not supported on aix, dragonfly, illumos, netbsd, plan9 or solaris
	due to no serial support.

	Serial is not supported on linux on loong64, ppc64 or ppc64le
	due to fdset having NFDBITS undefined
*/
import (
	_ "github.com/peter-mount/piweather.center/sensors/devices/gq/gmc320"
)
