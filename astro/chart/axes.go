package chart

import (
	"github.com/soniakeys/unit"
)

// RaDecAxesLayer generates axes for RA and Dec lines
func RaDecAxesLayer(proj Projection) ConfigurableLayer {
	p := NewPath(proj)

	for x := 0.0; x < 24; x += 1.0 {
		p.Start()
		for y := -80.0; y <= 80.0; y += .5 {
			p.Add(p.Project(unit.RAFromHour(x).Angle(), unit.AngleFromDeg(y)))
		}
	}

	for y := -80.0; y <= 80.0; y += 10.0 {
		p.Start()
		for x := 0.0; x <= 24.1; x += 0.1 {
			p.Add(p.Project(unit.RAFromHour(x).Angle(), unit.AngleFromDeg(y)))
		}
	}

	return p
}
