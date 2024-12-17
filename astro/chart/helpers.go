package chart

import (
	"github.com/llgcode/draw2d"
	"github.com/soniakeys/unit"
	"image"
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

// HorizonLayer is a Layer which fills everything below the "horizon" - e.g. negative Y
func HorizonLayer(proj Projection) ConfigurableLayer {
	cx, cy := float64(proj.Bounds().Dx()>>1), float64(proj.Bounds().Dy()>>1)
	correct := func(x0, y0 unit.Angle) (float64, float64) {
		x, y := proj.Project(x0, y0)
		return cx - x, cy - y
	}

	return NewDrawableLayer(func(gc draw2d.GraphicContext) {
		xStep := unit.RAFromHour(1).Angle()
		for y := -80.0; y <= 0.0; y += 10.0 {
			y0 := unit.AngleFromDeg(y)
			y1 := y0 - unit.AngleFromDeg(10)

			x0 := unit.RAFromHour(0).Angle()
			for x := 0.0; x < 24; x++ {
				x1 := x0 + xStep

				visible := false
				var pt []float64
				for z := 0.0; z <= 1.0; z += 0.2 {
					px1, py1 := correct(x0+xStep.Mul(z), y0)
					visible = visible || image.Pt(int(px1), int(py1)).In(proj.Bounds())
					pt = append(pt, px1, py1)
				}
				for z := 1.0; z >= 0.0; z -= 0.2 {
					px1, py1 := correct(x0+xStep.Mul(z), y1)
					visible = visible || image.Pt(int(px1), int(py1)).In(proj.Bounds())
					pt = append(pt, px1, py1)
				}
				if visible {
					gc.BeginPath()
					gc.MoveTo(pt[0], pt[1])
					for i := 2; i < len(pt); i += 2 {
						gc.LineTo(pt[i], pt[i+1])
					}
					gc.Close()
					gc.FillStroke()
				}

				x0 = x1
			}
		}
	})
}
