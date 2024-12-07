package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type StationEntry struct {
	Pos         lexer.Position
	Calculation *Calculation `parser:"( @@"`
	Sensor      *Sensor      `parser:"| @@"`
	Dashboard   *Dashboard   `parser:"| @@ )"`
}

func (c *visitor[T]) StationEntry(d *StationEntry) error {
	var err error
	if d != nil {
		if c.stationEntry != nil {
			err = c.stationEntry(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			switch {
			case d.Calculation != nil:
				err = c.Calculation(d.Calculation)

			case d.Dashboard != nil:
				err = c.Dashboard(d.Dashboard)

			case d.Sensor != nil:
				err = c.Sensor(d.Sensor)
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) StationEntry(f func(Visitor[T], *StationEntry) error) Builder[T] {
	b.stationEntry = f
	return b
}
