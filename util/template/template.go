package template

import (
	"bytes"
	"context"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"os"
	"time"
)

func (m *Manager) addTemplate(path string, _ os.FileInfo) error {
	name := path[len(m.rootDir)+1:]

	log.Printf("Loading template %q from %q", name, path)

	b, err := os.ReadFile(path)
	if err == nil {
		_, err = m.rootTemplate.New(name).
			Funcs(m.funcMap).
			Parse(string(b))
	}
	return err
}

// ExecuteTemplate executes the named template
func (m *Manager) ExecuteTemplate(r *rest.Rest, n string, d interface{}) error {

	// Form the root map. If d is compatible then use it rather than a new one.
	var m1 map[string]interface{}
	if m2, ok := d.(map[string]interface{}); ok {
		m1 = m2
	} else {
		m1 = map[string]interface{}{
			"data": d,
		}
	}
	m1["request"] = r.Request()
	m1["now"] = time.Now()
	m1["version"] = version.Version

	//r.CacheMaxAge(60)
	if v, exists := m1["generated"]; exists {
		r.SetDate("last-modified", v.(time.Time).Unix())
	}

	var buf bytes.Buffer
	if err := m.rootTemplate.ExecuteTemplate(&buf, n, m1); err != nil {
		return err
	}
	r.HTML().Value(buf.Bytes())
	return nil
}

// Render executes the named template. Similar to ExecuteTemplate but runs against a context
func (m *Manager) Render(ctx context.Context, n string, d interface{}) error {
	return m.ExecuteTemplate(rest.GetRest(ctx), n, d)
}

// Do adds a simple handler that renders a template without any other processing
func (m *Manager) Do(path, template string, methods ...string) *Manager {
	m.restService.Do(path, func(ctx context.Context) error {
		return m.Render(ctx, template, nil)
	}).Methods(methods...)
	return m
}
