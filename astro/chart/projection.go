package chart

import (
	"github.com/soniakeys/unit"
	"image"
)

type Projection interface {
	// Bounds of the projection - effectively the bounds of the Image
	Bounds() image.Rectangle

	// InsetBounds of the projection - same as Bounds().Index(-50)
	InsetBounds() image.Rectangle

	// GetCenter returns the coordinates of the center of the chart in pixels
	GetCenter() (float64, float64)

	// Contains returns true if Point is within the bounds of this Projection
	Contains(Point) bool

	// Project (x,y) on the sphere to coordinates on the chart.
	//
	// Note: the returned coordinates are positive left and up.
	// These would need to be corrected to the Image which is positive right and down.
	Project(p Point) (float64, float64)

	// Transform returns a new Projection which accepts a different coordinate system on Project.
	//
	// e.g. If we need Horizon coordinates for the chart but we need to project Equatorial
	// coordinates then we need to transform between Equatorial and Ecliptic
	Transform(ProjectionTransform) Projection
}

// A Point is an X, Y coordinate pair. The axes increase right and down.
//
// Note: This is the same as image.Point except these use unit.Angle
type Point struct {
	X, Y unit.Angle
}

func Pt(x, y float64) Point {
	return Point{X: unit.AngleFromDeg(x), Y: unit.AngleFromDeg(y)}
}

func (p Point) GetXY() (int, int) {
	return int(p.X), int(p.Y)
}

type baseProjection struct {
	bounds       image.Rectangle // Bounds of the plot
	cx, cy       float64         // coordinate of center within bounds
	insetBounds  image.Rectangle // bounds expanded by 50 pixels, so we still plot lines crossing the boundary
	lat          unit.Angle      // latitude of observer on Earth
	long         unit.Angle      // longitude of observer on Earth
	R            float64         // radius of the sphere in pixels
	sLat, cLat   float64         // sin & cos of lat
	sLong, cLong float64         // sin & cos of long
}

func (p *baseProjection) Bounds() image.Rectangle {
	return p.bounds
}

func (p *baseProjection) InsetBounds() image.Rectangle {
	return p.insetBounds
}

func newBaseProjection(long, lat unit.Angle, R float64, bounds image.Rectangle) baseProjection {
	p := baseProjection{
		bounds: bounds,
		cx:     float64(bounds.Dx()) / 2.0,
		cy:     float64(bounds.Dy()) / 2.0,
		long:   long,
		lat:    lat,
		R:      R,
	}
	p.insetBounds = image.Rectangle{
		Min: image.Point{X: int(-p.cx) - 50, Y: int(-p.cy) - 50},
		Max: image.Point{X: int(p.cx) + 50, Y: int(p.cy) + 50},
	}
	p.sLong, p.cLong = long.Sincos()
	p.sLat, p.cLat = lat.Sincos()
	return p
}

func (p *baseProjection) GetCenter() (float64, float64) {
	return p.cx, p.cy
}

func NewStereographicProjection(long, lat unit.Angle, R float64, bounds image.Rectangle) Projection {
	return &stereographicProjection{
		baseProjection: newBaseProjection(long, lat, R, bounds),
	}
}

type stereographicProjection struct {
	baseProjection
}

func (prj *stereographicProjection) Contains(p Point) bool {
	x, y := prj.Project(p)
	return image.Pt(int(x), int(y)).In(prj.insetBounds)
}

func (prj *stereographicProjection) Project(p Point) (float64, float64) {
	sDx, cDx := (p.X - prj.long).Sincos()
	sY, cY := p.Y.Sincos()
	px := cY * sDx
	py := (prj.cLat * sY) - (prj.sLat * cY * cDx)
	k := (2.0 * prj.R) / (1.0 + (prj.sLat * sY) + (prj.cLat * cY * cDx))
	return k * px, k * py
}

func (prj *stereographicProjection) Transform(f ProjectionTransform) Projection {
	return TransformProjection(prj, f)
}

// NewPlainProjection returns a simple Projection where the X and Y axes are plotted as-is
// with the given long at the center
func NewPlainProjection(long unit.Angle, bounds image.Rectangle) Projection {
	return &plainProjection{
		baseProjection: newBaseProjection(long, 0, 1, bounds),
		dx:             float64(bounds.Dx()) / 360.0,
		dy:             float64(bounds.Dy()) / 180.0,
	}
}

type plainProjection struct {
	baseProjection
	dx, dy float64
}

func (prj *plainProjection) Contains(p Point) bool {
	x, y := prj.Project(p)
	return image.Pt(int(x), int(y)).In(prj.insetBounds)
}

func (prj *plainProjection) Project(p Point) (float64, float64) {
	px := (p.X - prj.long).Deg()
	if px > 180 {
		px = px - 360
	}
	return px * prj.dx, p.Y.Deg() * prj.dy
}

func (prj *plainProjection) Transform(f ProjectionTransform) Projection {
	return TransformProjection(prj, f)
}

// NewAngularProjection returns an Angular Projection which works for projecting onto the
// view of a fisheye lens (also known as an f-theta lens).
//
// bounds 	of the image in pixels
// R 		radius of fisheye lens in pixels
// alpha	half of the field of view of the fisheye lens
func NewAngularProjection(bounds image.Rectangle, R float64, alpha unit.Angle) Projection {
	return &angularProjection{
		baseProjection: newBaseProjection(0, 0, R/alpha.Rad(), bounds),
		h0:             unit.AngleFromDeg(90),
	}
}

type angularProjection struct {
	baseProjection
	h0 unit.Angle
}

func (prj *angularProjection) Contains(p Point) bool {
	x, y := prj.Project(p)
	return image.Pt(int(x), int(y)).In(prj.insetBounds)
}

func (prj *angularProjection) Project(p Point) (float64, float64) {
	// As we use the origin as the zenith:
	// k = r*H/alpha, x = -k sin(A), y = k cos(A)
	// conversion from h0 is because H is angular distance from zenith whereas p.Y is altitude from horizon
	k := prj.R * float64(prj.h0-p.Y)
	sa, ca := p.X.Sincos()
	return -(k * sa), +(k * ca)
}

func (prj *angularProjection) Transform(f ProjectionTransform) Projection {
	return TransformProjection(prj, f)
}

type ProjectionTransform func(Point) Point

func TransformProjection(proj Projection, f ProjectionTransform) Projection {
	return &transformProjection{
		wrapped:   proj,
		transform: f,
	}
}

type transformProjection struct {
	wrapped   Projection
	transform ProjectionTransform
}

func (t *transformProjection) GetCenter() (float64, float64) {
	return t.wrapped.GetCenter()
}

func (t *transformProjection) Contains(p Point) bool {
	// We must implement Contains here as we need to use the transformed projection and not the wrapped one
	x, y := t.Project(p)
	return image.Pt(int(x), int(y)).In(t.wrapped.InsetBounds())
}

func (t *transformProjection) Project(p Point) (float64, float64) {
	p1 := t.transform(p)
	return t.wrapped.Project(p1)
}

func (t *transformProjection) Bounds() image.Rectangle {
	return t.wrapped.Bounds()
}

func (t *transformProjection) InsetBounds() image.Rectangle {
	return t.wrapped.InsetBounds()
}

func (t *transformProjection) Transform(f ProjectionTransform) Projection {
	return TransformProjection(t, f)
}
