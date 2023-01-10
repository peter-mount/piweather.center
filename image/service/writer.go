package service

import (
	"errors"
	"github.com/peter-mount/piweather.center/image"
)

func (i *imageService) Write(img *image.Image) error {
	if img.Filename == "" {
		return errors.New("no filename defined")
	}
	i.Debug.Println("Writing", img.Filename)
	return img.Write()
}
