package videogenerator

import (
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/util"
	xdraw "golang.org/x/image/draw"
	"image"
	"image/color"
	"image/draw"
	"math"
	"time"
)

type Frame struct {
	Time     time.Time       // time of frame
	Filename string          // Full path to frame
	Source   *Source         // Pointer to original source
	Next     *Frame          // Used to link frames with the same timestamp
	bounds   image.Rectangle // Bounds of image in frame
}

func (f *Frame) Render(ctx *Context, g graph.Graphics) error {
	img, err := util.ReadFile(f.Filename)
	if err != nil {
		return err
	}

	r := f.Source.Render
	f.bounds = r.Draw.Rect(img.Bounds())

	if r.Draw != nil {

		tmp := image.NewRGBA(f.bounds.Sub(f.bounds.Min))

		xdraw.NearestNeighbor.Scale(tmp, tmp.Bounds(), img, img.Bounds(), xdraw.Over, nil)

		g.DrawImage(f.bounds, tmp, image.Point{}, draw.Over)
	}

	if r.Keogram != nil {
		r.Keogram.render(f, ctx, g, img)
	}

	if f.Next != nil {
		return f.Next.Render(ctx, g)
	}

	return nil
}

func (k *Keogram) render(f *Frame, ctx *Context, g graph.Graphics, img image.Image) {
	// Default is vertical at middle of image
	ib := img.Bounds()
	if k.X == 0 && k.Y == 0 {
		k.X = ib.Dx() / 2
	}

	b := f.bounds.Add(image.Point{X: 0, Y: f.bounds.Dy() + 1})
	b.Max.Y = b.Min.Y + k.Height

	tickSize := int(math.Max(1, float64(k.Height)/32.0))

	dx := float64(b.Dx()) / float64(ctx.End.Sub(ctx.Start).Seconds())

	// x is for this frame, x0 is for the previous frame
	x := b.Min.X + int(f.Time.Sub(ctx.Start).Seconds()*dx)
	x0 := b.Min.X + int(ctx.LastTime.Sub(ctx.Start).Seconds()*dx)

	scale, start, end := k.setup(ib)
	for i := start; i <= end; i++ {
		col, y := k.getPixel(img, i, start, b.Max.Y, scale)
		g.Background(col).FillRect(x0, y, x, y)
	}

	g.Foreground(image.White).
		DrawRectangle(b)

	for h := ctx.Start.Truncate(time.Hour); h.Before(ctx.End); h = h.Add(time.Minute * 10) {
		x := b.Min.X + int(h.Sub(ctx.Start).Seconds()*dx)
		dy := tickSize
		switch {
		case h.Minute() == 0:
			dy = dy * 3
		case (h.Minute() % 15) == 0:
			dy = dy * 2
		}
		g.DrawLine(x, b.Max.Y, x, b.Max.Y-dy)
	}
}

// setup returns (scale,start,end)
func (k *Keogram) setup(b image.Rectangle) (float64, int, int) {
	max := b.Max.X
	if k.X == 0 {
		max = b.Max.Y
	}

	if k.End == 0 {
		k.End = max
	}
	if k.Start > k.End {
		k.Start, k.End = k.End, k.Start
	}

	return float64(k.Height) / float64(k.End-k.Start), k.Start, k.End
}

// getPixel returns color,y
func (k *Keogram) getPixel(img image.Image, i, start, maxY int, scale float64) (col color.Color, y int) {
	if k.X > 0 {
		col = img.At(k.X, i)
	} else {
		col = img.At(i, k.Y)
	}

	y = maxY - int(scale*float64(i-start))

	return
}
