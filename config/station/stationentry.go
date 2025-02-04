package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type StationEntry struct {
	Pos           lexer.Position
	CalculateFrom *CalculateFrom `parser:"( 'calculate' ( 'from' @@"`
	Calculation   *Calculation   `parser:"  | @@ )"`
	Dashboard     *Dashboard     `parser:"| @@"`
	Ephemeris     *Ephemeris     `parser:"| @@"`
	Tasks         *Tasks         `parser:"| @@"`
	Sensor        *Sensor        `parser:"| @@ )"`
}

func (e *StationEntry) GetTarget() string {
	switch {
	case e.CalculateFrom != nil:
		return e.CalculateFrom.From
	case e.Calculation != nil:
		return e.Calculation.OriginalTarget
	case e.Dashboard != nil:
		return e.Dashboard.Name
	case e.Ephemeris != nil:
		return e.Ephemeris.Target
	case e.Sensor != nil:
		return e.Sensor.Target.OriginalName
	default:
		// No target
		return ""
	}
}

func (c *visitor[T]) StationEntry(d *StationEntry) error {
	var err error
	if d != nil {
		if c.stationEntry != nil {
			err = c.stationEntry(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			switch {
			case d.CalculateFrom != nil:
				err = c.CalculateFrom(d.CalculateFrom)

			case d.Calculation != nil:
				err = c.Calculation(d.Calculation)

			case d.Dashboard != nil:
				err = c.Dashboard(d.Dashboard)

			case d.Ephemeris != nil:
				err = c.Ephemeris(d.Ephemeris)

			case d.Tasks != nil:
				err = c.Tasks(d.Tasks)

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
