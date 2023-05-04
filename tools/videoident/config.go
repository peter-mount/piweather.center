package videointro

import "time"

type Config struct {
	Title       string    `yaml:"title"`        // Top project title
	Subtitle    string    `yaml:"subtitle"`     // Top subtitle
	Description string    `yaml:"description"`  // Description
	Date        time.Time `yaml:"date"`         // Date of project
	Start       int       `yaml:"start"`        // Time to start from in seconds, usually a value between 5 and 45, default 5
	ClockSmooth bool      `yaml:"clock_smooth"` // true for smooth motion, false per second
	TestCard    bool      `yaml:"test_card"`    // true to show test card
	Format      Format    `yaml:"format"`       // Video format
}

type Format struct {
	Width     int `yaml:"width"`
	Height    int `yaml:"height"`
	FrameRate int `yaml:"frame_rate"` // Frame rate, defaults to 30
}

func (c *Config) Init() {
	if c.Start <= 0 {
		c.Start = 5
	}

	if c.Format.FrameRate <= 0 {
		c.Format.FrameRate = 30
	}

	if c.Format.Width == 0 || c.Format.Height == 0 {
		c.Format.Width = 3840
		c.Format.Height = 2160
	}
}
