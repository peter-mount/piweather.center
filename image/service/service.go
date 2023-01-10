package service

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/image"
)

func init() {
	kernel.RegisterAPI((*Service)(nil), &imageService{})
}

type Service interface {
	ReadWithExif(fileName string) (*image.Image, error)
	ReadRaw(fileName string) (*image.Image, error)
	Write(img *image.Image) error
}

type imageService struct{}

func (i *imageService) Start() error {
	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	//	exif.RegisterParsers(mknote.All...)

	return nil
}
