package keogram

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/peter-mount/go-anim/graph"
	"github.com/peter-mount/go-anim/layout"
	"github.com/peter-mount/go-anim/renderer"
	"github.com/peter-mount/go-anim/util"
	"github.com/peter-mount/go-anim/util/frames"
	"image"
	"image/color"
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
	layout.BaseComponent                        // We are a Component for display within a Layout
	img                  draw.Image             // drawable image to contain the output
	axisMetric           *util.AlignmentMetrics // Used for drawing the time axis
	plotHeight           float64                // height of the plot
}

func New() *Keogram {
	k := &Keogram{}
	k.BaseComponent.SetPainter(k.paint)
	k.BaseComponent.SetType("Keogram")
	return k
}

func (k *Keogram) Sample(frame *frames.Frame, src image.Image) {
	srcBounds := src.Bounds()
	// Center of source image
	cx := srcBounds.Min.X + (srcBounds.Dx() >> 1)

	keoBounds := k.InsetBounds()
	if k.img == nil {
		// Use the size but keep origin at 0,0 otherwise it can be lost
		keoBounds = image.Rect(0, 0, keoBounds.Dx(), keoBounds.Dy())
		//fmt.Println("sam", keoBounds, k.plotHeight)
		k.img = image.NewRGBA(keoBounds)

		// Blank the image
		ctx := renderer.NewImageContext(k.img)
		gc := ctx.Gc()
		gc.SetFillColor(color.Black)
		gc.BeginPath()
		draw2dkit.Rectangle(gc, 0, 0, float64(keoBounds.Dx()), float64(keoBounds.Dy()))
		gc.Fill()

		gc.SetFillColor(color.White)

		// Add the initial date to the image
		gc.Save()
		_ = graph.SetFont(gc, "luxi 20 mono bold")
		r := image.Rect(0, 0, keoBounds.Dy(), keoBounds.Dx()-50)
		gc.Translate(float64(keoBounds.Dx()-35), float64(keoBounds.Max.Y))
		gc.Rotate(util.ToRad * -90.0)
		util.CenterAlignment.Fill(gc, r, 0, frame.Time.Format("2006-01-02 15:04"))
		gc.Restore()
	}

	b := k.img.Bounds()

	// Move graph 1 pixel left
	draw.Draw(k.img, b, k.img, image.Pt(1, 0), draw.Src)

	// Plot data from source to keogram
	x := b.Max.X - 2
	dy := k.plotHeight / float64(srcBounds.Dy())
	for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
		k.img.Set(x, int(float64(y)*dy), src.At(cx, y))
	}

	// Clear down the axis area
	for i := int(k.plotHeight); i < b.Max.Y; i++ {
		k.img.Set(x, i, color.Black)
	}

	// Every 30 minutes add the time to the x-axis
	if (frame.Time.Minute() % 30) == 0 {
		// But only if the time changes
		s := frame.Time.Format("15:04")
		if k.axisMetric.Lines[0] != s {
			k.axisMetric.Lines[0] = s
			ctx := renderer.NewImageContext(k.img)
			gc := ctx.Gc()
			gc.SetFillColor(color.White)
			_ = graph.SetFont(gc, "luxi 14 mono bold")
			gc.Translate(float64(keoBounds.Dx())-k.axisMetric.MaxLineHeight-1, float64(keoBounds.Max.Y))
			gc.Rotate(util.ToRad * -90.0)
			gc.Translate(25, 0)
			k.axisMetric.Fill(gc)
		}
	}
}

func (k *Keogram) Layout(ctx draw2d.GraphicContext) bool {
	if k.img == nil {
		// Work out the dimensions of the x-axis labels
		ctx1 := renderer.NewContext(200, 100)
		gc := ctx1.Gc()
		// First to get the dimensions
		_ = graph.SetFont(gc, "luxi 14 mono bold")
		m := util.RightAlignment.Metrics(gc, ctx1.Bounds().Rect(), 0, "00:00")
		mb := image.Rect(0, 0, int(m.MaxLineWidth), int(m.MaxLineHeight))
		// now to get the metric we'll use for rendering
		k.axisMetric = util.CenterAlignment.Metrics(gc, mb, 0, "00:00")

		// Now the final keogram image
		keoBounds := k.InsetBounds()
		l, t, r, b := k.GetInsets()
		k.plotHeight = float64(360 - t - b)
		keoBounds = image.Rect(keoBounds.Min.X, keoBounds.Min.Y,
			keoBounds.Dx()-l-r,
			int(k.plotHeight)+50+int(k.axisMetric.MaxLineWidth))
		k.SetBounds(keoBounds)
	}

	return k.BaseComponent.Layout(ctx)
}

func (k *Keogram) paint(gc *draw2dimg.GraphicContext) {
	if k.img == nil {
		return
	}

	// Draw captured data
	gc.DrawImage(k.img)
}
