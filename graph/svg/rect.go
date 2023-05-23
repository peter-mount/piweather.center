package svg

import "github.com/peter-mount/piweather.center/weather/value"

type Rect struct {
	x0, y0, x1, y1 float64
}

func NewRect(x0, y0, x1, y1 float64) Rect {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rect{x0: x0, y0: y0, x1: x1, y1: y1}
}

func (r Rect) IsZero() bool {
	return value.IsZero(r.x0) && value.IsZero(r.y0) && value.IsZero(r.x1) && value.IsZero(r.y1)
}

func (r Rect) X0() float64 { return r.x0 }

func (r Rect) Y0() float64 { return r.y0 }

func (r Rect) X1() float64 { return r.x1 }

func (r Rect) Y1() float64 { return r.y1 }

func (r Rect) Width() float64 { return r.x1 - r.x0 }

func (r Rect) Height() float64 { return r.y1 - r.y0 }

func (r Rect) Add(x, y float64) Rect {
	return r.include(x, y, x, y)
}

func (r Rect) Include(b Rect) Rect {
	if r.IsZero() {
		return b
	}
	if b.IsZero() {
		return r
	}
	return r.include(b.x0, b.y0, b.x1, b.y1)
}

func (r Rect) include(x0, y0, x1, y1 float64) Rect {
	if !r.IsZero() {
		if r.x0 < x0 {
			x0 = r.x0
		}
		if r.x1 > x1 {
			x1 = r.x1
		}
		if r.y0 < y0 {
			y0 = r.y0
		}
		if r.y1 > y1 {
			y1 = r.y1
		}
	}
	return NewRect(x0, y0, x1, y1)
}

func (r Rect) Contains(x, y float64) bool {
	return value.Within(x, r.x0, r.x1) && value.Within(y, r.y0, r.y1)
}

func (r Rect) Draw(s SVG, a ...string) {
	s.Rect(r.x0, r.y0, r.x1, r.y1, a...)
}

// Reduce reduces the size of the Rect with the specified margins
func (r Rect) Reduce(left, top, right, bottom float64) Rect {
	return NewRect(r.x0+left, r.y0+top, r.x1-right, r.y1-bottom)
}

func (r Rect) Projection() *Projection {
	return NewProjection(r.x0, r.y0, r.x1, r.y1)
}
