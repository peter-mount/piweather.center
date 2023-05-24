package svg

import (
	"github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

type Projection struct {
	x0, y0, x1, y1 float64 // Dimensions of area
	maxX, maxY     float64 // x axis range
	minX, minY     float64 // y axis range
	yOrigin        float64 // y coord of y origin
	width, height  float64 // width,height in pixels
	xScale, yScale float64 // x & y axis scale
	stepX          float64 // step size used in NearestX
	stepY          float64 // step size used in NearestY
}

func NewProjection(x0, y0, x1, y1 float64) *Projection {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return &Projection{
		x0:    x0,
		y0:    y0,
		x1:    x1,
		y1:    y1,
		stepX: 1.0,
		stepY: 1.0,
	}
}

func (p *Projection) SetPeriod(period time.Period) *Projection {
	return p.SetXRange(0, period.DurationMinutes())
}

func (p *Projection) SetXRange(min, max float64) *Projection {
	p.minX, p.maxX = min, max
	return p
}

func (p *Projection) GetXRange() (float64, float64) { return p.minX, p.maxX }

func (p *Projection) SetValueRange(r *value.Range) *Projection {
	return p.SetYRange(r.Min().Float(), r.Max().Float())
}

func (p *Projection) SetYRange(min, max float64) *Projection {
	if min > max {
		min, max = max, min
	}
	p.minY, p.maxY = min, max
	return p
}

func (p *Projection) GetYRange() (float64, float64) { return p.minY, p.maxY }

func (p *Projection) ZeroY() *Projection {
	if p.minY > 0 {
		p.minY = 0
	}
	return p
}

func (p *Projection) NearestX(step float64) *Projection {
	p.stepX = step
	p.minX, p.maxX = Nearest(p.minX, p.maxX, step)
	return p
}

func (p *Projection) StepX() float64 { return p.stepX }

func (p *Projection) NearestY(step float64) *Projection {
	p.stepY = step
	p.minY, p.maxY = Nearest(p.minY, p.maxY, step)
	return p
}

func (p *Projection) StepY() float64 { return p.stepY }

func Nearest(minVal, maxVal, step float64) (float64, float64) {
	return math.Floor(minVal/step) * step, math.Ceil(maxVal/step) * step
}

func (p *Projection) Width() float64 { return p.width }

func (p *Projection) Height() float64 { return p.height }

func (p *Projection) Build() *Projection {
	p.width = p.x1 - p.x0
	p.height = p.y1 - p.y0

	p.xScale = p.width / (p.maxX - p.minX)
	p.yScale = p.height / (p.maxY - p.minY)

	p.yOrigin = p.y1 + (p.minY * p.yScale)

	return p
}

func (p *Projection) X0() float64 { return p.x0 }

func (p *Projection) Y0() float64 { return p.y0 }

func (p *Projection) X1() float64 { return p.x1 }

func (p *Projection) Y1() float64 { return p.y1 }

func (p *Projection) Xc() float64 { return p.x0 + (p.width / 2) }

func (p *Projection) Yc() float64 { return p.y0 + (p.height / 2) }

func (p *Projection) Rect() (float64, float64, float64, float64) { return p.x0, p.y0, p.x1, p.y1 }

func (p *Projection) XScale() float64 { return p.xScale }

func (p *Projection) YScale() float64 { return p.yScale }

func (p *Projection) YOrigin() float64 { return p.yOrigin }

func (p *Projection) Project(x, y float64) (float64, float64) {
	return p.x0 + (p.xScale * (x - p.minX)),
		p.y1 - ((y - p.minY) * p.yScale)
}

func (p *Projection) InsideX(x float64) bool { return x >= p.x0 && x <= p.x1 }

func (p *Projection) InsideY(y float64) bool { return y >= p.y0 && y <= p.y1 }

// YAxisTicks returns minY, maxY values and stepSize for plotting ticks on the Y-axis
func (p *Projection) YAxisTicks() (float64, float64, float64) {
	step := CalculateStep(p.minY, p.maxY)
	minV, maxV := Nearest(p.minY, p.maxY, step)
	return minV, maxV, step
}

// XAxisTicks returns minX, maxX values and stepSize for plotting ticks on the X-axis
func (p *Projection) XAxisTicks() (float64, float64, float64) {
	step := CalculateStep(p.minX, p.maxX)

	d := p.maxX - p.minX
	switch {
	// 1 hour
	case d <= 60:
		step = 10
	// 1 day in minutes
	case d <= 1440:
		step = 60
	}

	minV, maxV := Nearest(p.minX, p.maxX, step)
	return minV, maxV, step
}

func CalculateStep(min, max float64) float64 {
	if max < min {
		min, max = max, min
	}

	vRange := max - min
	e := 1.0
	for i := 1; i < 10; i++ {
		if vRange < e {
			// return previous e value
			return e / 10
		}
		e = e * 10.0
	}
	// range is >= 1e10 so default to 1e10
	return e
}
