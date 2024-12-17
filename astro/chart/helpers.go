package chart

import (
	"github.com/llgcode/draw2d"
	"github.com/soniakeys/unit"
)

// FloodFillLayer is a Layer which will fill the entire chart with the Fill colour.
func FloodFillLayer(proj Projection) ConfigurableLayer {
	return borderLayer(proj, func(gc draw2d.GraphicContext) {
		gc.Fill()
	})
}

// HorizonLayer is a Layer which fills everything below the "horizon" with a fill.
//
// Note: This specific layer only works for HorizonProjection showing the entire sky.
//
// If the projection is not centered around the zenith then this will not generate a correct horizon
// due to it assuming the horizon is visible in its entirety.
func HorizonLayer(proj Projection) ConfigurableLayer {
	return borderLayer(proj, func(gc draw2d.GraphicContext) {
		// Up to this point we have the entire chart as the outer region.
		// Now add a hole for the horizon

		zero := unit.AngleFromDeg(0)
		xStep := unit.AngleFromDeg(1)
		for x := zero; x <= unit.AngleFromDeg(360); x += xStep {
			pt := Point{X: x, Y: zero}
			px, py := proj.Project(pt)
			if x == zero {
				gc.MoveTo(px, py)
			} else {
				gc.LineTo(px, py)
			}
		}

		gc.Close()
		gc.FillStroke()
	})
}

// borderLayer used by BorderLayer, FloodFillLayer and HorizonLayer
func borderLayer(proj Projection, drawable Drawable) ConfigurableLayer {
	return NewDrawableLayer(func(gc draw2d.GraphicContext) {
		b := proj.Bounds()
		cx, cy := float64(b.Dx())/2, float64(b.Dy())/2
		x0, y0 := cx, cy
		x1, y1 := -cx, -cy
		gc.BeginPath()
		gc.MoveTo(x0, y0)
		gc.LineTo(x1, y0)
		gc.LineTo(x1, y1)
		gc.LineTo(x0, y1)
		gc.Close()
		drawable(gc)
	})
}
