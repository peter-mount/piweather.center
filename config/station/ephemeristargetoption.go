package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/config/util"
	"strings"
)

// EphemerisTargetOption declares the values to create metrics
type EphemerisTargetOption struct {
	Pos        lexer.Position
	Target     string `parser:"@('altitude' | 'azimuth' | 'ra' | 'dec' | 'distance')"`
	As         string `parser:"( 'as' @String )?"`
	targetType api.EphemerisOption
}

func (c *visitor[T]) EphemerisTargetOption(d *EphemerisTargetOption) error {
	var err error
	if d != nil {
		if c.ephemerisTargetOption != nil {
			err = c.ephemerisTargetOption(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initEphemerisTargetOption(v Visitor[*initState], d *EphemerisTargetOption) error {
	d.targetType = api.ParseEphemerisOption(d.Target)

	d.As = strings.ToLower(strings.TrimSpace(d.As))
	if d.As == "" {
		d.As = d.targetType.String()
	}

	d.As = v.Get().ephemerisTarget.Target + "." + d.As

	return nil
}

func (b *builder[T]) EphemerisTargetOption(f func(Visitor[T], *EphemerisTargetOption) error) Builder[T] {
	b.ephemerisTargetOption = f
	return b
}

func (d EphemerisTargetOption) TargetType() api.EphemerisOption {
	return d.targetType
}
