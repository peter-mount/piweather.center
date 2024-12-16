package chart

import (
	"github.com/llgcode/draw2d"
	"image"
)

// PixelStarsRenderer is a StarRenderer which renders stars as a series of pixels.
//
// This StarsRenderer only works for black stars.
//
// The pixel layout is based on an old star chart program for the Apple Macintosh 128 as
// published by Sky and Telescope back in the 1980's, so it looks crude but works for simple charts.
func PixelStarsRenderer(gc draw2d.GraphicContext, p Star) {
	i := 3
	switch {
	case p.Mag < 3:
		i = 2
	case p.Mag < 1:
		i = 1
	case p.Mag < 0:
		i = 0
	}
	gc.Save()
	gc.Translate(p.X, p.Y)
	gc.DrawImage(pixelStars[i])
	gc.Restore()
}

var (
	pixelStars []image.Image
)

func init() {
	r := image.Rect(-2, -2, 2, 2)
	for m := 0; m < 4; m++ {
		img := image.NewRGBA(r)
		img.Set(0, 0, image.Black)
		if m < 3 {
			img.Set(0, -1, image.Black)
			img.Set(0, +1, image.Black)
			img.Set(-1, 0, image.Black)
			img.Set(+1, 0, image.Black)
		}
		if m < 1 {
			img.Set(-1, -1, image.Black)
			img.Set(+1, -1, image.Black)
			img.Set(-1, +1, image.Black)
			img.Set(+1, +1, image.Black)
		}
		if m < 0 {
			img.Set(0, -2, image.Black)
			img.Set(0, +2, image.Black)
			img.Set(-2, 0, image.Black)
			img.Set(+2, 0, image.Black)
		}
		pixelStars = append(pixelStars, img)
	}
}
