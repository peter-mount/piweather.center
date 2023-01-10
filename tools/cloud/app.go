package cloud

import (
	"github.com/peter-mount/piweather.center/image/service"
	"github.com/peter-mount/piweather.center/util"
	"io/ioutil"
	"path"
)

type App struct {
	ImageService service.Service `kernel:"inject"`
	ImageName    *string         `kernel:"flag,i,Image to process"`
	ImageMask    *string         `kernel:"flag,m,Image mask"`
	Debug        util.Debug      `kernel:"inject"`
	Dir          *string         `kernel:"flag,d,Directory to use,."`
	Json         *string         `kernel:"flag,json,Json format of image"`
}

func (a *App) Start() error {
	dir, fileName := path.Split(*a.ImageName)

	dir = path.Clean(dir)
	if dir == "." {
		dir = *a.Dir
	}

	src, err := a.ImageService.ReadWithExif(path.Join(dir, *a.ImageName))
	if err != nil {
		return err
	}

	// Write exif data
	if err = src.WriteExif(); err != nil {
		return err
	}

	_, suffix := util.BaseName(fileName)

	if *a.ImageMask != "" {
		mask, err := a.ImageService.ReadRaw(path.Join(*a.Dir, *a.ImageMask))
		if err != nil {
			return err
		}

		src = src.Overlay(mask)

		/*img := graph.MaskImage(src.Image, mask.Image)
		src = src.Duplicate(img)*/

		src.Filename = path.Join(*a.Dir, "masked."+suffix)
		if err = a.ImageService.Write(src); err != nil {
			return err
		}
	}

	if *a.Json != "" {
		a.Debug.Println("Writing", *a.Json)
		b, err := src.JSON()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(*a.Json, b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
