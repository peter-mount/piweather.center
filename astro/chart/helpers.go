package chart

import (
	"github.com/llgcode/draw2d"
)

// FloodFillLayer is a Layer which will fill the entire chart with the Fill colour.
func FloodFillLayer(proj Projection) ConfigurableLayer {
	return NewDrawableLayer(func(gc draw2d.GraphicContext) {
		b := proj.Bounds()
		gc.BeginPath()
		gc.MoveTo(0, 0)
		gc.LineTo(float64(b.Dx()-1), 0)
		gc.LineTo(float64(b.Dx()-1), float64(b.Dy()-1))
		gc.LineTo(0, float64(b.Dy()-1))
		gc.Close()
		gc.FillStroke()
	})
}
