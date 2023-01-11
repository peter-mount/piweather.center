package videogenerator

import (
	"errors"
	piweather_center "github.com/peter-mount/piweather.center"
	"github.com/peter-mount/piweather.center/log"
)

type VideoGenerator struct {
	Config      *string `kernel:"flag,c,Config"`
	SrcFrame    *int    `kernel:"flag,sf,Source Frame Rate,24"`
	OutFrame    *int    `kernel:"flag,of,Output Frame Rate,24"`
	Output      *string `kernel:"flag,o,Output file"`
	DebugOutput *string `kernel:"flag,debug-output,Output each frame as png for debugging"`
	config      *Definition
}

func (v *VideoGenerator) PostInit() error {
	if *v.Config == "" {
		return errors.New("-c is required")
	}

	if *v.Output == "" {
		return errors.New("-o is required")
	}

	return nil
}

func (v *VideoGenerator) Start() error {
	log.Println("VideoGenerator", piweather_center.Version)

	v.config = &Definition{}
	if err := v.config.load(*v.Config); err != nil {
		return err
	}

	if err := v.config.findFrames(); err != nil {
		return err
	}

	if frames, err := v.config.collateFrames(); err != nil {
		return err
	} else {
		err := v.render(frames)
		if err != nil {
			return err
		}
	}

	return nil
}
