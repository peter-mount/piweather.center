package api

import (
	"github.com/peter-mount/go-kernel/v2/util"
	util2 "github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/unit"
	"time"
)

type EphemerisDay interface {
	EphemerisCommon
	util.List[EphemerisResult]
	SetObliquity(ε unit.Angle) EphemerisDay
	NewResult(name string) EphemerisResult
}

type ephemerisDay struct {
	ephemerisCommon
	util2.GenericList[EphemerisResult]
}

func NewEphemerisDay(name string, t value.Time) EphemerisDay {
	r := &ephemerisDay{}
	r.ephemerisCommon.init(name, t.Time(), t.Location(), coord.NewObliquity(defaultObliquity))
	return r
}

func (e *ephemeris) NewDay(t time.Time) EphemerisDay {
	r := &ephemerisDay{}
	r.ephemerisCommon.init("", t, e.ephemerisCommon.loc, e.ephemerisCommon.obliquity)
	e.Add(r)
	return r
}

func newEphemerisDay(name string, common *ephemerisCommon) *ephemerisDay {
	r := &ephemerisDay{}
	if common != nil {
		r.ephemerisCommon = *common
	}
	r.ephemerisCommon.name = name
	return r
}

func (r *ephemerisDay) SetObliquity(ε unit.Angle) EphemerisDay {
	r.ephemerisCommon.setObliquity(ε)
	return r
}

func (r *ephemerisDay) NewResult(name string) EphemerisResult {
	e := newEphemerisResult(name, &r.ephemerisCommon)
	r.Add(e)
	return e
}
