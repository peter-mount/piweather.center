package keogram

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/peter-mount/go-anim/layout"
	"golang.org/x/image/colornames"
	"image"
	"image/draw"
)

// Keogram captures a vertical slice of the sky and records it as a timeline
// of the sky's condition.
//
// Normally they are used for displaying the intensity of auroral display
// (keo in keogram is from "keoeeit" which os Inuit for Aurora Borealis)
// But we also use it here for showing clouds over time.
//
// see also: https://en.wikipedia.org/wiki/Keogram
type Keogram struct {
	layout.BaseComponent            // We are a Component for display within a Layout
	img                  draw.Image // drawable image to contain the output
}

func New() *Keogram {
	k := &Keogram{}
	k.BaseComponent.SetPainter(k.paint)
	k.BaseComponent.SetType("Keogram")
	return k
}

func (k *Keogram) Sample(src image.Image) {
	srcBounds := src.Bounds()
	// Center of source image
	cx := srcBounds.Min.X + (srcBounds.Dx() >> 1)

	keoBounds := k.Bounds()
	if k.img == nil {
		k.img = image.NewRGBA(keoBounds)
	}

	draw.Draw(k.img, keoBounds, k.img, image.Pt(1, 0), draw.Src)

	dy := float64(keoBounds.Dy()) / float64(srcBounds.Dy())
	for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
		k.img.Set(keoBounds.Max.X-2, int(float64(y)*dy), src.At(cx, y))
	}
}

func (k *Keogram) Layout(ctx draw2d.GraphicContext) bool {
	keoBounds := k.Bounds()
	if keoBounds.Dy() < 180 {
		keoBounds.Max.Y = keoBounds.Min.Y + 360
		k.SetBounds(keoBounds)
	}
	//log.Printf("%q Layout %v", k.GetType(), keoBounds)
	return k.BaseComponent.Layout(ctx)
}

func (k *Keogram) paint(gc *draw2dimg.GraphicContext) {
	keoBounds := k.Bounds()
	//log.Printf("%q paint %v", k.GetType(), keoBounds)

	// Border around the results
	p1 := keoBounds.Max
	gc.SetFillColor(colornames.Black)
	gc.SetStrokeColor(colornames.White)
	gc.BeginPath()
	gc.MoveTo(0, 0)
	gc.LineTo(float64(p1.X), 0)
	gc.LineTo(float64(p1.X), float64(p1.Y))
	gc.LineTo(0, float64(p1.Y))
	gc.Close()

	// Fill the component background
	gc.Save()
	gc.Fill()
	gc.Restore()

	// Draw captured data
	if k.img != nil {
		gc.DrawImage(k.img)
	}

	// Border around the component
	gc.Stroke()
}
