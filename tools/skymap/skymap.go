package skymap

import (
	"bufio"
	"flag"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	io2 "github.com/peter-mount/piweather.center/util/io"
	"image"
	"image/color"
	"image/png"
	"io"
)

type Skymap struct {
	overview *string `kernel:"flag,skymap-overview,Generate overview map"`
	catalog  *catalogue.Catalog
}

func (s *Skymap) Start() error {
	// Load the YBSC catalog
	s.catalog = &catalogue.Catalog{}
	if err := io2.NewReader(s.catalog.Read).
		Decompress().
		Open(application.FileName(application.STATIC, "bsc5.bin")); err != nil {
		return err
	}

	done := false

	if *s.overview != "" {
		done = true
		if err := s.renderOverview(); err != nil {
			return err
		}
	}

	if !done {
		flag.PrintDefaults()
	}
	return nil
}

func (s *Skymap) renderOverview() error {

	dest := image.NewRGBA(image.Rect(0, 0, 1080, 530))

	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(color.White)
	gc.SetStrokeColor(color.Black)
	gc.SetLineWidth(1)

	b := dest.Bounds()
	gc.BeginPath()
	gc.MoveTo(0, 0)
	gc.LineTo(float64(b.Dx()-1), 0)
	gc.LineTo(float64(b.Dx()-1), float64(b.Dy()-1))
	gc.LineTo(0, float64(b.Dy()-1))
	gc.Close()
	gc.FillStroke()

	_ = s.catalog.ForEach(func(e catalogue.Entry) error {
		x, y, m := int((180-e.RA().Deg())*3), int((90-e.Dec().Deg())*3), e.Mag()
		if x < 0 {
			x = x + b.Dx()
		}

		dest.Set(x, y, image.Black)
		if m < 3 {
			dest.Set(x, y-1, image.Black)
			dest.Set(x, y+1, image.Black)
			dest.Set(x-1, y, image.Black)
			dest.Set(x+1, y, image.Black)
		}
		if m < 1 {
			dest.Set(x-1, y-1, image.Black)
			dest.Set(x+1, y-1, image.Black)
			dest.Set(x-1, y+1, image.Black)
			dest.Set(x+1, y+1, image.Black)
		}
		if m < 0 {
			dest.Set(x, y-2, image.Black)
			dest.Set(x, y+2, image.Black)
			dest.Set(x-2, y, image.Black)
			dest.Set(x+2, y, image.Black)
		}
		return nil
	})

	return io2.NewWriter(func(w io.Writer) error {
		b := bufio.NewWriter(w)
		// Write the image into the buffer
		if err := png.Encode(b, dest); err != nil {
			return err
		}
		return b.Flush()
	}).CreateFile(*s.overview)
}
