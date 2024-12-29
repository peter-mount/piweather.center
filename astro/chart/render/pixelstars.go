package render

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/chart"
	"image/color"
	"math"
)

// PixelStarsRenderer is a StarRenderer which renders stars as a series of pixels.
//
// This StarsRenderer only works for black stars.
//
// The pixel layout is based on an old star chart program for the Apple Macintosh 128 as
// published by Sky and Telescope back in the 1980's, so it looks crude but works for simple charts.
func PixelStarsRenderer(gc draw2d.GraphicContext, p chart.Projection, s catalogue.Star) {
	x, y := p.Project(s.P)

	gc.BeginPath()
	gc.MoveTo(x-0.5, y)
	gc.LineTo(x+0.5, y)
	if s.Mag < 3 {
		gc.MoveTo(x, y-1)
		gc.LineTo(x, y+1)
	}
	if s.Mag < 1 {
		gc.MoveTo(x-1, y-1)
		gc.LineTo(x+1, y-1)
		gc.LineTo(x+1, y+1)
		gc.LineTo(x-1, y+1)
	}
	if s.Mag < 0 {
		gc.MoveTo(x, y-2)
		gc.LineTo(x, y+2)
		gc.MoveTo(x-2, y)
		gc.LineTo(x+2, y)
	}
	gc.Stroke()
}

// BrightnessPixelStarRenderer is like PixelStarsRenderer except this also sets the colour of the
// star to be grey with White the brightest
func BrightnessPixelStarRenderer(gc draw2d.GraphicContext, p chart.Projection, s catalogue.Star) {
	// YBSC ranges from -1.46 to 7.96
	m := int(math.Max(-1.46, math.Min(8, s.Mag)) + 1.46)
	m = 0xffff - (m * ((0xffff - 0x8000) / 8))
	col := color.Gray16{Y: uint16(m)}

	x, y := p.Project(s.P)

	gc.Save()
	gc.SetStrokeColor(col)
	gc.BeginPath()
	gc.MoveTo(x-0.5, y)
	gc.LineTo(x+0.5, y)
	if s.Mag < 3 {
		gc.MoveTo(x, y-1)
		gc.LineTo(x, y+1)
	}
	gc.Stroke()
	gc.Restore()
}

func SizeStarRenderer(gc draw2d.GraphicContext, p chart.Projection, s catalogue.Star) {
	x, y := p.Project(s.P)

	r := 1 + math.Min(0, 2*(4.0-s.Mag))
	gc.BeginPath()
	draw2dkit.Circle(gc, x, y, r)
	gc.Fill()

	if s.Name != "" {
		gc.Save()
		gc.Translate(x+5+r, y-5-r)
		gc.Scale(-1, 1)
		gc.FillString(s.Name)
		gc.Restore()
	}
}
