package cloud

import (
	piweather_center "github.com/peter-mount/piweather.center"
	"github.com/peter-mount/piweather.center/image/service"
	"github.com/peter-mount/piweather.center/log"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/dir"
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
