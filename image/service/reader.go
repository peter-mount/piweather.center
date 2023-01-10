package service

import (
	"github.com/peter-mount/piweather.center/image"
	exif2 "github.com/peter-mount/piweather.center/image/exif"
	"time"
)

func (i *imageService) ReadWithExif(fileName string) (*image.Image, error) {
	i.Debug.Println("Reading", fileName)

	x, err := exif2.ReadExif(fileName)
	if err != nil {
		return nil, err
	}

	return i.read(&image.Image{Filename: fileName, Exif: x, Time: x.Date})
}

func (i *imageService) ReadRaw(fileName string) (*image.Image, error) {
	i.Debug.Println("Reading", fileName)
	return i.read(&image.Image{Filename: fileName, Time: time.Now()})
}

func (i *imageService) read(img *image.Image) (*image.Image, error) {
	if err := img.Read(); err != nil {
		return nil, err
	}

	i.Debug.Printf("Image time: %v Bounds: %v", img.Time, img.Image.Bounds())

	return img, nil
}
