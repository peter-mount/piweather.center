package api

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sidereal"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
	"time"
)

const (
	defaultObliquity = 23.4392911
)

type EphemerisCommon interface {
	// Name of this entry
	Name() string

	// Time of this entry
	Time() time.Time

	// JD julian day number of this entry
	JD() julian.Day

	// SiderialTime of this entry
	SiderialTime() unit.Time

	// GetObliquity obliquity of the Ecliptic on this time
	GetObliquity() *coord.Obliquity
}
type ephemerisCommon struct {
	name      string           // name of object
	time      time.Time        // time of result
	jd        julian.Day       // date of result
	siderial  unit.Time        // siderial time
	loc       *globe.Coord     // Location of observer
	obliquity *coord.Obliquity // Obliquity of ecliptic on date
}

func (e *ephemerisCommon) init(name string, t time.Time, loc *globe.Coord, ε *coord.Obliquity) {
	e.name = name
	e.time = t
	e.jd = julian.FromTime(t)
	e.siderial = sidereal.FromJD(e.jd)
	e.loc = loc
	e.obliquity = ε
}

func (e *ephemerisCommon) Name() string {
	return e.name
}

func (e *ephemerisCommon) Time() time.Time {
	return e.time
}

func (e *ephemerisCommon) JD() julian.Day {
	return e.jd
}

func (e *ephemerisCommon) SiderialTime() unit.Time {
	return e.siderial
}

func (e *ephemerisCommon) GetObliquity() *coord.Obliquity {
	return e.obliquity
}

func (e *ephemerisCommon) setObliquity(ε unit.Angle) {
	e.obliquity = coord.NewObliquity(ε)
}
