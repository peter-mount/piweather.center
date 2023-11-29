package template

import (
	"errors"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"html/template"
	"math"
	"sort"
	"strings"
	"time"
)

func (m *Manager) PostInit() error {
	m.funcMap = template.FuncMap{
		"hhmm":           hhmm,
		"html":           html,
		"lower":          strings.ToLower,
		"upper":          strings.ToUpper,
		"replaceAll":     strings.ReplaceAll,
		"rfc3339":        rfc3339,
		"split":          split,
		"trim":           strings.TrimSpace,
		"trimPrefix":     trimPrefix,
		"utc":            utc,
		"min":            genCalc(math.Min),
		"max":            genCalc(math.Max),
		"add":            genCalc(value.Add),
		"subtract":       genCalc(value.Subtract),
		"multiply":       genCalc(value.Multiply),
		"divide":         genCalc(value.Divide),
		"sequence":       sequence,
		"defVal":         defVal,
		"dict":           dict,
		"decimalAlign":   NewDecimalAlign,
		"getReadingKeys": m.getReadingKeys,
		//"getReadingHistory": m.Latest.GetHistory,
		"getReading":           m.getReading,
		"getLatestReadingTime": m.getLatestReadingTime,
	}
	return nil
}

func (m *Manager) AddFunction(name string, handler interface{}) *Manager {
	m.funcMap[name] = handler
	return m
}

func genCalc(f func(float64, float64) float64) func(a, b interface{}) float64 {
	return func(a, b interface{}) float64 {
		af, _ := util.ToFloat64(a)
		bf, _ := util.ToFloat64(b)
		return f(af, bf)
	}
}

func rfc3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func utc(t time.Time) time.Time {
	return t.UTC()
}

func html(s string) template.HTML {
	return template.HTML(s)
}

func split(sep, s string) []string {
	a := strings.Split(s, sep)
	for len(a) < 3 {
		a = append(a, "")
	}
	return a
}

func trimPrefix(p, s string) string {
	return strings.TrimPrefix(s, p)
}

// hhmm takes "01:23:45.0000" and returns "01:23"
func hhmm(s string) string {
	v := strings.Split(s, ":")
	switch len(v) {
	case 0:
		return ""
	case 1:
		return v[0]
	default:
		return v[0] + ":" + v[1]
	}
}

func dict(values ...any) (map[string]any, error) {

	root := make(map[string]any)

	for i := 0; i < len(values); i += 2 {
		dict := root
		var key string
		switch v := values[i].(type) {
		case string:
			key = v
		case []string:
			for i := 0; i < len(v)-1; i++ {
				key = v[i]
				var m map[string]any
				v, found := dict[key]
				if found {
					m = v.(map[string]any)
				} else {
					m = make(map[string]any)
					dict[key] = m
				}
				dict = m
			}
			key = v[len(v)-1]
		default:
			return nil, errors.New("invalid dictionary key")
		}
		dict[key] = values[i+1]
	}

	return root, nil
}

func (m *Manager) getReadingKeys() []string {
	metrics := m.Latest.Metrics()
	sort.SliceStable(metrics, func(i, j int) bool {
		return strings.ToLower(metrics[i]) < strings.ToLower(metrics[j])
	})
	return metrics
}

func (m *Manager) getReading(name string) record.Record {
	r, _ := m.Latest.Latest(name)
	return r
}

func (m *Manager) getLatestReadingTime() time.Time {
	return m.Latest.LatestTime()
}

// seq start end step -> slice of values between start and step
func sequence(start, end, step float64) []float64 {
	step = math.Abs(step)
	if step == 0 {
		step = 1
	}

	var v []float64

	if start < end {
		for start <= end {
			v = append(v, start)
			start += step
		}
	} else {
		for start >= end {
			v = append(v, start)
			start -= step
		}
	}

	return v
}

// defVal returns v if not nil, otherwise d
func defVal(v *float64, d float64) float64 {
	if v == nil {
		return d
	}
	return *v
}
