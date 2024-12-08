package api

import (
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
	ei := int(e)
	b := 1
	for i := 0; i < 32; i++ {
		if (ei & b) == b {
			r = append(r, EphemerisOption(b))
		}
		b = b << 1
	}
	return r
}

func ParseEphemerisOption(s string) EphemerisOption {
	s = strings.ToLower(strings.TrimSpace(s))
	return ephemerisOptionNames[strings.ToLower(strings.TrimSpace(s))]
}

const (
	// AllOptions indicates all possible options
	AllOptions = EphemerisOption(0xffffffff)

	// Distance of object
	Distance = EphemerisOption(0x0001)

	// Horizon coordinates based on the observers location
	Horizon         = HorizonAltitude | HorizonAzimuth
	HorizonAltitude = EphemerisOption(0x010000)
	HorizonAzimuth  = EphemerisOption(0x020000)

	// Equatorial coordinates
	Equatorial    = EquatorialRA | EquatorialDec
	EquatorialRA  = EphemerisOption(0x040000)
	EquatorialDec = EphemerisOption(0x080000)

	// Ecliptic coordinates
	Ecliptic          = EclipticLatitude | EclipticLongitude
	EclipticLatitude  = EphemerisOption(0x100000)
	EclipticLongitude = EphemerisOption(0x200000)

	// Galactic coordinates
	Galactic          = GalacticLatitude | GalacticLongitude
	GalacticLatitude  = EphemerisOption(0x400000)
	GalacticLongitude = EphemerisOption(0x800000)

	// There's room for 8 more coordinate pairs after this point
)

var (
	ephemerisOptionNames = map[string]EphemerisOption{
		"distance":          Distance,
		"altitude":          HorizonAltitude,
		"azimuth":           HorizonAzimuth,
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
		Distance:          "distance",
		HorizonAltitude:   "hz.altitude",
		HorizonAzimuth:    "hz.azimuth",
		EquatorialRA:      "eq.ra",
		EquatorialDec:     "eq.dec",
		EclipticLatitude:  "ecl.lat",
		EclipticLongitude: "ecl.lon",
		GalacticLatitude:  "gal.lat",
		GalacticLongitude: "gal.lon",
	}
)
