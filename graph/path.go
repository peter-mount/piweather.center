package graph

import (
	svg "github.com/ajstarks/svgo/float"
	"strings"
)

type Path struct {
	p []string
}

func (p *Path) IsEmpty() bool { return len(p.p) == 0 }

func (p *Path) add(op string, x, y float64) *Path {
	p.p = append(p.p, op+Number(x), ",", Number(y))
	return p
}

func (p *Path) MoveTo(x, y float64) *Path {
	return p.add("M", x, y)
}

func (p *Path) LineTo(x, y float64) *Path {
	return p.add("L", x, y)
}

func (p *Path) RelMoveTo(x, y float64) *Path {
	return p.add("m", x, y)
}

func (p *Path) RelLineTo(x, y float64) *Path {
	return p.add("l", x, y)
}

func (p *Path) Draw(canvas *svg.SVG, styles ...string) {
	s := strings.TrimSpace(strings.Join(p.p, ""))
	if s != "" {
		canvas.Path(s, styles...)
	}
}

func (p *Path) Line(x0, y0, x1, y1 float64) *Path {
	return p.MoveTo(x0, y0).LineTo(x1, y1)
}

func (p *Path) Add(x, y float64) *Path {
	if p.IsEmpty() {
		return p.MoveTo(x, y)
	}
	return p.LineTo(x, y)
}

func (p *Path) AddProjectX(x, y float64, proj *Projection) *Path {
	px, py := proj.Project(x, y)
	if proj.InsideX(px) {
		p.Add(px, py)
	}
	return p
}
