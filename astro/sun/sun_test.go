package sun

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/rise"
	"github.com/soniakeys/unit"
	"testing"
)

func TestApparentEquatorial(t *testing.T) {
	type args struct {
		jd julian.Day
	}
	tests := []struct {
		name string
		args args
		want coord.Equatorial
	}{
		{
			name: "1992 10 13",
			args: args{jd: 2448908.5},
			want: coord.New(
				unit.NewAngle('+', 13, 13, 31).Mul(15.0).RA(),
				unit.NewAngle('-', 7, 47, 6),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApparentEquatorial(tt.args.jd); !got.Equals(&tt.want) {
				fmt.Println(util.String(tt.want))
				fmt.Println(util.String(got))
				t.Errorf("ApparentEquatorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApparentEquatorial_RiseSet(t *testing.T) {
	london := &coord.LatLong{
		Name:     "London, England",
		Coord:    globe.Coord{Lat: unit.AngleFromDeg(51.51), Lon: unit.AngleFromDeg(-0.13)},
		Altitude: 113.13,
	}

	jd20230107 := julian.FromDate(2023, 1, 7, 0, 0, 0)

	type Args struct {
		H0     unit.Angle
		Fields coord.Equatorial
		Loc    *coord.LatLong
		JD     julian.Day
	}
	type Test struct {
		Name string `xml:"name,attr"`
		Args Args
		Want coord.RiseSet
	}
	tests := []Test{
		{
			Name: "Sun_20230106_london",
			Args: Args{
				H0:     rise.Stdh0Stellar,
				Fields: ApparentEquatorial(jd20230107),
				Loc:    london,
				JD:     jd20230107,
			},
			Want: coord.RiseSet{
				Rise:    unit.TimeFromDay(.33746237546923313),
				Transit: unit.TimeFromDay(.5037688659344396),
				Set:     unit.TimeFromDay(.670075356399646),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Args.Fields.RiseSet(tt.Args.Loc.Coord, tt.Args.JD.Apparent(), tt.Args.H0)
			if !got.Equals(&tt.Want) {
				fmt.Println(util.String(&tt))
				fmt.Println(util.String(&got), got.Rise.Day(), got.Transit.Day(), got.Set.Day())
				t.Errorf("RiseSet() = %s, want %s", got.String(), tt.Want.String())
			}
		})
	}
}
