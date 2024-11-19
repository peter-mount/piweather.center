package sensors

import (
	"github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type Builder[T any] interface {
	Sensors(func(sensors.SensorVisitor[T], *sensors.Sensors) error) Builder[T]
	Sensor(func(sensors.SensorVisitor[T], *sensors.Sensor) error) Builder[T]
	I2C(func(sensors.SensorVisitor[T], *sensors.I2C) error) Builder[T]
	Serial(func(sensors.SensorVisitor[T], *sensors.Serial) error) Builder[T]
	Poll(func(sensors.SensorVisitor[T], *time.CronTab) error) Builder[T]
	Publisher(func(sensors.SensorVisitor[T], *sensors.Publisher) error) Builder[T]
	Build() sensors.SensorVisitor[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

type builder[T any] struct {
	common[T]
}

func (b *builder[T]) Build() sensors.SensorVisitor[T] {
	v := &visitor[T]{}
	v.common = b.common
	return v
}

func (b *builder[T]) Sensors(f func(sensors.SensorVisitor[T], *sensors.Sensors) error) Builder[T] {
	b.sensors = f
	return b
}
func (b *builder[T]) Sensor(f func(sensors.SensorVisitor[T], *sensors.Sensor) error) Builder[T] {
	b.sensor = f
	return b
}
func (b *builder[T]) I2C(f func(sensors.SensorVisitor[T], *sensors.I2C) error) Builder[T] {
	b.i2c = f
	return b
}

func (b *builder[T]) Serial(f func(sensors.SensorVisitor[T], *sensors.Serial) error) Builder[T] {
	b.serial = f
	return b
}

func (b *builder[T]) Poll(f func(sensors.SensorVisitor[T], *time.CronTab) error) Builder[T] {
	b.poll = f
	return b
}

func (b *builder[T]) Publisher(f func(sensors.SensorVisitor[T], *sensors.Publisher) error) Builder[T] {
	b.publisher = f
	return b
}
