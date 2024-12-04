package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type SensorList struct {
	Pos     lexer.Position
	Sensors []*Sensor `parser:"@@*"`
}

func (c *visitor[T]) SensorList(d *SensorList) error {
	var err error
	if d != nil {
		if c.sensorList != nil {
			err = c.sensorList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Sensors {
				err = c.Sensor(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initSensorList(v Visitor[*initState], d *SensorList) error {
	v.Get().sensors = make(map[string]*Sensor)
	return nil
}

func (b *builder[T]) SensorList(f func(Visitor[T], *SensorList) error) Builder[T] {
	b.sensorList = f
	return b
}
