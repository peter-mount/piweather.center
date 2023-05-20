package graph

import svg "github.com/ajstarks/svgo/float"

type PolyLine struct {
	x []float64
	y []float64
}

func (p *PolyLine) Add(x, y float64) *PolyLine {
	p.x = append(p.x, x)
	p.y = append(p.y, y)
	return p
}

func (p *PolyLine) AddProjectX(x, y float64, proj *Projection) *PolyLine {
	px, py := proj.Project(x, y)
	if proj.InsideX(px) {
		p.Add(px, py)
	}
	return p
}

// Draw draws a Polyline
func (p *PolyLine) Draw(canvas *svg.SVG, s ...string) {
	// Only plot if we have any data
	if len(p.x) > 0 && len(p.x) == len(p.y) {
		canvas.Polyline(p.x, p.y, s...)
	}
}

// DrawClosed draws a Polygon
func (p *PolyLine) DrawClosed(canvas *svg.SVG, s ...string) {
	// Only plot if we have any data
	if len(p.x) > 0 && len(p.x) == len(p.y) {
		canvas.Polygon(p.x, p.y, s...)
	}
}
