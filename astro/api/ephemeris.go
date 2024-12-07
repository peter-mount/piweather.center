package api

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/unit"
	"strings"
)

// EphemerisResult of a specific object at a specific time.
//
// Note: this replaces the old ephemeris package
type EphemerisResult interface {
	GetAltitude() value.Value
	GetAzimuth() value.Value
	GetDistanceAU() value.Value
	GetJD() julian.Day
	GetName() string
	GetDec() value.Value
	GetRa() value.Value
	SetAltAz(alt, az unit.Angle) EphemerisResult
	SetDistanceAU(f float64) EphemerisResult
	SetEquatorial(ra unit.RA, dec unit.Angle) EphemerisResult
	SetJD(jd julian.Day) EphemerisResult
	SetName(name string) EphemerisResult
	Value(t EphemerisOption) value.Value
}

type ephemerisResult struct {
	jd         julian.Day // time of ephemeris
	name       string
	ra         value.Value // Equatorial coordinates
	dec        value.Value // Equatorial coordinates
	altitude   value.Value // Elevation from local horizon
	azimuth    value.Value // Azimuth, measured westward from the south
	distanceAU value.Value // range in AU
}

func NewEphemerisResult() EphemerisResult {
	return &ephemerisResult{}
}

func (r *ephemerisResult) SetJD(jd julian.Day) EphemerisResult {
	r.jd = jd
	return r
}

func (r *ephemerisResult) GetJD() julian.Day {
	return r.jd
}

func (r *ephemerisResult) SetName(name string) EphemerisResult {
	r.name = name
	return r
}

func (r *ephemerisResult) GetName() string {
	return r.name
}

func (r *ephemerisResult) GetRa() value.Value {
	return r.ra
}

func (r *ephemerisResult) GetDec() value.Value {
	return r.dec
}

func (r *ephemerisResult) GetAltitude() value.Value {
	return r.altitude
}

func (r *ephemerisResult) GetAzimuth() value.Value {
	return r.azimuth
}

func (r *ephemerisResult) GetDistanceAU() value.Value {
	return r.distanceAU
}

func (r *ephemerisResult) SetAltAz(alt, az unit.Angle) EphemerisResult {
	r.altitude = measurement.Degree.Value(alt.Deg())
	r.azimuth = measurement.Degree.Value(az.Deg())
	return r
}

func (r *ephemerisResult) SetEquatorial(ra unit.RA, dec unit.Angle) EphemerisResult {
	r.ra = measurement.RA.Value(ra.Hour())
	r.dec = measurement.Declination.Value(dec.Deg())
	return r
}

func (r *ephemerisResult) SetDistanceAU(f float64) EphemerisResult {
	r.distanceAU = measurement.AU.Value(f)
	return r
}

func (r *ephemerisResult) Value(t EphemerisOption) value.Value {
	switch t {
	case EphemerisOptionAltitude:
		return r.altitude
	case EphemerisOptionAzimuth:
		return r.azimuth
	case EphemerisOptionRa:
		return r.ra
	case EphemerisOptionDec:
		return r.dec
	case EphemerisOptionDistance:
		return r.distanceAU
	default:
		return value.Value{}
	}
}

type EphemerisOption uint8

func (e EphemerisOption) String() string {
	if e >= ephemerisOptionEnd {
		return ephemerisOptionNames[0]
	}
	return ephemerisOptionNames[e]
}

func ParseEphemerisOption(s string) EphemerisOption {
	s = strings.ToLower(strings.TrimSpace(s))
	for i, e := range ephemerisOptionNames {
		if e == s {
			return EphemerisOption(i)
		}
	}
	return ephemerisOptionUnknown
}

const (
	ephemerisOptionUnknown EphemerisOption = iota
	EphemerisOptionAltitude
	EphemerisOptionAzimuth
	EphemerisOptionRa
	EphemerisOptionDec
	EphemerisOptionDistance
	ephemerisOptionEnd // end marker
)

var (
	ephemerisOptionNames = []string{"??", "altitude", "azimuth", "ra", "dec", "distance"}
)
