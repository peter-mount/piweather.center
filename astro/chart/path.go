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

	// IsEmpty returns true if there are no segments present.
	IsEmpty() bool
}

type path struct {
	BaseLayer
	BaseProjectionLayer
	paths       []*draw2d.Path
	currentPath *draw2d.Path
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

func (p *path) Start() Path {
	p.End()
	p.currentPath = &draw2d.Path{}
	return p
}

func (p *path) End() Path {
	if p.currentPath != nil && !p.currentPath.IsEmpty() {
		p.paths = append(p.paths, p.currentPath)
	}
	p.currentPath = nil
	return p
}

func (p *path) Add(x, y float64) Path {
	if p.currentPath == nil {
		p.Start()
	}
	if p.Contains(x, y) {
		if p.currentPath.IsEmpty() {
			p.currentPath.MoveTo(x, y)
		} else {
			p.currentPath.LineTo(x, y)
		}
	} else {
		p.End()
	}
	return p
}

func (p *path) draw(gc draw2d.GraphicContext) {
	p.End()

	if !p.IsEmpty() {
		gc.Stroke(p.paths...)
	}
}
