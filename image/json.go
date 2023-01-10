package image

import (
	"bytes"
	"encoding/json"
	"github.com/peter-mount/piweather.center/image/exif"
	"image/jpeg"
	"image/png"
	"time"
)

type jsonImage struct {
	Filename string     `json:"filename"`
	Time     time.Time  `json:"time"`
	Meta     *exif.Exif `json:"meta,omitempty"`
	Image    []byte     `json:"image"`
}

func (i *Image) JSON() ([]byte, error) {
	d := jsonImage{
		Filename: i.Filename,
		Time:     i.Time,
		Meta:     i.Exif,
		Image:    nil,
	}

	// Encode image
	b := &bytes.Buffer{}
	switch i.Type() {
	case "jpg":
		if err := jpeg.Encode(b, i.Image, &jpeg.Options{Quality: 90}); err != nil {
			return nil, err
		}

	case "png":
		if err := png.Encode(b, i.Image); err != nil {
			return nil, err
		}

	default:
		return nil, unsupportedType
	}
	d.Image = b.Bytes()

	return json.Marshal(d)
}

func (i *Image) MarshalJSON() ([]byte, error) {
	if i == nil {
		return []byte("null"), nil
	}

	b := &bytes.Buffer{}
	b.WriteByte('{')

	b.WriteByte('}')
	return b.Bytes(), nil
}
