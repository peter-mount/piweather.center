package chart

import (
	"github.com/llgcode/draw2d"
)

type Path interface {
	ConfigurableLayer
	ProjectionLayer

	// Start a new path segment.
	//
	// If a segment is currently being generated it will end that segment and start a new one.
	Start() Path

	// End the current path segment.
	//
	// If no segment is currently being generated this does nothing.
	End() Path

	// Add a point to the current segment.
	//
	// If no segment is being generated a new one will be created.
	//
	// If the point is outside the bounds of the path, then it is ignored. However, if a segment
	// is being generated then it will be completed as the segment is now out of bounds
	Add(x, y float64) Path
	AddPoint(Point) Path
	AddPoints(...Point) Path

	// IsEmpty returns true if there are no segments present.
	IsEmpty() bool

	SetClosed(bool)
}

type path struct {
	BaseLayer
	BaseProjectionLayer
	paths       [][]Point
	currentPath []Point
	closed      bool
}

func NewPath(proj Projection) Path {
	p := &path{}
	p.BaseProjectionLayer.SetProjection(proj)
	p.BaseLayer.Drawable = p.draw
	return p
}

func (p *path) IsEmpty() bool {
	return len(p.paths) == 0
}

func (p *path) SetClosed(closed bool) {
	p.closed = closed
}

func (p *path) Start() Path {
	p.End()
	p.currentPath = nil
	return p
}

func (p *path) End() Path {
	if p.currentPath != nil {
		if p.closed {
			p.currentPath = append(p.currentPath, p.currentPath[0])
		}

		p.paths = append(p.paths, p.currentPath)
	}

	p.currentPath = nil
	return p
}

func (p *path) Add(x, y float64) Path {
	return p.AddPoint(Pt(x, y))
}

func (p *path) AddPoint(pt Point) Path {
	if p.currentPath == nil {
		p.Start()
	}
	// If we are not a closed polygon then we can limit the path to just those points
	// that are contained within the plot.
	// If we filter out when closed then the fill breaks
	if p.closed || p.projection.Contains(pt) {
		p.currentPath = append(p.currentPath, pt)
	} else {
		p.End()
	}
	return p
}

func (p *path) AddPoints(pints ...Point) Path {
	for _, point := range pints {
		p.AddPoint(point)
	}
	return p
}

func (p *path) draw(gc draw2d.GraphicContext) {
	p.End()

	if !p.IsEmpty() {
		for _, points := range p.paths {
			gc.BeginPath()
			gc.MoveTo(p.projection.Project(points[0]))
			for _, point := range points[1:] {
				gc.LineTo(p.projection.Project(point))
			}

			if p.closed {
				gc.FillStroke()
			} else {
				gc.Stroke()
			}
		}
	}
}
