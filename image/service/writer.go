package service

import (
	"errors"
	"github.com/peter-mount/piweather.center/image"
	"github.com/peter-mount/piweather.center/log"
	"os"
)

func (i *imageService) Write(img *image.Image) error {
	if img.Filename == "" {
		return errors.New("no filename defined")
	}
	log.Println("Writing", img.Filename)
	err := img.Write()
	if err == nil {
		// If Image has a time then set the file time to it.
		if !img.Time.IsZero() {
			err = os.Chtimes(img.Filename, img.Time, img.Time)
		}
	}
	return err
}
