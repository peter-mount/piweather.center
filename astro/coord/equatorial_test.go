package coord

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/rise"
	"github.com/soniakeys/unit"
	"testing"
)

func TestEquatorial_RiseSet(t *testing.T) {
	boston := &LatLong{
		Name:  "Boston, MA",
		Coord: globe.Coord{Lon: unit.AngleFromDeg(71.0833), Lat: unit.AngleFromDeg(42.3333)},
	}

	london := &LatLong{
		Name:     "London, England",
		Coord:    globe.Coord{Lat: unit.AngleFromDeg(51.51), Lon: unit.AngleFromDeg(-0.13)},
		Altitude: 113.13,
	}

	type Args struct {
		H0     unit.Angle
		Fields Equatorial
		Loc    *LatLong
		JD     julian.Day
	}
	type Test struct {
		Name string `xml:"name,attr"`
		Args Args
		Want RiseSet
	}
	tests := []Test{
		{
			Name: "Circumpolar",
			Args: Args{
				H0:     rise.Stdh0Stellar,
				Fields: New(unit.RAFromDeg(41.73129), unit.AngleFromDeg(88.44092)),
				Loc:    london,
				JD:     julian.FromDate(1988, 3, 20, 0, 0, 0),
			},
			Want: RiseSet{Circumpolar: true},
		},
		{
			Name: "Venus_19880320_boston",
			Args: Args{
				H0:     rise.Stdh0Stellar,
				Fields: New(unit.RAFromDeg(41.73129), unit.AngleFromDeg(18.44092)),
				Loc:    boston,
				JD:     julian.FromDate(1988, 3, 20, 0, 0, 0),
			},
			Want: RiseSet{
				Rise:    unit.TimeFromDay(.518161679226393),
				Transit: unit.TimeFromDay(.8196459024007566),
				Set:     unit.TimeFromDay(.12113012557512),
			},
		},
		{
			Name: "Venus_20230106_london",
			Args: Args{
				H0: rise.Stdh0Stellar,
				Fields: New(
					unit.NewAngle('+', 20, 34, 21).Mul(15.0).RA(),
					unit.NewAngle('-', 20, 16, 14),
				),
				Loc: london,
				JD:  julian.FromDate(2023, 1, 7, 0, 0, 0),
			},
			Want: RiseSet{
				Rise:    unit.TimeFromDay(.38539004637161817),
				Transit: unit.TimeFromDay(.5615411587907412),
				Set:     unit.TimeFromDay(.737692271209864),
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
