package api

import (
	"github.com/peter-mount/go-kernel/v2/util"
	"github.com/peter-mount/piweather.center/store/api"
	util2 "github.com/peter-mount/piweather.center/util"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/globe"
	_ "slices"
	"time"
)

type Ephemeris interface {
	EphemerisCommon
	util.List[EphemerisDay]
	NewDay(t time.Time) EphemerisDay
	Table(EphemerisOption) *api.Table
}

type ephemeris struct {
	ephemerisCommon
	util2.GenericList[EphemerisDay]
}

func NewEphemeris(name string, t time.Time, loc *globe.Coord) Ephemeris {
	e := &ephemeris{
		ephemerisCommon: ephemerisCommon{
			name: name,
			loc:  loc,
		},
	}
	e.ephemerisCommon.init(name, t, loc, coord.NewObliquity(defaultObliquity))

	return e
}
