package cloud

import (
	"github.com/peter-mount/go-kernel/util/task"
)

const (
	imageKey  = "rawImage"
	maskedKey = imageKey
	cloudKey  = imageKey
)

// Service handles the capture of an image
type Service struct {
	Config *CaptureConfig `kernel:"config,cloud"` // Cloud config
	Worker task.Queue     `kernel:"worker"`       // worker queue
}

type CaptureConfig struct {
	Schedule string `yaml:"schedule"` // Schedule to run the capture
	Dir      string `yaml:"dir"`      // Directory where images are stored
	Name     string `yaml:"name"`     // Name of captured image
	RawName  string `yaml:"rawName"`  // Raw name, usually Name with dng suffix
	Capture  string `yaml:"capture"`  // Command to use to capture the image
}
