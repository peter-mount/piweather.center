package videogenerator

import (
	"github.com/peter-mount/go-graphics/graphics"
	"github.com/peter-mount/go-graphics/util"
	"github.com/peter-mount/piweather.center/log"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Context struct {
	Start        time.Time
	End          time.Time
	LastTime     time.Time
	CurrentFrame int
	FrameCount   int
}

func (v *VideoGenerator) render(frames []*Frame) error {

	ctx := &Context{
		Start:        frames[0].Time,
		LastTime:     frames[0].Time,
		End:          frames[len(frames)-1].Time,
		CurrentFrame: 0,
		FrameCount:   len(frames),
	}

	log.Println("Rendering video.\n\n" +
		"+----------------------------------------------------+\n" +
		"| This will appear slow due to how much rendering is |\n" +
		"| required for each frame as it's passed to ffmpeg!  |\n" +
		"+----------------------------------------------------+")

	// Note: to ensure the frame rate is correct -framerate is before -i
	// to set the source framerate and then -r after to set the output framerate
	//
	// Discovered this with another animation where, when both were set to 30
	// the output was fixed to 25fps so with 30fps input a 60s animation ended
	// up playing for 72s (01:12) and ran slow.
	//
	// see: https://stackoverflow.com/a/66466918/1994472
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-framerate", strconv.Itoa(*v.SrcFrame),
		"-i", "-", // pipe from stdin
		"-r", strconv.Itoa(*v.OutFrame),
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		*v.Output,
	)

	pr, pw := io.Pipe()

	// We will pipe frames into ffmpeg
	cmd.Stdin = pr

	// ffmpeg output only if we are also showing output
	if log.IsVerbose() {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	var cmdErr error
	go func() {
		cmdErr = cmd.Run()
	}()

	bounds := image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: v.config.OutputSize.Width, Y: v.config.OutputSize.Height},
	}
	g := graphics.NewRect(bounds).
		Background(color.Black).
		Foreground(image.White).
		FillRectangle(bounds)

	for i, f := range frames {
		ctx.CurrentFrame = i

		if err := f.Render(ctx, g); err != nil {
			return err
		}

		if *v.DebugOutput != "" {
			if err := util.WritePNG(*v.DebugOutput, g.Image()); err != nil {
				return err
			}
		}

		if err := png.Encode(pw, g.Image()); err != nil {
			return err
		}

		ctx.LastTime = f.Time
	}
	if cmdErr != nil {
		return cmdErr
	}

	log.Println("Render completed")
	_ = pw.Close()

	return cmd.Wait()
}
