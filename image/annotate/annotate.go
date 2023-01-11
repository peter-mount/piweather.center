package annotate

import (
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/graphics"
	"github.com/peter-mount/go-graphics/text"
	"image"
	"image/color"
)

func AnnotateTop(g graph.Graphics, top, bottom string, height int, fontSize float64) graph.Graphics {
	return graphics.New(Expand(g.Image(), height, 0, height, 0)).
		Background(image.Black).
		Foreground(color.White).
		FillRect(0, 0, g.Width(), height).
		FillRect(0, g.Bounds().Max.Y+height, g.Width(), g.Bounds().Max.Y+height+height).
		SetFont(text.Mono, fontSize).
		DrawText(image.Point{X: 0, Y: 0}, top).
		DrawText(image.Point{X: 0, Y: g.Bounds().Max.Y + height + height - int(fontSize)}, bottom)
}
