package cloud

import (
	"github.com/peter-mount/go-graphics/util"
	"image"
	"path"
)

func (c *CaptureConfig) ReadImage(name string) (image.Image, error) {
	return util.ReadFile(path.Join(c.Dir, name))
}

func (c *CaptureConfig) WriteImage(name string, img image.Image) error {
	return util.WritePNG(path.Join(c.Dir, name), img)
}
