package cloud

import (
	"github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/filter"
	"github.com/peter-mount/go-graphics/filter/cloud"
	"github.com/peter-mount/go-graphics/graphics"
	"github.com/peter-mount/go-graphics/histogram"
	"github.com/peter-mount/go-graphics/text"
	"image"
	"image/color"
	"log"
	"path"
)

type Stats struct {
	Total  int // total pixels with data
	Cloud  int // Cloud pixels
	Sky    int // Sky pixels
	Sun    int // Sun pixels
	NoData int // NoData pixels
}

func (c Service) Start() error {
	img, err := c.Config.ReadImage(c.Config.Name)
	if err != nil {
		return err
	}
	/*
	   mask, err := c.Config.ReadImage("mask.png")
	   if err != nil {
	     return err
	   }
	*/
	histBefore := histogram.New()
	histAfter := histogram.New()

	cloudFilter := cloud.New()

	// We need a mutable one
	g := graphics.New(graph.DuplicateImage(img))

	err = g.Filter(histBefore.AnalyzeFilter)
	if err != nil {
		return err
	}
	histBefore.ResetValuesBelow(int(cloudFilter.DarkLim)).
		ResetValuesAbove(int(cloudFilter.LightLim))

	deltaRGB := filter.DeltaRGBFromHistogram(histBefore)
	log.Println(deltaRGB)
	deltaRGB.R = -9000

	err = g.Filter(graph.Of(
		deltaRGB.Filter,
		histAfter.AnalyzeFilter))
	if err != nil {
		return err
	}

	histAfter.ResetValuesBelow(int(cloudFilter.DarkLim)).
		ResetValuesAbove(int(cloudFilter.LightLim))
	/*
	   err = g.WritePNG(path.Join(c.Config.Dir, "adjusted.png"))
	   if err != nil {
	     return err
	   }
	*/
	/*
	   g.DrawImage(mask.Bounds(), mask, image.Point{}, draw.Over)
	   err = g.WritePNG(path.Join(c.Config.Dir, "masked-sky.png"))
	   if err != nil {
	     return err
	   }
	*/
	g.Map(cloudFilter.Mapper)

	stats := cloudFilter.Stats()
	log.Printf("Cloud %.2f Sky %.2f Okta %v\n",
		stats.Cloud(),
		stats.Sky(),
		stats.OKTA(),
	)

	g.Background(image.Black).FillRect(0, 0, g.Width(), 100)

	g.Background(image.Black).
		Foreground(color.White).
		SetFont(text.Mono, 90).
		DrawTextf(image.Point{X: 400, Y: 0},
			"Cloud %.0f%% Sky %.0f%% Okta %d %v\n",
			stats.Cloud(),
			stats.Sky(),
			stats.OKTA(),
			stats.OKTA(),
		) /*.
		  Draw(histBefore.Drawable(image.Point{X: 10})).
		  Draw(histAfter.Drawable(image.Point{X: 10, Y: 280}))*/

	err = g.WritePNG(path.Join(c.Config.Dir, "cloud.png"))
	if err != nil {
		return err
	}

	return nil
}

/*
func (w *Work) ApplyMask() error {
  // Load image mask
  mask, err := w.Config.ReadImage("mask.png")
  if err != nil {
    return err
  }

  log.Printf("Creating mask %v", w.SrcImage.Bounds())

  w.MaskedImage = graph.MaskImage(w.SrcImage, mask)

  // We now have the image with the mask applied
  // so queue the write of that masked image
  err = w.Config.WriteImage("masked-sky.png", w.MaskedImage)
  if err != nil {
    return err
  }

  // Now calculate the cloud cover
  return nil
}
*/
