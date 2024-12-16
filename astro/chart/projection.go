package chart

import (
	"github.com/soniakeys/unit"
	"image"
	"math"
)

type Projection interface {
	Bounds() image.Rectangle
	Project(x, y unit.Angle) (float64, float64)
}

// A Point is an X, Y coordinate pair. The axes increase right and down.
//
// Note: This is the same as image.Point except these use unit.Angle
type Point struct {
	X, Y float64
}

func (p Point) GetXY() (int, int) {
	return int(p.X), int(p.Y)
}

type baseProjection struct {
	bounds       image.Rectangle // Bounds of the plot
	lat          unit.Angle      // latitude of observer on Earth
	long         unit.Angle      // longitude of observer on Earth
	R            float64         // radius of the sphere in pixels
	sLat, cLat   float64         // sin & cos of lat
	sLong, cLong float64         // sin & cos of long
}

func (p *baseProjection) Bounds() image.Rectangle {
	return p.bounds
}

func newBaseProjection(long, lat unit.Angle, R float64, bounds image.Rectangle) baseProjection {
	p := baseProjection{
		bounds: bounds,
		long:   long,
		lat:    lat,
		R:      R,
	}
	p.sLong, p.cLong = long.Sincos()
	p.sLat, p.cLat = lat.Sincos()
	return p
}

func NewStereographicProjection(long, lat unit.Angle, R float64, bounds image.Rectangle) Projection {
	return &stereographicProjection{
		baseProjection: newBaseProjection(long, lat, R, bounds),
	}
}

type stereographicProjection struct {
	baseProjection
}

func (prj *stereographicProjection) Project(x, y unit.Angle) (float64, float64) {
	sDx, cDx := math.Sincos((x - prj.long).Rad())
	sY, cY := math.Sincos(y.Rad())
	px := cY * sDx
	py := (prj.cLat * sY) - (prj.sLat * cY * cDx)
	k := (2.0 * prj.R) / (1.0 + (prj.sLat * sY) + (prj.cLat * cY * cDx))
	return k * px, k * py
}

// NewPlainProjection returns a simple Projection where the X and Y axes are plotted as-is
// with the given long at the center
func NewPlainProjection(long unit.Angle, bounds image.Rectangle) Projection {
	return &plainProjection{
		baseProjection: newBaseProjection(long, 0, 1, bounds),
		cx:             float64(bounds.Dx()) / 2.0,
		cy:             float64(bounds.Dy()) / 2.0,
		dx:             float64(bounds.Dx()) / 360.0,
		dy:             float64(bounds.Dy()) / 180.0,
	}
}

type plainProjection struct {
	baseProjection
	cx, cy float64
	dx, dy float64
}

func (prj *plainProjection) Project(x, y unit.Angle) (float64, float64) {
	px := (x - prj.long).Deg()
	if px > 180 {
		px = px - 360
	}
	return px * prj.dx, y.Deg() * prj.dy
}
