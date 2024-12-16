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

func (s *Skymap) renderOverview() error {
	log.Printf("Generating overview %q", *s.overview)

	w, h := 1080, 530
	bounds := image.Rect(0, 0, w, h)

	proj := chart.NewPlainProjection(
		unit.RAFromHour(0.0).Angle(),
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

	mw, err := s.Manager.Feature("milkyway")
	if err != nil {
		return err
	}
	layers.Add(mw.GetLayerAll(proj).
		SetFill(color.Gray16{Y: 0x1111}).
		SetStroke(color.Gray16{Y: 0x1111}))

	layers.Add(chart.RaDecAxesLayer(proj).SetStroke(color.Gray16{Y: 0x3333}))

	layers.Add(catalogue.NewCatalogLayer(s.catalog, render.BrightnessPixelStarRenderer, proj))

	layers.Draw(gc)

	return io2.NewWriter(func(w io.Writer) error {
		b := bufio.NewWriter(w)
		// Write the image into the buffer
		if err := png.Encode(b, dest); err != nil {
			return err
		}
		return b.Flush()
	}).CreateFile(*s.overview)
}
