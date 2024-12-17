package skymap

import (
	"bufio"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/astro/chart/render"
	io2 "github.com/peter-mount/piweather.center/util/io"
	"github.com/soniakeys/unit"
	"image"
	"image/color"
	"image/png"
	"io"
)

func (s *Skymap) spherical() error {
	log.Printf("Generating spherical %q", *s.sphericalMap)

	w, h := 900, 900
	bounds := image.Rect(0, 0, w, h)

	proj := chart.NewStereographicProjection(
		unit.RAFromHour(5.0).Angle(),
		unit.AngleFromDeg(0.0),
		//unit.AngleFromDeg(90.0),
		float64(w)/2.0,
		bounds,
	)

	dest := image.NewRGBA(bounds)

	gc := draw2dimg.NewGraphicContext(dest)

	layers := chart.NewLayers()

	// Common values for drawing
	layers.SetFill(color.Black).
		SetStroke(color.White).
		SetLineWidth(1)

	layers.Add(chart.FloodFillLayer(proj))

	layers.Add(s.Manager.FeatureLayer("milkyway", proj).
		SetFill(color.Gray16{Y: 0x1111}).
		SetStroke(color.Gray16{Y: 0x1111}))

	layers.Add(chart.RaDecAxesLayer(proj).SetStroke(color.Gray16{Y: 0x3333}))

	layers.Add(s.Manager.FeatureLayer("const.border", proj).SetStroke(color.Gray16{Y: 0x4444}))

	layers.Add(s.Manager.FeatureLayer("const.line", proj).SetStroke(color.Gray16{Y: 0x5555}))

	layers.Add(catalogue.NewCatalogLayer(s.catalog, render.BrightnessPixelStarRenderer, proj))

	gc.Save()
	gc.Translate(proj.GetCenter())
	// We need to flip both axes when plotting equatorial coordinates
	gc.Scale(-1, -1)
	layers.Draw(gc)
	gc.Restore()

	return io2.NewWriter(func(w io.Writer) error {
		b := bufio.NewWriter(w)
		// Write the image into the buffer
		if err := png.Encode(b, dest); err != nil {
			return err
		}
		return b.Flush()
	}).CreateFile(*s.sphericalMap)
}