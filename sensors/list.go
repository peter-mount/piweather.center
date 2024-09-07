package sensors

import (
	"sort"
	"strings"
)

func ListDevices() []DeviceInfo {
	mutex.Lock()
	defer mutex.Unlock()

	var ret []DeviceInfo
	for _, device := range devices {
		info := device.Info()
		// IDs are case-insensitive internally so ensure this when exposing them
		info.ID = strings.ToLower(info.ID)
		ret = append(ret, info)
	}

	sort.SliceStable(ret, func(i, j int) bool {
		c := strings.Compare(ret[i].Manufacturer, ret[j].Manufacturer)
		if c == 0 {
			c = strings.Compare(ret[i].Model, ret[j].Model)
		}
		return c < 0
	})

	return ret
}
