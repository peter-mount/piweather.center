package sensors

import (
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type visitor[T any] struct {
	common[T]
}

type common[T any] struct {
	sensors   func(sensors.SensorVisitor[T], *sensors.Sensors) error
	sensor    func(sensors.SensorVisitor[T], *sensors.Sensor) error
	i2c       func(sensors.SensorVisitor[T], *sensors.I2C) error
	serial    func(sensors.SensorVisitor[T], *sensors.Serial) error
	poll      func(sensors.SensorVisitor[T], *time.CronTab) error
	publisher func(sensors.SensorVisitor[T], *sensors.Publisher) error
}

func (v *visitor[T]) Sensors(b *sensors.Sensors) error {
	var err error
	if b != nil {
		if v.sensors != nil {
			err = v.sensors(v, b)
		}
		if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
			return nil
		}

		for _, s := range b.Sensors {
			err = v.Sensor(s)
		}
	}
	return err
}

func (v *visitor[T]) Sensor(b *sensors.Sensor) error {
	var err error
	if b != nil {
		if v.sensor != nil {
			err = v.sensor(v, b)
		}
		if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
			return nil
		}

		if err == nil {
			err = v.I2C(b.I2C)
		}

		if err == nil {
			err = v.Serial(b.Serial)
		}

		if err == nil {
			err = v.Poll(b.Poll)
		}

		for _, p := range b.Publisher {
			if err == nil {
				err = v.Publisher(p)
			}
		}
	}
	return err
}

func (v *visitor[T]) I2C(b *sensors.I2C) error {
	var err error
	if b != nil && v.i2c != nil {
		err = v.i2c(v, b)
	}
	return err
}

func (v *visitor[T]) Serial(b *sensors.Serial) error {
	var err error
	if b != nil && v.serial != nil {
		err = v.serial(v, b)
		if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Poll(b *time.CronTab) error {
	var err error
	if b != nil && v.poll != nil {
		err = v.poll(v, b)
	}
	return err
}

func (v *visitor[T]) Publisher(b *sensors.Publisher) error {
	var err error
	if b != nil {
		if v.publisher != nil {
			err = v.publisher(v, b)
		}
		if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
			return nil
		}
	}
	return err
}
