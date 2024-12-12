package sidereal

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	util2 "github.com/peter-mount/piweather.center/util/strings"
	"github.com/soniakeys/unit"
	"time"
)

// AtGreenwichMidnight returns the Sidereal Time at Greenwich at 0h UT
func AtGreenwichMidnight(jd julian.Day) unit.Time {
	T := jd.JDMidnight().CenturiesJ2k()
	Theta0 := util.Polynomial(T, 100.46061837, 36000.770053608, 0.000387933, -1/38710000)
	return unit.TimeFromRad(util2.Deg2Rad(util2.DegRange(Theta0)))
}

// FromTime returns the Sidereal time at Greenwich for a specific time.
// Technically this is in UTC however this function ensures the conversion into UTC.
func FromTime(t time.Time) unit.Time {
	return FromJD(julian.FromTime(t))
}

// FromJD returns the Greenwich Sidereal time for a specific Julian Day.
func FromJD(jd julian.Day) unit.Time {
	T := jd.CenturiesJ2k()
	Theta0 := util.Polynomial(T, 280.46061837, 0.000387933, -1/38710000) + (360.98564736629 * (jd.JD() - 2451545.0))
	return unit.TimeFromRad(util2.Deg2Rad(util2.DegRange(Theta0)))
}
