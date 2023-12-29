package api

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

type WindRose struct {
	Count        int               `json:"count"`        // Total number of entries
	Min          float64           `json:"min"`          // Min speed
	Max          float64           `json:"max"`          // Max speed
	MaxPerBucket int               `json:"maxPerBucket"` // Max count in each bucket
	Steps        []*WindRoseStep   `json:"steps"`        // step for each compass point
	Buckets      []*WindRoseBucket `json:"buckets"`      // bucket for each compass point
}

type WindRoseStep struct {
	Value   float64 `json:"value"`   // Max value in this step
	Count   int     `json:"count"`   // Total number of entries in this step
	Percent float64 `json:"percent"` // Percentage of entries in this step
}

func (s WindRoseStep) contains(f float64) bool {
	return value.LessThanEqual(f, s.Value)
}

type WindRoseBucket struct {
	Index  int       `json:"index"` // Index of bucket
	Count  int       `json:"count"` // Number of values in this bucket
	Max    float64   `json:"max"`   // Max value in bucket
	Steps  []int     `json:"steps"` // Number of entries per step
	Values []float64 `json:"-"`     // Values within this bucket
}

func (r *Result) AddWindRose(wr *WindRose) {
	r.WindRose = append(r.WindRose, wr)
}

func NewWindRose() *WindRose {
	// Note: we have 6 steps but 7 here to account for fence post problem
	w := &WindRose{}

	for i := 0; i < 7; i++ {
		w.Steps = append(w.Steps, &WindRoseStep{})
	}

	for i := 0; i < 16; i++ {
		w.Buckets = append(w.Buckets, &WindRoseBucket{Index: i})
	}

	return w
}

func (w *WindRose) Add(degree, speed float64) {
	// For zero speed just record under the first step
	if value.IsZero(speed) {
		w.Steps[0].Count++
		return
	}

	if speed < w.Min || w.Count == 0 {
		w.Min = speed
	}

	if speed > w.Max {
		w.Max = speed
	}

	w.Count++

	// Add to the relevant bucket.
	//
	// Here the bucket is a 22.5° swathe centered on one of the 16 compass points.
	//
	// We ensure that degree is positive but are not concerned of degree>=360 as that's managed by the %16
	for degree < 0 {
		degree += 360
	}
	w.Buckets[int((degree+11.25)/22.5)%16].Add(speed)
}

func (b *WindRoseBucket) Add(speed float64) {
	b.Values = append(b.Values, speed)
	b.Count++
	if speed > b.Max {
		b.Max = speed
	}
}

func (w *WindRose) Finalise() {
	if w.Count == 0 {
		return
	}

	// Set the ranges of each step.
	stepSize := (w.Max - w.Min) / 6.0
	for i := 1; i < len(w.Steps); i++ {
		w.Steps[i].Value = w.Steps[i-1].Value + stepSize
	}

	// Now finalise each bucket
	for i := 0; i < len(w.Buckets); i++ {
		w.Buckets[i].finalise(i, w)

		if w.Buckets[i].Count > w.MaxPerBucket {
			w.MaxPerBucket = w.Buckets[i].Count
		}
	}

	// Finalise the steps
	for i := 0; i < len(w.Steps); i++ {
		w.Steps[i].Percent = math.Floor(10000*float64(w.Steps[i].Count)/float64(w.Count)) / 10000
	}
}

func (b *WindRoseBucket) finalise(id int, w *WindRose) {
	b.Steps = make([]int, len(w.Steps))

	// Allocate each value to it's appropriate step
	for i := 0; i < len(b.Values); i++ {
		sid := w.findStep(b.Values[i])
		b.Steps[sid]++
		w.Steps[sid].Count++
	}
}

func (w *WindRose) findStep(f float64) int {
	for i := 0; i < len(w.Steps); i++ {
		if w.Steps[i].contains(f) {
			return i
		}
	}
	return len(w.Steps) - 1
}
