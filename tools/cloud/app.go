package cloud

import (
	"fmt"
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/filter/cloud"
	"github.com/peter-mount/go-graphics/graphics"
	piweather_center "github.com/peter-mount/piweather.center"
	image2 "github.com/peter-mount/piweather.center/image"
	"github.com/peter-mount/piweather.center/image/annotate"
	"github.com/peter-mount/piweather.center/image/service"
	"github.com/peter-mount/piweather.center/log"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/dir"
	"image"
	"image/color"
	"image/draw"
	"os"
	"path"
	"time"
)

type App struct {
	ImageService service.Service `kernel:"inject"`
	ImageName    *string         `kernel:"flag,i,Image to process"`
	ImageMask    *string         `kernel:"flag,m,Image mask"`
	Json         *string         `kernel:"flag,json,Json format of image"`
	OverlaySrc   *bool           `kernel:"flag,overlay-source,Draw source instead of black on cloud image"`
	Directory    dir.Directory   `kernel:"inject"`
	SrcOutput    *string         `kernel:"flag,so,Optional output src annotated"`
	CloudOutput  *string         `kernel:"flag,co,Output file,cloud"`
}

func (a *App) Start() error {
	log.Println("Cloud Cover", piweather_center.Version)

	baseDir, fileName := a.Directory.Split(*a.ImageName)

	src, err := a.ImageService.ReadWithExif(path.Join(baseDir, *a.ImageName))
	if err != nil {
		return err
	}

	// Write exif data
	if err = src.WriteExif(); err != nil {
		return err
	}

	_, suffix := util.BaseName(fileName)

	if *a.SrcOutput != "" {
		g := graphics.New(graph.DuplicateImage(src.Image))
		srcOut := src.Duplicate(
			annotate.AnnotateTop(
				g,
				fmt.Sprintf("%s", src.Time.Format(time.RFC1123)),
				"",
				100,
				90.0,
			).
				Image())
		srcOut.Filename = path.Join(baseDir, *a.SrcOutput+"."+suffix)
		if err = a.ImageService.Write(srcOut); err != nil {
			return err
		}
	}

	// Apply a mask. masked is the masked image or src if no mask applied
	masked := src
	var mask *image2.Image
	haveImageMask := *a.ImageMask != ""
	if haveImageMask {
		mask, err = a.ImageService.ReadRaw(path.Join(baseDir, *a.ImageMask))
		if err != nil {
			return err
		}

		g := graphics.New(graph.DuplicateImage(src.Image)).
			DrawMask(src.Bounds(), mask.Image, image.Point{}, src.Image, image.Point{}, draw.Over)

		masked = src.Duplicate(g.Image())

		/*masked.Filename = path.Join(baseDir, "masked."+suffix)
		if err = a.ImageService.Write(masked); err != nil {
			return err
		}*/
	}

	cloudFilter := cloud.New()
	g := graphics.New(graph.DuplicateImage(masked.Image))
	g.Map(cloudFilter.Mapper)

	// If we have a mask then optionally draw what was masked out back in
	if haveImageMask && *a.OverlaySrc {
		_ = g.Filter(func(x, y int, col color.Color) (color.Color, error) {
			_, _, _, a := mask.Image.At(x, y).RGBA()
			if a >= 65535 {
				return src.Image.At(x, y), nil
			}
			return col, nil
		})
	}

	stats := cloudFilter.Stats()

	log.Printf("%s Cloud %.2f Sky %.2f Okta %v",
		src.Time.Format(time.RFC1123),
		stats.Cloud(),
		stats.Sky(),
		stats.OKTA(),
	)

	g = annotate.AnnotateTop(g,
		fmt.Sprintf("%s", src.Time.Format(time.RFC1123)),
		fmt.Sprintf("Cloud %.0f%% Sky %.0f%% Okta %d %v",
			stats.Cloud(),
			stats.Sky(),
			stats.OKTA(),
			stats.OKTA()),
		100,
		90.0,
	)

	/*
		g.Foreground(colornames.Red).DrawRectangle(g.Bounds())

		g = graphics.New(annotate.Expand(g.Image(), 100, 50, 50, 50))

		g.Foreground(colornames.Green).DrawRectangle(g.Bounds())

		g.Background(image.Black).
			Foreground(color.White).
			FillRect(0, 0, g.Width(), 100).
			FillRect(0, g.Bounds().Max.Y-100, g.Width(), masked.Bounds().Max.Y).
			SetFont(text.Mono, 90).
			DrawTextf(image.Point{X: 0, Y: 0}, "%s", src.Time.Format(time.RFC1123)).
			DrawTextf(image.Point{X: 0, Y: g.Bounds().Max.Y - 90},
				"Cloud %.0f%% Sky %.0f%% Okta %d %v",
				stats.Cloud(),
				stats.Sky(),
				stats.OKTA(),
				stats.OKTA(),
			)
	*/

	masked.Image = g.Image()

	masked.Filename = path.Join(baseDir, *a.CloudOutput+"."+suffix)
	if err = a.ImageService.Write(masked); err != nil {
		return err
	}

	if *a.Json != "" {
		log.Println("Writing", *a.Json)
		b, err := masked.JSON()
		if err != nil {
			return err
		}

		err = os.WriteFile(path.Join(baseDir, *a.Json), b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
