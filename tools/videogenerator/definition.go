package videogenerator

import (
	"errors"
	"fmt"
	"github.com/peter-mount/piweather.center/log"
	"gopkg.in/yaml.v2"
	"os"
	"sort"
	"time"
)

type Definition struct {
	OutputSize Size      `yaml:"outputSize"` // Size of generated output
	Sources    []*Source `yaml:"sources"`    // List of sources to use in output
}

type Size struct {
	Width  int `yaml:"width"`  // Width of output
	Height int `yaml:"height"` // Height of output
}

func (d *Definition) load(fileName string) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, d)
	if err != nil {
		return err
	}

	for _, s := range d.Sources {
		err := s.init()
		if err != nil {
			return err
		}
	}

	// We require at least 1 source
	if len(d.Sources) == 0 {
		return errors.New("no sources defined")
	}

	// Output size needs setting. Limits here are arbitrary but anything smaller than this is obviously not viable
	if d.OutputSize.Width < 320 || d.OutputSize.Height < 200 {
		return fmt.Errorf("invalid output frame size %d,%d", d.OutputSize.Width, d.OutputSize.Height)
	}

	return nil
}

func (d *Definition) findFrames() error {
	for _, s := range d.Sources {
		if err := s.findFrames(); err != nil {
			return err
		}
	}
	return nil
}

func (d *Definition) collateFrames() ([]*Frame, error) {
	log.Println("Collating frames")

	frames := make(map[int64]*Frame)
	for _, s := range d.Sources {
		for _, f := range s.Frames {
			k := f.Time.Truncate(time.Second * 30).Unix()
			if e, exists := frames[k]; exists {
				f.Next = e
			}
			frames[k] = f
		}
	}

	log.Printf("Collated %d frames", len(frames))

	var f []*Frame
	for _, v := range frames {
		f = append(f, v)
	}

	sort.SliceStable(f, func(i, j int) bool {
		return f[i].Time.Before(f[j].Time)
	})

	return f, nil
}
