package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"sort"
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

func (s *StationEntryList) AddStationEntry(e *StationEntry) {
	s.Entries = append(s.Entries, e)
}

func (s *StationEntryList) Merge(b *StationEntryList) {
	if b != nil {
		for _, e := range b.Entries {
			s.AddStationEntry(e)
		}
	}
}

func (s *StationEntryList) Sort() {
	e := s.Entries
	sort.SliceStable(e, func(i, j int) bool {
		a, b := e[i], e[j]
		at, bt := a.GetTarget(), b.GetTarget()

		switch {
		// Sort by target A < B
		case at != "" && bt != "":
			return at < bt
			// A before non target B
		case at != "":
			return false
			// B before non target A
		case bt != "":
			return true
		default:
			// Sort by file position
			return a.Pos.String() < b.Pos.String()
		}
	})
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
