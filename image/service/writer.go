package service

import (
	"errors"
	"github.com/peter-mount/piweather.center/image"
	"github.com/peter-mount/piweather.center/log"
)

func (i *imageService) Write(img *image.Image) error {
	if img.Filename == "" {
		return errors.New("no filename defined")
	}
	log.Println("Writing", img.Filename)
	return img.Write()
}
