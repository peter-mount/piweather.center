package ephemeris

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
	"testing"
)

func TestSunProvider(t *testing.T) {
	london := coord.LatLong{
		Name:     "London, England",
		Coord:    globe.Coord{Lat: unit.AngleFromDeg(51.51), Lon: unit.AngleFromDeg(-0.13)},
		Altitude: 113.13,
	}

	type args struct {
		start julian.Day
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "2023_Jan_Wk1",
			args: args{
				start: julian.FromDate(2023, 1, 1, 0, 0, 0),
				count: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := &Ephemeris{Name: "Sun test: " + tt.name}

			ep.Meta.LatLong = london

			err := ep.Include(tt.args.start).
				AppendDuration(float64(tt.args.count)).
				Generate(1, &SunProvider{})

			if err != nil {
				t.Errorf("Error %v", err)
			} else {
				fmt.Println(util.String(ep))
			}
		})
	}
}
