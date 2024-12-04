package weathersensor

import (
	"fmt"
	"github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/util/table"
	"os"
	"sort"
)

type ListDevices struct {
	ListDevices *bool `kernel:"flag,list-devices,List Devices"`
}

func (s *ListDevices) PostInit() error {
	// PostInit if we want to list devices do so now and exit so we don't
	// start the rest of the system, especially connecting to rabbitmq
	if *s.ListDevices {
		s.listDevices()
		os.Exit(0)
	}
	return nil
}

func (s *ListDevices) listDevices() {
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
}
