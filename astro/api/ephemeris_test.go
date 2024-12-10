package api

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
	"github.com/soniakeys/unit"
	"strings"
	"testing"
	"time"
)

func TestNewEphemeris(t *testing.T) {

	earth, err := planetposition.LoadPlanetPath(planetposition.Earth, "/home/peter/area51/piweather.center/builds/linux/amd64/share/vsop87b")
	if err != nil {
		t.Fatalf("failed to find planet Earth %v", err)
	}

	loc := coord.LatLong{
		Longitude: unit.AngleFromDeg(0),
		Latitude:  unit.AngleFromDeg(51.5),
		Altitude:  0,
	}

	tm := time.Date(2024, 12, 1, 12, 0, 0, 0, time.UTC)
	e := NewEphemeris("test", tm, loc.Coord())

	for day := 0; day < 7; day++ {
		dt := tm.AddDate(0, 0, day)

		jd := julian.FromTime(dt)
		ra, dec, R := solar.ApparentEquatorialVSOP87(earth, jd.Float())

		e.NewDay(dt).
			NewResult("sun").
			SetEquatorial(ra, dec).
			SetDistance(measurement.AU.Value(R))
	}

	table := e.Table(AllOptions).
		Finalise().
		String(nil)

	fmt.Println(strings.Join(table, "\n"))
}
