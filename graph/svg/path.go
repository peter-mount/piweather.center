package svg

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strings"
)

// Path allows for a SVG Path element to be constructed.
type Path struct {
	p            []string // Path being generated
	x, y         float64  // Absolute coordinates of last position
	optimalCount int      // Used to limit how many optimisations are done
}

// IsEmpty returns true of no points have been added to the Path
func (p *Path) IsEmpty() bool { return len(p.p) == 0 }

func (p *Path) add(op string, x, y float64) *Path {
	p.p = append(p.p, op+Number(x), ",", Number(y))
	return p
}

// MoveTo an absolute coordinate
func (p *Path) MoveTo(x, y float64) *Path {
	// Do nothing if the move is to the same point
	if !p.IsEmpty() && value.Equal(p.x, x) && value.Equal(p.y, y) {
		return p
	}
	p.x, p.y, p.optimalCount = x, y, 0
	return p.add("M", x, y)
}

// LineTo draws a line to an absolute coordinate.
// If the Path is empty this will be changed to a MoveTo
func (p *Path) LineTo(x, y float64) *Path {
	// If first point then use MoveTo
	if p.IsEmpty() {
		return p.MoveTo(x, y)
	}

	// Do nothing if the move is to the same point
	if value.Equal(p.x, x) && value.Equal(p.y, y) {
		return p
	}
	p.x, p.y, p.optimalCount = x, y, 0
	return p.add("L", x, y)
}

// RelMoveTo moves the drawing point by a relative coordinate from the current point
func (p *Path) RelMoveTo(x, y float64) *Path {
	// Ignore 0,0 relative moves
	if value.Equal(x, 0) && value.Equal(y, 0) {
		return p
	}

	p.x, p.y = p.x+x, p.y+y
	return p.add("m", x, y)
}

// RelLineTo draws a line to a relative coordinate from the current point.
func (p *Path) RelLineTo(x, y float64) *Path {
	// Ignore 0,0 relative lines
	if value.Equal(x, 0) && value.Equal(y, 0) {
		return p
	}

	// If first point then use RelMoveTo
	if p.IsEmpty() {
		return p.RelMoveTo(x, y)
	}

	p.x, p.y = p.x+x, p.y+y
	return p.add("l", x, y)
}

// OptimalMoveTo will use RelMoveTo if the distance from the last point
// is smaller than x,y. If not it does a MoveTo.
// Using this can compress the path's output.
func (p *Path) OptimalMoveTo(x, y float64) *Path {
	return p.optimalTo(x, y, p.MoveTo, p.RelMoveTo)
}

// OptimalLineTo will use RelLineTo if the distance from the last point
// is smaller than x,y. If not it does a LineTo.
// Using this can compress the path's output.
func (p *Path) OptimalLineTo(x, y float64) *Path {
	return p.optimalTo(x, y, p.LineTo, p.RelLineTo)
}

func (p *Path) optimalTo(x, y float64, abs, rel func(float64, float64) *Path) *Path {
	// Only use rel if we already have a point
	if !p.IsEmpty() {
		dx, dy := x-p.x, y-p.y
		if math.Abs(dx) < math.Abs(x) && math.Abs(dy) < math.Abs(y) {
			// If we have 10 successive optimisations then skip and keep the
			// absolute coordinate. This keeps the axes accurate as, over time
			// errors happen due to the limited precision in the output
			p.optimalCount++
			if p.optimalCount < 10 {
				return rel(dx, dy)
			}
		}
	}
	p.optimalCount = 0
	return abs(x, y)
}

// Draw implements Drawable to allow the Path to be drawn using SVG.Draw()
func (p *Path) Draw(svg SVG, styles ...string) {
	s := strings.TrimSpace(strings.Join(p.p, ""))
	if s != "" {
		svg.Tag("path", nil, AttrMerge(styles, Attr("d", s))...)
	}
}

// Line draws a line between two coordinates.
// The line will be optimized so if the start is the current position then
// it will not include an initial move, whilst the line to the
// destination will be changed to a relative line if it's nearby.
func (p *Path) Line(x0, y0, x1, y1 float64) *Path {
	return p.OptimalMoveTo(x0, y0).OptimalLineTo(x1, y1)
}

// Add a point to the path.
// If the path is empty then this will just add a MoveTo.
// Otherwise, it will use LineTo. It's handy for generating lines.
func (p *Path) Add(x, y float64) *Path {
	if p.IsEmpty() {
		return p.MoveTo(x, y)
	}
	return p.OptimalLineTo(x, y)
}

// AddProjectX is the same as Add except it will use a Projection first.
func (p *Path) AddProjectX(x, y float64, proj *Projection) *Path {
	px, py := proj.Project(x, y)
	return p.Add(px, py)
}
