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
	HorizonAltitude = EphemerisOption(0x0010000)
	// HorizonAzimuth is the astronomical azimuth in degrees, positive west, from due South.
	HorizonAzimuth = EphemerisOption(0x0020000)
	// HorizonBearing is the azimuth in degrees, positive clockwise, from due North.
	HorizonBearing = EphemerisOption(0x0040000)

	// Equatorial coordinates
	Equatorial    = EquatorialRA | EquatorialDec
	EquatorialRA  = EphemerisOption(0x0080000)
	EquatorialDec = EphemerisOption(0x0100000)

	// Ecliptic coordinates
	Ecliptic          = EclipticLatitude | EclipticLongitude
	EclipticLatitude  = EphemerisOption(0x0200000)
	EclipticLongitude = EphemerisOption(0x0400000)

	// Galactic coordinates
	Galactic          = GalacticLatitude | GalacticLongitude
	GalacticLatitude  = EphemerisOption(0x0800000)
	GalacticLongitude = EphemerisOption(0x1000000)

	// There's room for 7 more coordinate pairs after this point
)

var (
	ephemerisOptionNames = map[string]EphemerisOption{
		"distance":          Distance,
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
		Distance:          "distance",
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
)
