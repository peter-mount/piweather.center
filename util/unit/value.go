package unit

import (
	time2 "github.com/peter-mount/piweather.center/util/time"
	"strconv"
	"time"
)

type Value struct {
	tp Type      // Type of value
	i  int       // value when tp==Int
	f  float64   // value when tp==Float
	s  string    // value when tp==String
	t  time.Time // value when tp==Time
}

type Type int8

const (
	None   = iota // No value, "" or nil
	Int           // Integer
	Float         // Floating Point
	String        // String
	Time          // Timestamp
)

func (v Value) IsNil() bool    { return v.tp == None }
func (v Value) IsInt() bool    { return v.tp == Int }
func (v Value) IsFloat() bool  { return v.tp == Float }
func (v Value) IsString() bool { return v.tp == String }
func (v Value) IsTime() bool   { return v.tp == Time }

func (v Value) Int() int {
	switch v.tp {
	case Int:
		return v.i
	case Float:
		return int(v.f)
	default:
		return 0
	}
}

func (v Value) Float() float64 {
	switch v.tp {
	case Int:
		return float64(v.i)
	case Float:
		return v.f
	default:
		return 0.0
	}
}

func (v Value) String() string {
	switch v.tp {
	case Int:
		return strconv.Itoa(v.i)
	case Float:
		return strconv.FormatFloat(v.f, 'f', -1, 64)
	case String:
		return v.s
	default:
		return ""
	}
}

func (v Value) Time() time.Time {
	var t time.Time
	t.IsZero()
	switch v.tp {
	case Time:
		return v.t
	default:
		return time.Time{}
	}
}

func ParseValue(s string) Value {
	if s == "" {
		return Value{tp: None}
	}

	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return Value{tp: Int, i: int(i)}
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return Value{tp: Float, f: f}
	}

	if t := time2.ParseTime(s); !t.IsZero() {
		return Value{tp: Time, t: t}
	}

	return Value{tp: String, s: s}
}
