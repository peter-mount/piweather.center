package videogenerator

import (
	"github.com/peter-mount/go-kernel/v2/util/walk"
	exif2 "github.com/peter-mount/piweather.center/image/exif"
	"github.com/peter-mount/piweather.center/log"
	"io"
	"os"
	"regexp"
)

type Source struct {
	Name      string   `yaml:"name"`      // Name of source, optional, used in some places
	Directory string   `yaml:"directory"` // Directory containing frames
	Pattern   string   `yaml:"pattern"`   // Filename pattern
	Render    Render   `yaml:"render"`    // Render instructions
	Frames    []*Frame `yaml:"-"`         // Discovered frames
	regexp    *regexp.Regexp
}

type Render struct {
	Draw    *Bounds  `yaml:"draw,omitempty"`    // Render at location in frame
	Keogram *Keogram `yaml:"keogram,omitempty"` // Keogram
}

type Keogram struct {
	X      int `yaml:"x,omitempty"` // Position of Keogram in image.
	Y      int `yaml:"y,omitempty"` // X is normal, Y optional, only 1 is valid
	Height int `yaml:"height"`      // Height of Keogram in pixels. Width is that of image
	Start  int `yaml:"start"`       // Start offset, default 0
	End    int `yaml:"end"`         // End offset, default height (or width when Y mode used)
}

func (s *Source) init() error {
	if s.Directory == "" {
		s.Directory = "."
	}

	reg, err := regexp.Compile(s.Pattern)
	if err != nil {
		return err
	}
	s.regexp = reg
	return nil
}

func (s *Source) matches(_ string, info os.FileInfo) bool {
	return s.regexp.MatchString(info.Name())
}

func (s *Source) append(p string, info os.FileInfo) error {
	f := &Frame{
		Time:     info.ModTime(),
		Filename: p,
		Source:   s,
	}

	x, err := exif2.ReadExif(p)
	if err == nil && x.Date.Before(f.Time) {
		f.Time = x.Date
	}

	if err == nil || err == io.EOF {
		s.Frames = append(s.Frames, f)
		return nil
	}

	return err
}

func (s *Source) findFrames() error {
	log.Println("Scanning", s.Directory)
	err := walk.NewPathWalker().
		Then(s.append).
		If(s.matches).
		IsFile().
		Walk(s.Directory)
	if err == nil {
		log.Printf("Found %d frames", len(s.Frames))
	}
	return err
}
