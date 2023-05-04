package videointro

import (
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/colornames"
	"math"
	"strings"
)

func (v *VideoIntro) DrawStats(frame int) {

	sec := float64(v.Config.Start) - (float64(frame) / float64(v.Config.Format.FrameRate))

	w, _ := float64(v.Config.Format.Width), float64(v.Config.Format.Height)
	cx, cy := float64(v.center.X), float64(v.center.Y)

	gc := v.gc

	gc.SetStrokeColor(colornames.Red)
	gc.SetFillColor(colornames.Red)
	gc.SetLineWidth(5)

	gc.SetFontData(draw2d.FontData{Name: "font/luxi", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontSize(80)

	secI, secF := math.Modf(sec)

	// Fit against 00.00 but still wobbles as it seems mono isn't mono
	l, t, r, b := stringSize(gc, "00.00")
	gc.FillStringAt(fmt.Sprintf("%02.0f.%02.0f", secI, secF*float64(v.Config.Format.FrameRate)),
		cx-(r-l)/2, cy+(b-t)/2)
	//drawString(gc, cx, cy, "%02.0f.%02.0f", secI, secF*float64(v.Config.Format.FrameRate))

	gc.SetStrokeColor(colornames.White)
	gc.SetFillColor(colornames.White)

	gc.SetFontSize(80)
	_, t, _, b = stringSize(gc, v.Config.Title)
	dy := b - t
	if v.Config.Subtitle != "" {
		gc.SetFontSize(40)
		_, t, _, b = stringSize(gc, v.Config.Subtitle)
		dy += b - t
	}

	x, y := w/2, 1830+(250-dy)/2
	gc.SetFontSize(80)
	y = drawString(gc, x, y, v.Config.Title)

	if v.Config.Subtitle != "" {
		gc.SetFontSize(40)
		y = drawString(gc, x, y, v.Config.Subtitle)
	}
}

func drawString(gc *draw2dimg.GraphicContext, x, y float64, s string, a ...interface{}) float64 {
	for _, str := range strings.Split(fmt.Sprintf(s, a...), "\n") {
		sl, st, sr, sb := gc.GetStringBounds(str)
		gc.FillStringAt(str, x-(sr-sl)/2, y+(sb-st)/2)
		y = y - st + sb
	}
	return y
}

func stringSize(gc *draw2dimg.GraphicContext, s string, a ...interface{}) (float64, float64, float64, float64) {
	var l, t, r, b float64
	for i, str := range strings.Split(fmt.Sprintf(s, a...), "\n") {
		sl, st, sr, sb := gc.GetStringBounds(str)
		if i == 0 {
			l, t, r, b = sl, st, sr, sb
		} else {
			l, t, r, b = fitString(l, t, r, b, sl, st, sr, sb)
		}
	}
	return l, t, r, b
}

func fitString(l, t, r, b, sl, st, sr, sb float64) (float64, float64, float64, float64) {
	return math.Min(l, sl), math.Min(t, st), math.Max(r, sr), math.Max(b, sb)
}
