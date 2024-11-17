package device

import (
	"sort"
	"strings"
)

func ListDevices() []DeviceInfo {
	mutex.Lock()
	defer mutex.Unlock()

	var ret []DeviceInfo

	for _, bus := range devices {
		for _, device := range bus {
			info := device.Info()
			// IDs are case-insensitive internally so ensure this when exposing them
			info.ID = strings.ToLower(info.ID)
			ret = append(ret, info)
		}
	}

	// Sort by Manufacturer, Model then BusType
	sort.SliceStable(ret, func(i, j int) bool {
		c := strings.Compare(ret[i].Manufacturer, ret[j].Manufacturer)
		if c == 0 {
			c = strings.Compare(ret[i].Model, ret[j].Model)
		}
		if c == 0 && ret[i].BusType < ret[j].BusType {
			c = -1
		}
		return c < 0
	})

	return ret
}
