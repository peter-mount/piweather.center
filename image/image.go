package image

import (
	"github.com/peter-mount/go-graphics/util"
	"github.com/peter-mount/piweather.center/image/exif"
	util2 "github.com/peter-mount/piweather.center/util"
	"image"
	"path"
	"strings"
	"time"
)

// Image is a representation of an Image being processed
type Image struct {
	Filename string      // File name
	Time     time.Time   // Time image taken
	Image    image.Image // Source image
	Exif     *exif.Exif  // Exif data
}

// Bounds of the underlying Image
func (i *Image) Bounds() image.Rectangle { return i.Image.Bounds() }

// Read the image data from disk
func (i *Image) Read() error {
	img, err := util.ReadFile(i.Filename)
	if err != nil {
		return err
	}

	i.Image = img
	return nil
}

// Write the image data to disk along with the exif meta data.
// This is the same as calling WriteImage() then WriteExif()
func (i *Image) Write() error {
	if i.Filename == "" {
		return noFilename
	}

	if err := i.WriteImage(); err != nil {
		return err
	}

	return i.WriteExif()
}

// WriteImage writes just the image data to disk.
func (i *Image) WriteImage() error {
	if i.Filename == "" {
		return noFilename
	}

	switch i.Type() {
	case "jpg":
		return util.WriteJPG(i.Filename, i.Image)

	case "png":
		return util.WritePNG(i.Filename, i.Image)

	default:
		return unsupportedType
	}
}

// WriteExif writes the exif data as json to disk.
func (i *Image) WriteExif() error {
	if i.Exif != nil {
		if i.Filename == "" {
			return noFilename
		}

		dir, fileName := path.Split(i.Filename)
		dir = path.Clean(dir)
		prefix, _ := util2.BaseName(fileName)
		return i.Exif.WriteFile(path.Join(dir, prefix+".json"))
	}

	return nil
}

func (i *Image) Type() string {
	switch {
	case strings.HasSuffix(i.Filename, ".jpg"), strings.HasSuffix(i.Filename, ".jpeg"):
		return "jpg"

	case strings.HasSuffix(i.Filename, ".png"):
		return "png"

	default:
		return ""
	}
}

// Duplicate returns a new Image with the same metadata but with a new underlying Image
func (i *Image) Duplicate(img image.Image) *Image {
	return &Image{Time: i.Time, Image: img, Exif: i.Exif}
}
