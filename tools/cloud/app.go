package cloud

import (
	"fmt"
	"github.com/peter-mount/go-build/version"
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/filter/cloud"
	"github.com/peter-mount/go-graphics/graphics"
	"github.com/peter-mount/go-graphics/text"
	"github.com/peter-mount/go-kernel/v2/log"
	image2 "github.com/peter-mount/piweather.center/image"
	"github.com/peter-mount/piweather.center/image/annotate"
	"github.com/peter-mount/piweather.center/image/service"
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
	Crop         *string         `kernel:"flag,crop,Crop image"`
}

func Join(baseDir, fileName string) string {
	if path.IsAbs(fileName) {
		return fileName
	}
	return path.Join(baseDir, fileName)
}

func (a *App) Start() error {
	log.Println("CloudDetect", version.Version)

	baseDir, fileName := a.Directory.Split(*a.ImageName)

	src, err := a.ImageService.ReadWithExif(Join(baseDir, *a.ImageName))
	if err != nil {
		return err
	}

	// Write exif data
	if err = src.WriteExif(); err != nil {
		return err
	}

	_, suffix := util.BaseName(fileName)

	if *a.Crop != "" {
		cp, err := annotate.ParseCoordinates(*a.Crop, 4)
		if err != nil {
			return err
		}
		if len(cp) != 4 {
			return fmt.Errorf("unsupported crop %q", *a.Crop)
		}
		src.Image = graphics.New(graph.Imutable(src.Image)).
			Crop(image.Rect(cp[0], cp[1], cp[2], cp[3])).
			Image()
	}

	if *a.SrcOutput != "" {
		//g := graphics.New(graph.DuplicateImage(src.Image))
		g := graphics.New(graph.Imutable(src.Image))
		g.Expand(100, 0, 100, 0).
			Background(image.Black).
			Foreground(color.White).
			FillRect(0, 0, g.Width(), 100).
			FillRect(0, g.Bounds().Max.Y+100, g.Width(), g.Bounds().Max.Y+200).
			SetFont(text.Mono, 90.0).
			DrawText(image.Point{X: 0, Y: 0}, fmt.Sprintf("%s", src.Time.Format(time.RFC1123)))

		srcOut := src.Duplicate(g.Image())
		srcOut.Filename = Join(baseDir, *a.SrcOutput+"."+suffix)
		if err = a.ImageService.Write(srcOut); err != nil {
			return err
		}
	}

	// Apply a mask. masked is the masked image or src if no mask applied
	masked := src
	var mask *image2.Image
	haveImageMask := *a.ImageMask != ""
	if haveImageMask {
		mask, err = a.ImageService.ReadRaw(Join(baseDir, *a.ImageMask))
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

	masked.Image = g.Image()

	masked.Filename = Join(baseDir, *a.CloudOutput+"."+suffix)
	if err = a.ImageService.Write(masked); err != nil {
		return err
	}

	if *a.Json != "" {
		log.Println("Writing", *a.Json)
		b, err := masked.JSON()
		if err != nil {
			return err
		}

		err = os.WriteFile(Join(baseDir, *a.Json), b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
