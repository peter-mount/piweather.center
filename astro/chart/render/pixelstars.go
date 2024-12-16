package render

import (
	"github.com/llgcode/draw2d"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"image/color"
	"math"
)

// PixelStarsRenderer is a StarRenderer which renders stars as a series of pixels.
//
// This StarsRenderer only works for black stars.
//
// The pixel layout is based on an old star chart program for the Apple Macintosh 128 as
// published by Sky and Telescope back in the 1980's, so it looks crude but works for simple charts.
func PixelStarsRenderer(gc draw2d.GraphicContext, s catalogue.Star) {
	gc.BeginPath()
	gc.MoveTo(s.X-0.5, s.Y)
	gc.LineTo(s.X+0.5, s.Y)
	if s.Mag < 3 {
		gc.MoveTo(s.X, s.Y-1)
		gc.LineTo(s.X, s.Y+1)
	}
	if s.Mag < 1 {
		gc.MoveTo(s.X-1, s.Y-1)
		gc.LineTo(s.X+1, s.Y-1)
		gc.LineTo(s.X+1, s.Y+1)
		gc.LineTo(s.X-1, s.Y+1)
	}
	if s.Mag < 0 {
		gc.MoveTo(s.X, s.Y-2)
		gc.LineTo(s.X, s.Y+2)
		gc.MoveTo(s.X-2, s.Y)
		gc.LineTo(s.X+2, s.Y)
	}
	gc.Stroke()
}

// BrightnessPixelStarRenderer is like PixelStarsRenderer except this also sets the colour of the
// star to be grey with White the brightest
func BrightnessPixelStarRenderer(gc draw2d.GraphicContext, s catalogue.Star) {
	// YBSC ranges from -1.46 to 7.96
	m := int(math.Max(-1.46, math.Min(8, s.Mag)) + 1.46)
	m = 0xffff - (m * ((0xffff - 0x8000) / 8))
	col := color.Gray16{Y: uint16(m)}
	gc.Save()
	gc.SetStrokeColor(col)
	gc.BeginPath()
	gc.MoveTo(s.X-0.5, s.Y)
	gc.LineTo(s.X+0.5, s.Y)
	if s.Mag < 3 {
		gc.MoveTo(s.X, s.Y-1)
		gc.LineTo(s.X, s.Y+1)
	}
	gc.Stroke()
	gc.Restore()
}
