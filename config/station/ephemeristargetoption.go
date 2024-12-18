package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/astro/api"
)

// EphemerisTargetOption declares the values to create metrics
type EphemerisTargetOption struct {
	Pos             lexer.Position
	Target          string `parser:"@Ident"`
	ephemerisOption api.EphemerisOption
}

func (c *visitor[T]) EphemerisTargetOption(d *EphemerisTargetOption) error {
	var err error
	if d != nil {
		if c.ephemerisTargetOption != nil {
			err = c.ephemerisTargetOption(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initEphemerisTargetOption(_ Visitor[*initState], d *EphemerisTargetOption) error {
	d.ephemerisOption = api.ParseEphemerisOption(d.Target)
	if d.ephemerisOption == 0 {
		return errors.Errorf(d.Pos, "Invalid target %q", d.Target)
	}

	return nil
}

func (b *builder[T]) EphemerisTargetOption(f func(Visitor[T], *EphemerisTargetOption) error) Builder[T] {
	b.ephemerisTargetOption = f
	return b
}

func (d EphemerisTargetOption) EphemerisOption() api.EphemerisOption {
	return d.ephemerisOption
}
