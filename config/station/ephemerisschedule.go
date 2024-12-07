package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
)

// EphemerisSchedule defines when the listed targets are calculated
type EphemerisSchedule struct {
	Pos     lexer.Position
	Every   *time.CronTab      `parser:"'every' @@"` // When to calculate the following entries
	Targets []*EphemerisTarget `parser:"@@+"`        // Targets to calculate
}

func (c *visitor[T]) EphemerisSchedule(d *EphemerisSchedule) error {
	var err error
	if d != nil {
		if c.ephemerisSchedule != nil {
			err = c.ephemerisSchedule(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.CronTab(d.Every)
		}

		if err == nil {
			for _, e := range d.Targets {
				err = c.EphemerisTarget(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) EphemerisSchedule(f func(Visitor[T], *EphemerisSchedule) error) Builder[T] {
	b.ephemerisSchedule = f
	return b
}
