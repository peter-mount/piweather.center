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

type Operation func(float64, float64) float64

type Finalizer func(*Entry) value.Value

type Predicate func(float64) bool

func (a Predicate) Or(b Predicate) Predicate {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(f float64) bool {
		return a(f) || b(f)
	}
}

func (a Predicate) And(b Predicate) Predicate {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(f float64) bool {
		return a(f) && b(f)
	}
}

func (a Predicate) Not() Predicate {
	if a == nil {
		return nil
	}
	return func(f float64) bool {
		return !a(f)
	}
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

func (r *Reducer) Reduce(d []*Reading, op Operation, f Finalizer) ([]*Reading, error) {
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

// Filter allows for a Reading set to be filtered, removing entries
// that do not match a set of criteria.
//
// Usually this is to remove readings that are abnormal, due to sensor errors etc.
type Filter struct {
}

func NewFilter() *Filter {
	return &Filter{}
}

func (_ Filter) Filter(d []*Reading, p Predicate) []*Reading {
	var ret []*Reading
	for _, reading := range d {
		if p(reading.Value.Float()) {
			ret = append(ret, reading)
		}
	}
	return ret
}

func (f Filter) Min(d []*Reading, min float64) []*Reading {
	return f.Filter(d, func(f float64) bool {
		return value.GreaterThanEqual(f, min)
	})
}

func (f Filter) Max(d []*Reading, max float64) []*Reading {
	return f.Filter(d, func(f float64) bool {
		return value.LessThanEqual(f, max)
	})
}

func (f Filter) Within(d []*Reading, min, max float64) []*Reading {
	return f.Filter(d, func(f float64) bool {
		return value.Within(f, min, max)
	})
}

func (f Filter) Without(d []*Reading, min, max float64) []*Reading {
	return f.Filter(d, func(f float64) bool {
		return value.Without(f, min, max)
	})
}
