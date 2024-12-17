package chart

import (
	"github.com/soniakeys/unit"
)

// RaDecAxesLayer generates axes for RA and Dec lines
func RaDecAxesLayer(proj Projection) ConfigurableLayer {
	p := NewPath(proj)

	for x := 0.0; x < 24; x += 1.0 {
		x0 := unit.RAFromHour(x).Angle()
		p.Start()
		for y := -80.0; y <= 80.0; y += .5 {
			p.AddPoint(Point{X: x0, Y: unit.AngleFromDeg(y)})
		}
	}

	for y := -80.0; y <= 80.0; y += 10.0 {
		y0 := unit.AngleFromDeg(y)
		p.Start()
		for x := 0.0; x <= 24.1; x += 0.1 {
			p.AddPoint(Point{X: unit.RAFromHour(x).Angle(), Y: y0})
		}
	}

	return p
}
