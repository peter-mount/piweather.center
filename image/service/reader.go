package service

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/image"
	exif2 "github.com/peter-mount/piweather.center/image/exif"
	"os"
	"time"
)

func (i *imageService) ReadWithExif(fileName string) (*image.Image, error) {
	log.Println("VisitReading", fileName)

	x, err := exif2.ReadExif(fileName)
	if err != nil {
		return nil, err
	}

	img := &image.Image{Filename: fileName, Exif: x, Time: x.Date}

	// No time then set it from the filesystem
	if img.Time.IsZero() {
		stat, err := os.Stat(fileName)
		if err != nil {
			return nil, err
		}
		img.Time = stat.ModTime()
	}

	return i.read(img)
}

func (i *imageService) ReadRaw(fileName string) (*image.Image, error) {
	log.Println("VisitReading", fileName)
	return i.read(&image.Image{Filename: fileName, Time: time.Now()})
}

func (i *imageService) read(img *image.Image) (*image.Image, error) {
	if err := img.Read(); err != nil {
		return nil, err
	}

	log.Printf("Image time: %v Bounds: %v", img.Time, img.Image.Bounds())

	return img, nil
}
