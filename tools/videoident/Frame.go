package videointro

import (
	"github.com/peter-mount/go-graphics/util"
	"github.com/peter-mount/go-kernel/v2/log"
	"golang.org/x/image/colornames"
	"image"
	"math"
)

func (v *VideoIntro) DrawFrame(frame int) {
	gc := v.gc

	clockThickness := 5.0

	// Radius of central clock
	radius := float64(util.Min(v.Config.Format.Width, v.Config.Format.Height)) / 4
	outerRadius := radius + (7 * clockThickness) + 15
	radius = radius - 60
	radiusInner := radius / 2

	cx, cy := float64(v.center.X), float64(v.center.Y)

	// Clock dial
	gc.SetFillColor(image.Black)
	gc.BeginPath()
	gc.ArcTo(cx, cy, outerRadius, outerRadius, 0, 2*math.Pi)
	gc.Fill()

	gc.SetLineWidth(clockThickness * 2)
	gc.SetStrokeColor(image.White)
	gc.SetFillColor(image.White)
	gc.BeginPath()
	gc.ArcTo(cx, cy, outerRadius, outerRadius, 0, 2*math.Pi)
	gc.Stroke()

	gc.SetLineWidth(clockThickness)

	// Outer dial
	gc.BeginPath()
	gc.ArcTo(cx, cy, radius, radius, 0, 2*math.Pi)
	gc.Stroke()

	// Inner dial
	gc.BeginPath()
	gc.ArcTo(cx, cy, radiusInner, radiusInner, 0, 2*math.Pi)
	gc.Stroke()

	// Draw second ticks
	for sec := 0; sec < 60; sec++ {
		if (sec%10) == 0 || sec <= 10 {
			deg := 360.0 * float64(45-sec) / 60.0

			// Small or Large ticks
			r1, r2 := radius, radius-(2*clockThickness)
			if sec == 45 || (sec%10) == 0 {
				r1, r2 = radius+(3*clockThickness), radius-(3*clockThickness)
			}

			sr1, cr1 := math.Sincos((deg - 0.25) * ToRad)
			sr2, cr2 := math.Sincos((deg + 0.25) * ToRad)

			gc.MoveTo(cx+(r1*cr1), cy+(r1*sr1))
			gc.LineTo(cx+(r2*cr1), cy+(r2*sr1))
			gc.LineTo(cx+(r2*cr2), cy+(r2*sr2))
			gc.LineTo(cx+(r1*cr2), cy+(r1*sr2))
			gc.Close()
		}
	}
	gc.FillStroke()

	// Clock arm
	sec := float64(v.Config.Start) - (float64(frame) / float64(v.Config.Format.FrameRate))

	// Only move on the second
	if !v.Config.ClockSmooth {
		s, f := math.Modf(sec)
		sec = s
		if f > 0 {
			sec = sec + 1
		}
	}

	deg := 360.0 * float64(45-sec) / 60.0
	log.Printf("frame %d start %d sec %f deg %f", frame, v.Config.Start, sec, deg)
	sr, cr := math.Sincos(deg * ToRad)
	gc.BeginPath()
	gc.MoveTo(cx+((radiusInner+10)*cr), cy+((radiusInner+10)*sr))
	gc.LineTo(cx+((radius-20)*cr), cy+((radius-20)*sr))
	gc.SetStrokeColor(colornames.Red)
	gc.FillStroke()
}

const (
	ToRad = math.Pi / 180.0
)
