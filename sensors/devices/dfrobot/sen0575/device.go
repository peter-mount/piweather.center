package rainfall

import (
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/sensors"
)

func init() {
	sensors.RegisterDevice(&device{})
}

type device struct {
}

func (d *device) Info() sensors.DeviceInfo {
	return sensors.DeviceInfo{
		ID:           "Sen0575",
		Manufacturer: "DFRobot",
		Model:        "SEN0575",
		Description:  "Rain Fall Detector",
		BusType:      sensors.BusI2C,
	}
}

func (d *device) NewTask() task.Task {
	i := &Sen0575{}
	return i.task
}
