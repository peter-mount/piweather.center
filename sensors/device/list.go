package device

import (
	"sort"
	"strings"
)

func ListDevices() []DeviceInfo {
	var devices []DeviceInfo
	devices = append(devices, listI2CDevices()...)
	devices = append(devices, listSerialDevices()...)

	// IDs are case-insensitive internally so ensure this when exposing them
	for _, info := range devices {
		info.ID = strings.ToLower(info.ID)
	}

	// Sort by Manufacturer, Model then BusType
	sort.SliceStable(devices, func(i, j int) bool {
		c := strings.Compare(devices[i].Manufacturer, devices[j].Manufacturer)
		if c == 0 {
			c = strings.Compare(devices[i].Model, devices[j].Model)
		}
		if c == 0 && devices[i].BusType < devices[j].BusType {
			c = -1
		}
		return c < 0
	})

	return devices
}
