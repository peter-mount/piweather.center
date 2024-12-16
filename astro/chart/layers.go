package chart

import "github.com/llgcode/draw2d"

// Layers is a collection of Layer which can be drawn in sequence.
type Layers interface {
	ConfigurableLayer
	Add(Layer) Layers
}

func NewLayers() Layers {
	return &layers{}
}

type layers struct {
	BaseLayer
	layers []Layer
}

func (l *layers) Add(layer Layer) Layers {
	l.layers = append(l.layers, layer)
	return l
}

func (l *layers) Draw(ctx draw2d.GraphicContext) {
	l.DrawLayer(ctx, l.draw)
}

func (l *layers) draw(ctx draw2d.GraphicContext) {
	for _, layer := range l.layers {
		layer.Draw(ctx)
	}
}
