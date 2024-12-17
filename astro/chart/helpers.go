package chart

import (
	"github.com/llgcode/draw2d"
)

// FloodFillLayer is a Layer which will fill the entire chart with the Fill colour.
func FloodFillLayer(proj Projection) ConfigurableLayer {
	return borderLayer(proj, func(gc draw2d.GraphicContext) {
		gc.Fill()
	})
}

// BorderLayer is a Layer which will draw a border around the entire chart with the Stroke colour.
func BorderLayer(proj Projection) ConfigurableLayer {
	return borderLayer(proj, func(gc draw2d.GraphicContext) {
		gc.Stroke()
	})
}

// borderLayer used by FloodFillLayer and BorderLayer
func borderLayer(proj Projection, drawable Drawable) ConfigurableLayer {
	return NewDrawableLayer(func(gc draw2d.GraphicContext) {
		b := proj.Bounds()
		x0, y0 := -1.0, -1.0
		x1, y1 := float64(b.Dx()), float64(b.Dy())
		gc.BeginPath()
		gc.MoveTo(x0, y0)
		gc.LineTo(x1, y0)
		gc.LineTo(x1, y1)
		gc.LineTo(x0, y1)
		gc.Close()
		drawable(gc)
	})
}
