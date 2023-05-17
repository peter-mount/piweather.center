package template

import (
	"html/template"
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
	}
	return nil
}

func (m *Manager) AddFunction(name string, handler interface{}) *Manager {
	m.funcMap[name] = handler
	return m
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
