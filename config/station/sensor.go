package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type Sensor struct {
	Pos       lexer.Position
	Target    *Metric      `parser:"'sensor' '(' @@"`
	Http      *Http        `parser:"( @@"`
	I2C       *I2C         `parser:"| @@"`
	Serial    *Serial      `parser:"| @@ )"`
	Poll      time.CronTab `parser:"('poll' '(' @@ ')')?"`
	Publisher []*Publisher `parser:"'publish' '(' @@+ ')' ')'"`
}

func (c *visitor[T]) Sensor(d *Sensor) error {
	var err error
	if d != nil {
		if c.sensor != nil {
			err = c.sensor(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Metric(d.Target)
		}

		if err == nil {
			err = c.sensorCommon(d)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

// Shared by visitor.Sensor() and initSensor()
// Used specifically to ensure we visit everything EXCEPT the target metric during init
func (c *visitor[T]) sensorCommon(d *Sensor) error {
	var err error
	switch {
	case d.Http != nil:
		err = c.Http(d.Http)
	case d.I2C != nil:
		err = c.I2C(d.I2C)
	case d.Serial != nil:
		err = c.Serial(d.Serial)
	}

	if err == nil {
		err = c.CronTab(d.Poll)
	}

	if err == nil {
		for _, e := range d.Publisher {
			err = c.Publisher(e)
			if err != nil {
				break
			}
		}
	}

	return err
}

func initSensor(v Visitor[*initState], d *Sensor) error {
	s := v.Get()

	// Should never occur
	if d.Target == nil {
		return participle.Errorf(d.Pos, "target is required")
	}

	if d.Target.Unit != nil {
		return participle.Errorf(d.Target.Unit.Pos, "unit is invalid as a target for sensors")
	}

	// Check Target is unique within the station
	if e, exists := s.sensors[d.Target.Name]; exists {
		return participle.Errorf(d.Pos, "sensor %q already defined at %s", d.Target.Name, e.String())
	}
	s.sensors[d.Target.Name] = d.Pos

	s.sensorPrefix = d.Target.Name + "."

	return (v.(*visitor[*initState])).sensorCommon(d)
}

func (b *builder[T]) Sensor(f func(Visitor[T], *Sensor) error) Builder[T] {
	b.sensor = f
	return b
}
