//go:build aix || plan9 || solaris || windows

package device

func registerI2CDevice(device Device) {
	// does nothing as those devices are not supported and in that instance they wouldn't be built anyhow.
	// This just allows RegisterDevice to work
}

func listI2CDevices() []DeviceInfo {
	return nil
}
