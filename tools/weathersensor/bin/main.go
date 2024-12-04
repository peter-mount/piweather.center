package main

import (
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/sensors/devices"
	"github.com/peter-mount/piweather.center/tools/weathersensor"
	"log"
)

// main entry point for the weathersensor binary.
//
// By default, this includes all devices on all supported buses for a platform.
//
// To customise what devices are included, e.g. you only want a subset and want a smaller binary:
//
//  1. Copy this file into an empty directory
//
//  2. Remove the following import, which includes everything:
//     _ "github.com/peter-mount/piweather.center/sensors/devices"
//
//  3. Add imports for the devices you require.
//     e.g. for just the GMC320 geiger-counter, add the following import:
//     _ "github.com/peter-mount/piweather.center/sensors/devices/gq/gmc320"
//
//  4. Compile the new file - this will then generate your custom binary.
//
// Note: If you want to add a driver that's not in the main source tree, follow the same instructions but instead
//       of removing the import on step 2, just add the appropriate import for your driver in step 3.

func main() {
	err := kernel.Launch(
		// Must be first to enable us to capture this first before anything else does
		&weathersensor.ListDevices{},
		&weathersensor.Service{},
		// If developing or soak testing device drivers, uncomment the following line which will then report
		// some useful metrics when the binary exits, e.g. memory use or runtime.
		//
		// This can be useful if the device fails intermittently.
		//
		// e.g. I had an intermittent issue with I2C where it crashed at random times and this told me how long it ran
		// before the failure. In that case it was a cabling fault causing the i2c bus to crash.
		//
		//&kernel.MemUsage{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
