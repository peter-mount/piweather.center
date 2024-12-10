package api

import (
	"slices"
	"sort"
	"strings"
)

type EphemerisOption uint32

// MetricSuffix returns the suggested suffix for metrics
func (e EphemerisOption) MetricSuffix() string {
	return ephemerisMetrics[e]
}

// Split will return a slice of distinct EphemerisOption's
// This is required as EphemerisOption is a bit mask but there are times we want
// to work against individual entries
func (e EphemerisOption) Split() []EphemerisOption {
	var r []EphemerisOption
	b := EphemerisOption(1)
	for i := 0; i < 32; i++ {
		if e.Is(b) {
			r = append(r, b)
		}
		b = b << 1
	}
	return r
}

func ParseEphemerisOption(s string) EphemerisOption {
	s = strings.ToLower(strings.TrimSpace(s))
	return ephemerisOptionNames[strings.ToLower(strings.TrimSpace(s))]
}

func (e EphemerisOption) Is(o EphemerisOption) bool {
	return (e & o) == o
}

func (e EphemerisOption) IsData() bool {
	return (e & AllDataOptions) == e
}

func (e EphemerisOption) IsCoordinate() bool {
	return (e & AllCoordinateOptions) == e
}

func (e EphemerisOption) GroupSize(o EphemerisOption) int {
	if !e.IsCoordinate() {
		return 1
	}
	return len(e.Split())
}

func (e EphemerisOption) Group() EphemerisOption {
	for _, g := range groupOptions {
		if g.Is(e) {
			return g
		}
	}
	return 0
}

const (
	EquatorialName EphemerisOption = 1 << iota

	EquatorialRA
	EquatorialDec

	// Distance of object
	Distance

	// LightTime to the object from Earth
	LightTime

	DistanceSun

	// SemiDiameter of object
	SemiDiameter

	HorizonAltitude
	// HorizonAzimuth is the astronomical azimuth in degrees, positive west, from due South.
	HorizonAzimuth
	// HorizonBearing is the azimuth in degrees, positive clockwise, from due North.
	HorizonBearing

	EclipticLatitude
	EclipticLongitude

	GalacticLatitude
	GalacticLongitude

	// AllOptions indicates all possible options
	AllOptions = AllDataOptions | AllCoordinateOptions

	// AllDataOptions possible data options
	AllDataOptions = Distance | DistanceSun | LightTime | SemiDiameter

	// AllCoordinateOptions possible coordinate options
	AllCoordinateOptions = Ecliptic | Equatorial | Galactic | Horizon

	// Equatorial coordinates
	Equatorial = EquatorialRA | EquatorialDec

	// Horizon coordinates based on the observers location
	Horizon = HorizonAltitude | HorizonAzimuth | HorizonBearing

	// Ecliptic coordinates
	Ecliptic = EclipticLatitude | EclipticLongitude

	// Galactic coordinates
	Galactic = GalacticLatitude | GalacticLongitude
)

var (
	ephemerisOptionNames = map[string]EphemerisOption{
		"distance":          Distance,
		"lightTime":         LightTime,
		"distanceSun":       DistanceSun,
		"semiDiameter":      SemiDiameter,
		"altitude":          HorizonAltitude,
		"azimuth":           HorizonAzimuth,
		"bearing":           HorizonBearing,
		"equatorial":        Equatorial,
		"ra":                EquatorialRA,
		"dec":               EquatorialDec,
		"ecliptic":          Ecliptic,
		"eclipticLatitude":  EclipticLatitude,
		"eclipticLongitude": EclipticLongitude,
		"galactic":          Galactic,
		"galacticLatitude":  GalacticLatitude,
		"galacticLongitude": GalacticLongitude,
	}

	ephemerisMetrics = map[EphemerisOption]string{
		Distance:          "dist.earth",
		LightTime:         "dist.lightTime",
		DistanceSun:       "dist.sun",
		SemiDiameter:      "semiDiameter",
		HorizonAltitude:   "hz.altitude",
		HorizonAzimuth:    "hz.azimuth",
		HorizonBearing:    "hz.bearing",
		EquatorialRA:      "eq.ra",
		EquatorialDec:     "eq.dec",
		EclipticLatitude:  "ecl.lat",
		EclipticLongitude: "ecl.lon",
		GalacticLatitude:  "gal.lat",
		GalacticLongitude: "gal.lon",
	}

	availableOptions []EphemerisOption

	groupOptions = []EphemerisOption{
		Equatorial,
		Horizon,
		AllDataOptions,
		Ecliptic,
		Galactic,
	}
)

func (e EphemerisOption) String() string {
	n, ok := e.Name()
	if !ok {
		return ""
	}
	return n
}

func (e EphemerisOption) Name() (string, bool) {
	n, exists, _ := e.nameGroup()
	return n, exists
}

func (e EphemerisOption) IsGroup() bool {
	_, _, isGroup := e.nameGroup()
	return isGroup
}

func (e EphemerisOption) nameGroup() (string, bool, bool) {
	switch e {
	case Ecliptic:
		return "Ecliptic", true, true
	case Equatorial:
		return "Equatorial", true, true
	case Galactic:
		return "Galactic", true, true
	case Horizon:
		return "Horizon", true, true
	default:
		s, exists := ephemerisMetrics[e]
		return s, exists, false
	}
}

func init() {
	for k, _ := range ephemerisMetrics {
		availableOptions = append(availableOptions, k)
	}
	sort.SliceStable(availableOptions, func(i, j int) bool {
		return availableOptions[i] < availableOptions[j]
	})
}

func AvailableOptions() []EphemerisOption {
	return slices.Clone(availableOptions)
}

func (e EphemerisOption) OptionNames() []string {
	var keys []EphemerisOption
	for _, k := range availableOptions {
		if e.Is(k) {
			keys = append(keys, k)
		}
	}

	var r []string
	for _, k := range keys {
		if e.Is(k) {
			r = append(r, ephemerisMetrics[k])
		}
	}
	return r
}

func GroupOptions() []EphemerisOption {
	return slices.Clone(groupOptions)
}

func (e EphemerisOption) Groups() []EphemerisOption {
	var r []EphemerisOption
	for _, g := range groupOptions {
		if g.Is(e) {
			r = append(r, g)
		}
	}
	return r
}
