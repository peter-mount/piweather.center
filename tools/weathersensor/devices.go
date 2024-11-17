package weathersensor

import (
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/util/table"
	"sort"
)

func (s *Service) listDevices() error {
	t := table.New("ID", "Description", "Manufacturer", "Model", "Bus", "Mode")

	devs := device.ListDevices()

	sort.SliceStable(devs, func(i, j int) bool {
		return devs[i].ID < devs[j].ID
	})

	for _, dev := range devs {
		t.NewRow().
			Add(dev.ID).
			Add(dev.Description).
			Add(dev.Manufacturer).
			Add(dev.Model).
			Add(dev.BusType.Label()).
			Add(dev.PollMode.Label())
	}

	fmt.Println(t.String())

	return nil
}
