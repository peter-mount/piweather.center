package weathersensor

import (
	"fmt"
	"github.com/peter-mount/piweather.center/sensors"
	"strings"
)

func (s *Service) listDevices() error {
	devs := sensors.ListDevices()

	format := "%-10.10s %-8.8s %-10.10s %-10.10s %s\n"

	str := fmt.Sprintf(format, "ID", "Bus", "Manufacturer", "Model", "Description")

	fmt.Println(str + strings.Repeat("-", len(str)))

	for _, dev := range devs {
		fmt.Printf(
			format,
			dev.ID,
			dev.BusType.Label(),
			dev.Manufacturer,
			dev.Model,
			dev.Description)
	}

	return nil
}
