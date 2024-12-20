package catalogue

import (
	"github.com/llgcode/draw2d"
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/chart"
)

type EphemerisDayLayer interface {
	chart.ConfigurableLayer
	SetEphemeris(day api.EphemerisDay)
}

type ephemerisDayLayer struct {
	chart.BaseLayer
	chart.BaseProjectionLayer
	renderer  StarRenderer
	ephemeris api.EphemerisDay
	stars     []Star
}

func NewEphemerisDayLayer(renderer StarRenderer, proj chart.Projection) EphemerisDayLayer {
	l := &ephemerisDayLayer{renderer: renderer}
	l.BaseLayer.Drawable = l.draw
	l.BaseProjectionLayer.SetProjection(proj)
	return l
}

func (l *ephemerisDayLayer) SetEphemeris(day api.EphemerisDay) {
	l.ephemeris = day
	l.stars = nil
	l.ephemeris.ForEach(func(result api.EphemerisResult) {
		eq := result.GetEquatorial()
		l.stars = append(l.stars, Star{
			Name: result.Name(),
			P: chart.Point{
				X: eq.RA.Angle(),
				Y: eq.Dec,
			},
			Mag: 0.0,
		})
	})
}

func (l *ephemerisDayLayer) draw(gc draw2d.GraphicContext) {
	for _, s := range l.stars {
		l.renderer(gc, l.Projection(), s)
	}
}
