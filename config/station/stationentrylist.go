package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type StationEntryList struct {
	Pos     lexer.Position
	Entries []*StationEntry `parser:"@@*"`
}

func (c *visitor[T]) StationEntryList(d *StationEntryList) error {
	var err error
	if d != nil {
		if c.stationEntry != nil {
			err = c.stationEntryList(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Entries {
				err = c.StationEntry(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initStationEntryList(v Visitor[*initState], d *StationEntryList) error {
	s := v.Get()

	if s.stationPrefix == "" {
		// should never occur
		return errors.Errorf(d.Pos, "stationPrefix not defined")
	}

	// Clear sensorPrefix. If the children need it they can set it, e.g. Sensor does
	s.sensorPrefix = ""

	return nil
}

func (b *builder[T]) StationEntryList(f func(Visitor[T], *StationEntryList) error) Builder[T] {
	b.stationEntryList = f
	return b
}

// HomeDashboard returns the home dashboard which is by definition the first dashboard in the definition
func (s *StationEntryList) HomeDashboard() *Dashboard {
	for _, e := range s.Entries {
		if e.Dashboard != nil {
			return e.Dashboard
		}
	}
	return nil
}
