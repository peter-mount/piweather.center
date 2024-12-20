package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/astro/api"
	"strings"
)

// EphemerisTarget defines the target to calculate
type EphemerisTarget struct {
	Pos             lexer.Position
	Target          string                   `parser:"@Ident"`                  // Object to calculate
	Options         []*EphemerisTargetOption `parser:"'(' (@@ (',' @@)*)? ')'"` // Parameters to include
	As              string                   `parser:"( 'as' @String )?"`       // Override the Target for the metric name
	targetType      EphemerisTargetType      // Computed target type
	ephemerisOption api.EphemerisOption      // Computed options
}

func (c *visitor[T]) EphemerisTarget(d *EphemerisTarget) error {
	var err error
	if d != nil {
		if c.ephemerisTarget != nil {
			err = c.ephemerisTarget(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Options {
				err = c.EphemerisTargetOption(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initEphemerisTarget(v Visitor[*initState], d *EphemerisTarget) error {
	st := v.Get()

	d.targetType = ParseEphemerisTargetType(d.Target)

	d.As = strings.ToLower(strings.TrimSpace(d.As))
	if d.As == "" {
		d.As = d.Target
	}
	st.ephemerisTarget = d

	// Work out the required targets
	var ephemerisOption api.EphemerisOption
	for _, opt := range d.Options {
		_ = initEphemerisTargetOption(v, opt)
		ephemerisOption = ephemerisOption | opt.EphemerisOption()
	}
	if ephemerisOption == 0 {
		ephemerisOption = api.AllOptions
	}
	d.ephemerisOption = ephemerisOption

	return errors.VisitorStop
}

func (b *builder[T]) EphemerisTarget(f func(Visitor[T], *EphemerisTarget) error) Builder[T] {
	b.ephemerisTarget = f
	return b
}

func (d *EphemerisTarget) GetTarget() EphemerisTargetType {
	return d.targetType
}

func (d *EphemerisTarget) GetEphemerisOption() api.EphemerisOption {
	return d.ephemerisOption
}

type EphemerisTargetType uint8

const (
	EphemerisTargetMercury EphemerisTargetType = iota
	EphemerisTargetVenus
	EphemerisTargetEarth // placeholder for the Earth, included as the values 0...7 match the VSOP planet id's.
	EphemerisTargetMars
	EphemerisTargetJupiter
	EphemerisTargetSaturn
	EphemerisTargetUranus
	EphemerisTargetNeptune
	EphemerisTargetSun // The sun
	EphemerisTargetMoon
	EphemerisTargetOther // Unused, but reserved for custom targets in the future
)

var (
	ephemerisTargetNames = []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
		"Sun", "Moon",
		"Other",
	}
)

func ParseEphemerisTargetType(s string) EphemerisTargetType {
	for i, _ := range ephemerisTargetNames {
		t := EphemerisTargetType(i)
		if t.Is(s) {
			return t
		}
	}
	return EphemerisTargetOther
}

func (t EphemerisTargetType) String() string {
	if t > EphemerisTargetOther {
		return ephemerisTargetNames[EphemerisTargetOther]
	}
	return ephemerisTargetNames[t]
}

func (t EphemerisTargetType) Is(s string) bool {
	return strings.ToLower(ephemerisTargetNames[t]) == strings.ToLower(s)
}

func (t EphemerisTargetType) IsSun() bool {
	return t == EphemerisTargetSun
}

func (t EphemerisTargetType) IsPlanet() bool {
	return t <= EphemerisTargetNeptune
}

func (t EphemerisTargetType) IsOther() bool {
	return t >= EphemerisTargetOther
}
