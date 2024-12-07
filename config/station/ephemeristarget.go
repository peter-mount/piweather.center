package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"sort"
	"strings"
)

// EphemerisTarget defines the target to calculate
type EphemerisTarget struct {
	Pos        lexer.Position
	Target     string                   `parser:"@( 'sun' )"`           // Object to calculate
	Options    []*EphemerisTargetOption `parser:"'(' @@ (',' @@)* ')'"` // Parameters to include
	As         string                   `parser:"( 'as' @String )?"`    // Override the Target for the metric name
	targetType EphemerisTargetType      // Computed target type
}

func (c *visitor[T]) EphemerisTarget(d *EphemerisTarget) error {
	var err error
	if d != nil {
		if c.ephemerisTarget != nil {
			err = c.ephemerisTarget(c, d)
			if util.IsVisitorStop(err) {
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

	d.targetType = EphemerisTargetOther
	if d.Target == "sun" {
		d.targetType = EphemerisTargetSun
	} else {
		for i, e := range ephemerisTargetNames {
			if strings.ToLower(e) == d.Target {
				d.targetType = EphemerisTargetType(i)
			}
		}
	}

	d.As = strings.ToLower(strings.TrimSpace(d.As))
	if d.As == "" {
		d.As = d.Target
	}
	st.ephemerisTarget = d

	// Ensure the options are unique
	opts := make(map[string]*EphemerisTargetOption)
	for _, opt := range d.Options {
		_ = initEphemerisTargetOption(v, opt)

		n := opt.As
		if e, exists := opts[n]; exists {
			return errors.Errorf(opt.Pos, "option %q already declared at %s", e.As, e.Pos.String())
		}
		opts[n] = opt

		// Turn As into a full metric id
		opt.As = st.sensorPrefix + opt.As
	}

	// Sort them, not really needed
	sort.SliceStable(d.Options, func(i, j int) bool {
		return d.Options[i].As < d.Options[j].As
	})

	return util.VisitorStop
}

func (b *builder[T]) EphemerisTarget(f func(Visitor[T], *EphemerisTarget) error) Builder[T] {
	b.ephemerisTarget = f
	return b
}

func (d *EphemerisTarget) GetTarget() EphemerisTargetType {
	return d.targetType
}

type EphemerisTargetType uint8

const (
	EphemerisTargetMercury EphemerisTargetType = iota
	EphemerisTargetVenus
	ephemerisTargetEarth // placeholder for the Earth, included as the values 0...7 match the VSOP planet id's.
	EphemerisTargetMars
	EphemerisTargetJupiter
	EphemerisTargetSaturn
	EphemerisTargetUranus
	EphemerisTargetNeptune
	EphemerisTargetSun   // The sun
	EphemerisTargetOther // Unused, but reserved for custom targets in the future
)

var (
	ephemerisTargetNames = []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
		"Sun",
		"Other",
	}
)

func (t EphemerisTargetType) String() string {
	if t > EphemerisTargetOther {
		return ephemerisTargetNames[EphemerisTargetOther]
	}
	return ephemerisTargetNames[t]
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
