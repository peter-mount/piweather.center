package sensors

import "github.com/peter-mount/piweather.center/config/util/time"

type SensorVisitor[T any] interface {
	Sensors(*Sensors) error
	Sensor(*Sensor) error
	I2C(*I2C) error
	Serial(*Serial) error
	Poll(tab *time.CronTab) error
	Publisher(*Publisher) error
}
