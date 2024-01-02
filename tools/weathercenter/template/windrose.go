package template

import (
	"fmt"
	util2 "github.com/peter-mount/go-anim/util"
	"github.com/peter-mount/piweather.center/store/api"
	"math"
	"sort"
)

type CirclePos struct {
	Deg, Radian float64
	Radius      float64
	X, Y        float64
}

func circlePos(r, a float64) CirclePos {
	d := a * util2.ToRad
	return CirclePos{
		Deg:    a,
		Radian: d,
		Radius: r,
		X:      r * math.Sin(d),
		Y:      r * math.Cos(d),
	}
}

type WindRoseBreakdown struct {
	C1     CirclePos
	C2     CirclePos
	Bucket int
	Entry  int
	Radius float64
}

func (wrb WindRoseBreakdown) Path() string {
	return fmt.Sprintf("M0,0L%.2f,%.2fL%.2f,%.2fz", wrb.C1.X, wrb.C1.Y, wrb.C2.X, wrb.C2.Y)
}

// Seq returns a plot ordering where we order by compass point and then by reverse order
// so the larger values are plotted first ensuring smaller entries are not obscured
func (wrb WindRoseBreakdown) Seq() int {
	return (wrb.Bucket << 8) + (16 - wrb.Entry)
}

func windRoseBreakdown(radius, hubRadius float64, wr *api.WindRose) []WindRoseBreakdown {
	var r []WindRoseBreakdown

	// TODO see why MaxValuePerBucket and v are not the same here?
	//dv := (radius - hubRadius) / wr.MaxValuePerBucket
	mv := 0.0
	for _, b := range wr.Buckets {
		for _, e := range b.MaxVal {
			mv = math.Max(mv, e)
		}
	}

	if mv > 0 {
		dv := (radius - hubRadius) / mv

		for i, b := range wr.Buckets {
			d := float64(i) * 22.5
			for j, v := range b.MaxVal {
				rad := dv * v
				if rad > 0 {
					wrb := WindRoseBreakdown{
						C1:     circlePos(rad+hubRadius, d-8),
						C2:     circlePos(rad+hubRadius, d+8),
						Bucket: i,
						Entry:  j,
						Radius: rad,
					}

					r = append(r, wrb)
				}
			}
		}
	}

	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Seq() < r[j].Seq()
	})

	return r
}
