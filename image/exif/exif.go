package exif

import (
	"encoding/json"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Exif struct {
	Date       time.Time              `json:"date"`
	Filename   string                 `json:"filename"`
	Properties map[string]interface{} `json:"props"`
}

func ReadExif(fileName string) (*Exif, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decode exif
	x, err := exif.Decode(f)
	if err != nil {
		return nil, err
	}

	// extract image time
	t, err := getExifTime(x)
	if err != nil {
		return nil, err
	}

	e := &Exif{Filename: fileName, Date: t, Properties: make(map[string]interface{})}

	// Extract properties
	err = x.Walk(e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Exif) Walk(name exif.FieldName, tag *tiff.Tag) error {
	v := DecodeTiffTag(tag)
	if v != nil {
		e.Properties[string(name)] = v
	}
	return nil
}

// DecodeTiffTag will attempt to decode a tiff.Tag and return its value in native format.
// Untested: tag values which are arrays
func DecodeTiffTag(tag *tiff.Tag) interface{} {
	s := tag.String()

	if t, err := time.Parse(dateFormat, s); err == nil {
		return t
	}

	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return strings.Trim(s, "\"\"")
	}

	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	return nil
}

func (e *Exif) WriteFile(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return e.Write(f)
}

func (e *Exif) Write(w io.Writer) error {
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	n, err := w.Write(b)
	if err != nil {
		return err
	}
	if n != len(b) {
		return fmt.Errorf("wrote %d expected %d", n, len(b))
	}

	return nil
}

const dateFormat = "\"2006:01:02 15:04:05\""

var dateTimeNames = []exif.FieldName{exif.DateTime, exif.DateTimeOriginal, exif.DateTimeDigitized}

func getExifTime(x *exif.Exif) (time.Time, error) {
	for _, n := range dateTimeNames {
		dt, err := x.Get(n)
		if err != nil {
			if !exif.IsTagNotPresentError(err) {
				return time.Time{}, err
			}
		} else {
			t, err := time.Parse(dateFormat, dt.String())
			if err != nil {
				return time.Time{}, err
			}
			return t, nil
		}
	}
	return time.Now(), nil
}
