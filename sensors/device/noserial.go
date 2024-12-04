//go:build aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le))

package device

func registerSerialDevice(device Device) {
	// does nothing as those devices are not supported and in that instance they wouldn't be built anyhow.
	// This just allows RegisterDevice to work
}

func listSerialDevices() []DeviceInfo {
	return nil
}
