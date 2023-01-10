package cloud

import (
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/filter/cloud"
	"github.com/peter-mount/go-graphics/graphics"
	"github.com/peter-mount/go-graphics/text"
	piweather_center "github.com/peter-mount/piweather.center"
	"github.com/peter-mount/piweather.center/image/service"
	"github.com/peter-mount/piweather.center/log"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/dir"
	"image"
	"image/color"
	"io/ioutil"
	"path"
)

type App struct {
	ImageService service.Service `kernel:"inject"`
	ImageName    *string         `kernel:"flag,i,Image to process"`
	ImageMask    *string         `kernel:"flag,m,Image mask"`
	Json         *string         `kernel:"flag,json,Json format of image"`
	Directory    dir.Directory   `kernel:"inject"`
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

	if *a.ImageMask != "" {
		mask, err := a.ImageService.ReadRaw(path.Join(baseDir, *a.ImageMask))
		if err != nil {
			return err
		}

		src = src.Overlay(mask)

		/*img := graph.MaskImage(src.Image, mask.Image)
		src = src.Duplicate(img)*/

		src.Filename = path.Join(baseDir, "masked."+suffix)
		if err = a.ImageService.Write(src); err != nil {
			return err
		}
	}

	cloudFilter := cloud.New()
	g := graphics.New(graph.DuplicateImage(src.Image))
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
		)
	src.Image = g.Image()

	src.Filename = path.Join(baseDir, "cloud."+suffix)
	if err = a.ImageService.Write(src); err != nil {
		return err
	}

	if *a.Json != "" {
		log.Println("Writing", *a.Json)
		b, err := src.JSON()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path.Join(baseDir, *a.Json), b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
