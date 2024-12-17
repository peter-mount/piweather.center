package skymap

import (
	"bufio"
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/astro/chart/render"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sidereal"
	io2 "github.com/peter-mount/piweather.center/util/io"
	"github.com/soniakeys/unit"
	"image"
	"image/color"
	"image/png"
	"io"
	"time"
)

func (s *Skymap) horizon() error {
	log.Printf("Generating horizon %q", *s.sphericalMap)

	w, h := 900, 900
	bounds := image.Rect(0, 0, w, h)

	// We need a CoordinateTransformer set to our location on the Earth and the greenwich siderial time
	jd := julian.FromTime(time.Now())
	transformer := coord.NewCoordinateTransformer(unit.AngleFromDeg(51.1), unit.AngleFromDeg(-0.5)).
		Sidereal(sidereal.FromJD(jd))

	// Next find the ra,dec of the position at the zenith
	x0, y0 := transformer.HzToEq(unit.AngleFromDeg(0.0), unit.AngleFromDeg(90.0))
	fmt.Printf("%v %v\n", x0.Hour(), y0.Deg())

	// Base Projection for the point at the zenith
	proj0 := chart.NewStereographicProjection(0, unit.AngleFromDeg(90), float64(w)/4.1, bounds)
	//proj0 := chart.NewPlainProjection(x0.Angle(), bounds)

	// Transformed Projection which maps Equatorial to Horizon coordinates
	proj := proj0.Transform(func(p chart.Point) chart.Point {
		A, h := transformer.EqToHz(p.X.RA(), p.Y)
		f := A.Deg() + 180.0
		for f < 0.0 {
			f = f + 360
		}
		for f >= 360 {
			f = f - 360
		}
		return chart.Point{X: unit.AngleFromDeg(f), Y: h}
	})

	dest := image.NewRGBA(bounds)

	gc := draw2dimg.NewGraphicContext(dest)

	layers := chart.NewLayers().
		// Used to draw with the correct origin
		SetProjection(proj).
		// We need to flip X axis when plotting horizon coordinates
		Flip(true, false)

	// Common values for drawing
	layers.SetFill(color.Black).
		SetStroke(color.White).
		SetLineWidth(1)

	layers.Add(chart.FloodFillLayer(proj0))

	layers.Add(s.Manager.FeatureLayer("milkyway", proj).SetFillStroke(color.Gray16{Y: 0x1111}))

	//layers.Add(chart.RaDecAxesLayer(proj).SetStroke(color.RGBA{G: 0x66, A: 0xff}))
	//layers.Add(chart.RaDecAxesLayer(proj).SetStroke(color.Gray16{Y: 0x3333}))

	//layers.Add(chart.RaDecAxesLayer(proj0).SetStroke(color.RGBA{R: 0x66, A: 0xff}))

	//layers.Add(s.Manager.FeatureLayer("const.border", proj).SetStroke(color.RGBA{B: 0xaa, A: 0xff}))
	layers.Add(s.Manager.FeatureLayer("const.line", proj).SetStroke(color.RGBA{B: 0xaa, A: 0xff}))

	layers.Add(s.catalog.NewLayer(render.BrightnessPixelStarRenderer, proj).
		MagLimit(6))

	// Solid green horizon
	//layers.Add(chart.HorizonLayer(proj0).SetFillStroke(colornames.Darkgreen))

	// Alternate dark green but Alpha allows for features below the horizon to still be visible.
	// Note: Darkgreen is G:0x64 but 0x32 looks better for this
	layers.Add(chart.HorizonLayer(proj0).SetFillStroke(color.RGBA{G: 0x32, A: 0x33}))
	//layers.Add(chart.HorizonLayer(proj0).SetFillStroke(color.RGBA{G: 0x32, A: 0xff}))

	//layers.Add(chart.BorderLayer(proj))

	layers.Draw(gc)

	return io2.NewWriter(func(w io.Writer) error {
		b := bufio.NewWriter(w)
		// Write the image into the buffer
		if err := png.Encode(b, dest); err != nil {
			return err
		}
		return b.Flush()
	}).CreateFile(*s.horizonMap)
}
