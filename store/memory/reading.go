package memory

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strconv"
	"strings"
	"time"
)

// Reading represents a raw reading in memory
type Reading struct {
	Name  string
	Value value.Value
	Time  time.Time
}

func (r *Reading) String() string {
	return strings.Join([]string{
		r.Name,
		strconv.FormatFloat(r.Value.Float(), 'f', 3, 64),
		strconv.Itoa(int(r.Time.UTC().Unix())),
	}, " ")
}

// Reducer performs reductions of data to reduce them to meaningful values
type Reducer struct {
	period time.Duration // Period between readings
}

type Entry struct {
	Reading
	Count int
}

func NewReducer(period time.Duration) *Reducer {
	return &Reducer{period: period}
}

func NewReducerMins(minutes int) *Reducer {
	return NewReducer(time.Minute * time.Duration(minutes))
}

func (r *Reducer) Reduce(d []*Reading, op func(float64, float64) float64, f func(*Entry) value.Value) ([]*Reading, error) {
	var entries []*Entry
	var e *Entry
	for _, reading := range d {
		if e != nil {
			if reading.Time.Sub(e.Time) > r.period {
				e = nil
			} else {
				rv, err := reading.Value.As(e.Value.Unit())
				if err != nil {
					return nil, err
				}

				e.Value = e.Value.Unit().Value(op(e.Value.Float(), rv.Float()))
				e.Count++
			}
		}

		if e == nil {
			e = &Entry{Reading: *reading, Count: 1}
			entries = append(entries, e)
		}
	}

	var ret []*Reading
	for _, e := range entries {
		e.Value = f(e)
		ret = append(ret, &e.Reading)
	}
	return ret, nil
}

func unity(e *Entry) value.Value { return e.Value }

func mean(e *Entry) value.Value {
	return e.Value.Unit().Value(e.Value.Float() / float64(e.Count))
}

func add(a, b float64) float64 { return a + b }

func (r *Reducer) Max(d []*Reading) ([]*Reading, error) {
	return r.Reduce(d, math.Max, unity)
}

func (r *Reducer) Min(d []*Reading) ([]*Reading, error) {
	return r.Reduce(d, math.Max, unity)
}

func (r *Reducer) Sum(d []*Reading) ([]*Reading, error) {
	return r.Reduce(d, add, unity)
}

func (r *Reducer) Mean(d []*Reading) ([]*Reading, error) {
	return r.Reduce(d, add, mean)
}
