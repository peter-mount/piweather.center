package util

import (
	"fmt"
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/piweather.center/weather/value"
	"strconv"
	"strings"
	"time"
)

const (
	na = "N/A"
)

// Sprintf is similar to fmt.Sprintf except it is customised for handling
// our own data types.
//
// It's primarily used by weatherbot with it's output formatting.
//
// The format comprises the following:
//
// %%	% character
//
// %d 	integer or floating point with 0 decimal places
//
// %e	The unit for Value's, "" otherwise
//
// %f 	floating point
//
// %.1f	floating point 1 decimal place
//
// %.2f	floating point 2 decimal place
//
// %.3f	floating point 3 decimal place
//
// %s	String value
//
// %T   Integer value in unix seconds - returns as "Jan 02 15:04:05 MST"
//
// %u 	Unit value with unit suffix
//
// %v	interface value
func Sprintf(f string, args ...interface{}) string {
	var r []string
	var arg interface{}

	for f != "" {
		i := strings.Index(f, "%")
		if i < 0 {
			r = append(r, f)
			f = ""
		} else {
			r = append(r, f[:i])
			f = f[i:]

			switch {
			case strings.HasPrefix(f, "%%"):
				r = append(r, "%")
				f = f[2:]

			case strings.HasPrefix(f, "%d"):
				arg, args = getArg(args)
				if arg != nil {
					if i, err := calculator.GetInt(arg); err != nil {
						r = append(r, err.Error())
					} else {
						r = append(r, strconv.Itoa(i))
					}
				} else {
					r = append(r, na)
				}
				f = f[2:]

			case strings.HasPrefix(f, "%e"):
				arg, args = getArg(args)
				if v, ok := arg.(value.Value); ok {
					r = append(r, strings.TrimSpace(v.Unit().Unit()))
				} else {
					r = append(r, "")
				}
				f = f[2:]

			case strings.HasPrefix(f, "%f"), strings.HasPrefix(f, "%.0f"),
				strings.HasPrefix(f, "%.1f"),
				strings.HasPrefix(f, "%.2f"),
				strings.HasPrefix(f, "%.3f"):

				// Split f against %...f, a[0] the format, a[1] the remainder
				a := strings.SplitN(f, "f", 2)
				format := a[0] + "f"
				f = a[1]

				arg, args = getArg(args)
				if arg != nil {
					if fv, err := calculator.GetFloat(arg); err != nil {
						r = append(r, err.Error())
					} else {
						r = append(r, fmt.Sprintf(format, fv))
					}
				} else {
					r = append(r, na)
				}

			case strings.HasPrefix(f, "%s"):
				arg, args = getArg(args)
				if arg != nil {
					s, err := calculator.GetString(arg)
					if err != nil {
						s = err.Error()
					}
					r = append(r, s)
				} else {
					r = append(r, na)
				}
				f = f[2:]

			case strings.HasPrefix(f, "%T"):
				arg, args = getArg(args)
				if v, ok := arg.(value.Value); ok {
					r = append(r, time.Unix(int64(v.Float()), 0).Format("Jan 02 15:04:05 MST"))
				} else {
					r = append(r, na)
				}
				f = f[2:]

			case strings.HasPrefix(f, "%u"):
				arg, args = getArg(args)
				if v, ok := arg.(value.Value); ok {
					r = append(r, v.String())
				} else {
					r = append(r, na)
				}
				f = f[2:]

			case strings.HasPrefix(f, "%v"):
				arg, args = getArg(args)
				r = append(r, fmt.Sprintf("%v", arg))
				f = f[2:]

			default:
				// Ignore and include the %
				r = append(r, "%")
				f = f[1:]
			}
		}
	}

	return strings.Join(r, "")
}

func getArg(args []interface{}) (interface{}, []interface{}) {
	if len(args) > 0 {
		return args[0], args[1:]
	}
	return nil, nil
}
