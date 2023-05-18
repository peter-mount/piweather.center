package template

import (
	"errors"
	"fmt"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"html/template"
	"math"
	"strings"
	"time"
)

func (m *Manager) PostInit() error {
	m.funcMap = template.FuncMap{
		"hhmm":       hhmm,
		"html":       html,
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"replaceAll": strings.ReplaceAll,
		"rfc3339":    rfc3339,
		"split":      split,
		"trim":       strings.TrimSpace,
		"trimPrefix": trimPrefix,
		"utc":        utc,
		"valLeftPad": valLeftPad,
		"min":        genCalc(math.Min),
		"max":        genCalc(math.Max),
		"add":        genCalc(value.Add),
		"subtract":   genCalc(value.Subtract),
		"multiply":   genCalc(value.Multiply),
		"divide":     genCalc(value.Divide),
		"dict":       dict,
	}
	return nil
}

func (m *Manager) AddFunction(name string, handler interface{}) *Manager {
	m.funcMap[name] = handler
	return m
}

func valLeftPad(v interface{}) float64 {
	f, _ := util.ToFloat64(v)
	lp := len(fmt.Sprintf("%.0f", f))
	return float64(lp)
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
