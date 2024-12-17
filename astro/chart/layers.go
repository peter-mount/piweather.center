package chart

import "github.com/llgcode/draw2d"

// Layers is a collection of Layer which can be drawn in sequence.
type Layers interface {
	ConfigurableLayer

	// Add a Layer to this Layers collection
	Add(Layer) Layers

	// Flip will flip the axes when drawing
	Flip(x, y bool) Layers

	SetProjection(Projection) Layers

	GetProjection() Projection
}

func NewLayers() Layers {
	return &layers{
		scaleX: 1.0,
		scaleY: 1.0,
	}
}

type layers struct {
	BaseLayer
	layers         []Layer
	scaleX, scaleY float64
	projection     Projection
}

func (l *layers) Add(layer Layer) Layers {
	if layer == nil {
		panic("cannot add nil layer")
	}
	l.layers = append(l.layers, layer)
	return l
}

func (l *layers) Flip(x, y bool) Layers {
	l.scaleX = 1
	l.scaleY = 1
	if x {
		l.scaleX = -1
	}
	if y {
		l.scaleY = -1
	}
	return l
}

func (l *layers) GetProjection() Projection {
	return l.projection
}

func (l *layers) SetProjection(projection Projection) Layers {
	l.projection = projection
	return l
}

func (l *layers) Draw(ctx draw2d.GraphicContext) {
	l.DrawLayer(ctx, l.draw)
}

func (l *layers) draw(ctx draw2d.GraphicContext) {
	// If we have a projection, translate to the origin
	if l.projection != nil {
		ctx.Translate(l.projection.GetCenter())
	}

	// Handle axis flipping
	ctx.Scale(l.scaleX, l.scaleY)

	for _, layer := range l.layers {
		layer.Draw(ctx)
	}
}
