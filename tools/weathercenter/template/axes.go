package template

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

type Axis struct {
	Delta  float64
	Num    float64
	Format string
	Ticks  []AxisTick
}

type AxisTick struct {
	Index int
	Label string
	Angle float64
}

func genAxis(min, max, num, ang float64) Axis {
	if num < 1 {
		num = 1
	}
	if min > max {
		min, max = max, min
	}

	da := ang / num
	dmm := math.Abs(max - min)
	dv := dmm / num

	var f string
	switch {
	case value.LessThan(dv, 1):
		f = "%.1f"
	case value.LessThan(dv, 0.1):
		f = "%.2f"
	case value.LessThan(dv, 0.01):
		f = "%.3f"
	default:
		f = "%.0f"
	}

	r := Axis{
		Delta:  da,
		Num:    num,
		Format: f,
	}

	for v := min; value.LessThanEqual(v, max); v += dv {
		i := len(r.Ticks)
		r.Ticks = append(r.Ticks, AxisTick{
			Index: i,
			Label: fmt.Sprintf(f, v),
			Angle: float64(i) * da,
		})
	}

	return r
}

func ensureWithin(v, min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return math.Max(min, math.Min(v, max))
}

type AxisScale struct {
	Min    float64
	Max    float64
	Ticks  int
	Format string
	Points []float64
	Labels []string
	Size   float64
	Scale  float64
	Range  float64
}

func (as AxisScale) Contains(v float64) bool {
	return value.GreaterThan(v, as.Min) && value.LessThanEqual(v, as.Max)
}

var (
	autoScaleRange = []AxisScale{
		{Min: 900, Max: 1000.0, Ticks: 10},
		{Min: 800, Max: 900, Ticks: 9},
		{Min: 700, Max: 800, Ticks: 8},
		{Min: 600, Max: 700, Ticks: 7},
		{Min: 500, Max: 600, Ticks: 6},
		{Min: 400, Max: 500, Ticks: 5},
		{Min: 300, Max: 400, Ticks: 4},
		{Min: 200, Max: 300, Ticks: 6},
		{Min: 100, Max: 200, Ticks: 10},
		{Min: 50, Max: 100, Ticks: 10},
		{Min: 40, Max: 50, Ticks: 10},
		{Min: 30, Max: 40, Ticks: 8},
		{Min: 20, Max: 30, Ticks: 6},
		{Min: 10, Max: 20, Ticks: 4},
		{Min: 5, Max: 10, Ticks: 10},
		{Min: 1, Max: 5, Ticks: 5},
		{Min: 0.1, Max: 1, Ticks: 10, Format: "%.1f"},
		{Min: 0.01, Max: 0.1, Ticks: 10, Format: "%.2f"},
		{Min: 0, Max: 0.01, Ticks: 10, Format: "%.3f"},
	}
)

func autoScale(min, max, size float64) AxisScale {
	var r AxisScale

	if min > max {
		min, max = max, min
	}
	d := max - min
	for _, l := range autoScaleRange {
		if l.Contains(d) {
			r = l
		}
	}

	// No ticks then copy a default
	if r.Ticks == 0 {
		r.Max = (math.Floor(max/autoScaleRange[0].Max) + 1) * autoScaleRange[0].Max
		r.Ticks = 10
	}

	r.Min = math.Min(r.Min, min)
	r.Max = math.Max(r.Max, max)
	r.Size = size
	r.Range = r.Max - r.Min
	r.Scale = r.Size / r.Range
	if r.Format == "" {
		r.Format = "%.0f"
	}

	ft := float64(r.Ticks)
	dy := size / ft
	dt := r.Range / ft
	for i := 0; i <= r.Ticks; i++ {
		fi := float64(i)
		r.Points = append(r.Points, dy*fi)
		r.Labels = append(r.Labels, fmt.Sprintf(r.Format, min+(dt*fi)))
	}

	return r
}
