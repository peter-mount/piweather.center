package station

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

// Axis represents the axis to a component, e.g. Gauge
type Axis struct {
	Pos   lexer.Position
	Min   float64 `parser:"('min' @('-'? Number) )?"` // Min value for the axis, defaults to 0
	Max   float64 `parser:"('max' @('-'? Number) )?"` // Max value for the axis, defaults to 100
	Ticks int     `parser:"('ticks' @(Number) )?"`    // Number of ticks to add after the origin
}

func (c *visitor[T]) Axis(d *Axis) error {
	var err error
	if d != nil {
		if c.axis != nil {
			err = c.axis(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initAxis(_ Visitor[*initState], d *Axis) error {
	var err error

	// ensure min < max
	if value.GreaterThan(d.Min, d.Max) {
		d.Min, d.Max = d.Max, d.Min
	}

	// default values
	if value.IsZero(d.Min) && value.IsZero(d.Max) {
		// Default to 0...100
		d.Min, d.Max = 0.0, 100.0
	}

	if d.Ticks == 0 {
		// Default to 10 ticks
		d.Ticks = 10
	}

	// verify station
	switch {
	case value.Equal(d.Min, d.Max):
		err = errors.Errorf(d.Pos, "Min and Max must not be the same")

	case d.Ticks < 0:
		err = errors.Errorf(d.Pos, "Ticks %d is invalid", d.Ticks)
	}

	return errors.Error(d.Pos, err)
}

func (b *builder[T]) Axis(f func(Visitor[T], *Axis) error) Builder[T] {
	b.axis = f
	return b
}

// AxisDef is the output of GenAxis generating a simple axis
type AxisDef struct {
	Delta  float64    // Difference offset between each tick
	Num    float64    // Number of ticks after the origin
	Format string     // format for fmt.Sprintf() to format values on the axis
	Ticks  []AxisTick // List of ticks. There will be Num+1 entries
}

// AxisTick represents each individual tick on the axis
type AxisTick struct {
	Index int     // index of this tick
	Label string  // Label
	Angle float64 // Angle of the tick, clockwise from the origin
}

// GenAxis generates an AxisDef based on this Axis. ang is the overall angle the axis will range over.
// e.g. ang=180 if the axis covers a semicircle.
func (a *Axis) GenAxis(ang float64) AxisDef {
	return GenAxis(a.Min, a.Max, float64(a.Ticks), ang)
}

func (a *Axis) EnsureWithin(v, delta, offset float64) float64 {
	return ((EnsureWithin(v, a.Min, a.Max) - a.Min) * delta) + offset
}

// GenAxis generates an AxisDef.
// min,max define the range of the axis,
// num the number of ticks after the origin,
// ang is the overall angle the axis will range over. e.g. ang=180 if the axis covers a semicircle.
func GenAxis(min, max, num, ang float64) AxisDef {
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

	r := AxisDef{
		Delta:  ang / dmm,
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

// EnsureWithin returns v so that it's always within min,max.
// e.g. if v<min then it returns min, v>max then max, otherwise v
func EnsureWithin(v, min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return math.Max(min, math.Min(v, max))
}

// AxisScale represents an axis which scales based on the value passed.
// This is primarily used for Rain Gauges where the axis can vary dependent on the amount of rain measured.
type AxisScale struct {
	Min    float64   // Minimum axis value
	Max    float64   // Maximum axis value
	Ticks  int       // Number of ticks after the origin
	Format string    // Format for fmt.Sprintf() to format values
	Points []float64 // slice of points for each tick, there will be Ticks+1 entries
	Labels []string  // Labels for each entry in Points
	Size   float64   // The size of the axis in the destination coordinate system
	Scale  float64   // scale to convert index in Points/Labels to the destination coordinate system
	Range  float64   // Range between Min and Max
}

// Contains true if the passed value is wtihin AxisScale
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
		{Min: 2, Max: 5, Ticks: 5},
		{Min: 1, Max: 2, Ticks: 4, Format: "%.1f"},
		{Min: 0.1, Max: 1, Ticks: 10, Format: "%.1f"},
		{Min: 0.01, Max: 0.1, Ticks: 10, Format: "%.2f"},
		{Min: 0, Max: 0.01, Ticks: 10, Format: "%.3f"},
	}
)

// AutoScale returns an AxisScale based on the input.
// min,max is the range required,
// size of the axis in the destination coordinate system
func AutoScale(min, max, size float64) AxisScale {
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
