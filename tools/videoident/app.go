package videointro

import (
	"errors"
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-kernel/v2/log"
	common "github.com/peter-mount/piweather.center"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type VideoIntro struct {
	Config       *Config `kernel:"config,intro"` // Our config yaml file
	Frame        *int    `kernel:"flag,frame,Render single frame,-1"`
	Output       *string `kernel:"flag,o,Write to file"`
	Format       *string `kernel:"flag,fmt,Image format,png"`
	img          graph.Image
	gc           *draw2dimg.GraphicContext
	bounds       image.Rectangle
	center       image.Point
	ffmpeg       *exec.Cmd
	ffmpegReader *io.PipeReader
	ffmpegWriter *io.PipeWriter
}

func (v *VideoIntro) PostInit() error {
	if *v.Output == "" {
		return errors.New("-o is required")
	}

	return nil
}

func (v *VideoIntro) Start() error {
	log.Println("videointro", common.Version)

	// Ensure defaults are applied
	v.Config.Init()

	v.bounds = image.Rect(0, 0, v.Config.Format.Width, v.Config.Format.Height)
	v.center = image.Point{
		X: v.Config.Format.Width >> 1,
		Y: v.Config.Format.Height >> 1,
	}

	var cmdErr error
	if strings.HasSuffix(*v.Output, ".mp4") {
		v.ffmpeg = exec.Command(
			"ffmpeg",
			"-y",
			"-framerate", strconv.Itoa(v.Config.Format.FrameRate),
			"-i", "-", // pipe from stdin
			"-r", strconv.Itoa(v.Config.Format.FrameRate),
			"-c:v", "libx264",
			"-pix_fmt", "yuv420p",
			*v.Output,
		)

		v.ffmpegReader, v.ffmpegWriter = io.Pipe()

		v.ffmpeg.Stdin = v.ffmpegReader

		if log.IsVerbose() {
			v.ffmpeg.Stdout = os.Stdout
			v.ffmpeg.Stderr = os.Stderr
		}

		go func() {
			cmdErr = v.ffmpeg.Run()
		}()
	}

	v.img = image.NewRGBA(v.bounds)
	v.gc = draw2dimg.NewGraphicContext(v.img)

	log.Printf("Frame size %dx%d rate %d fps",
		v.Config.Format.Width,
		v.Config.Format.Height,
		v.Config.Format.FrameRate,
	)

	if *v.Frame >= 0 {
		if err := v.RenderFrames(*v.Frame, *v.Frame); err != nil {
			return err
		}
	} else if err := v.RenderFrames(0, v.Config.Start*v.Config.Format.FrameRate); err != nil {
		log.Println(err)
		return err
	}

	if cmdErr != nil {
		return cmdErr
	}

	if v.ffmpeg != nil {

		log.Println("Render complete")
		_ = v.ffmpegWriter.Close()
		//return v.ffmpeg.Wait()
	}

	return nil
}

func (v *VideoIntro) RenderFrames(start, end int) error {
	if start > end {
		return errors.New("start frame after end")
	}

	if start != end {
		log.Printf("Rendering frames %d...%d", start, end)
	}

	for frame := start; frame <= end; frame++ {
		log.Printf("Rendering frame %d/%d", frame, end)
		if err := v.RenderFrame(frame); err != nil {
			return err
		}
	}
	return nil
}

func (v *VideoIntro) RenderFrame(frame int) error {
	// Erase image
	v.gc.SetFillColor(image.Black)
	v.gc.ClearRect(0, 0, v.bounds.Dx(), v.bounds.Dy())

	if v.Config.TestCard {
		v.TestCard(frame)
	}

	v.DrawFrame(frame)
	v.DrawStats(frame)
	return v.WriteFrame(frame)
}

func (v *VideoIntro) WriteFrame(frame int) error {
	if v.ffmpegWriter != nil {
		return v.writeFrame(v.ffmpegWriter, "png")
	}

	if *v.Output == "-" {
		return v.writeFrame(os.Stdout, "png")
	}

	fileName := *v.Output
	if matches, _ := regexp.Match("%([0-9]+)d", []byte(fileName)); matches {
		fileName = fmt.Sprintf(fileName, frame)
	}

	log.Printf("Writing %s", fileName)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return v.writeFrame(f, *v.Format)
}

func (v *VideoIntro) writeFrame(w io.Writer, f string) error {
	switch f {
	case "png":
		return png.Encode(w, v.img)
	case "jpg":
		return jpeg.Encode(w, v.img, &jpeg.Options{Quality: 90})
	default:
		return fmt.Errorf("unsupported image type %q", f)
	}
}
