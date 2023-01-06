package cloud

import (
	"context"
	"log"
	"os/exec"
)

// captureImage captures a new image if enabled
func (c *Service) captureImage(_ context.Context) error {
	if c.Config.Capture != "" {
		log.Print("Capturing image")
		cmd := exec.Command("libcamera-still", "-r", "-o", c.Config.Name)
		cmd.Dir = c.Config.Dir
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
