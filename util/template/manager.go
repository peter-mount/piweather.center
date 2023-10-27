package template

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	"github.com/peter-mount/piweather.center/store"
	"html/template"
	"os"
	"path"
	"path/filepath"
)

type Manager struct {
	restService  *rest.Server `kernel:"inject"`
	webRoot      *string      `kernel:"flag,webroot,Web root directory"`
	Store        store.Store  `kernel:"inject"`
	rootTemplate *template.Template
	funcMap      template.FuncMap
	rootDir      string
	disabled     bool
}

func (m *Manager) Disable() {
	m.disabled = true
}

func (m *Manager) GetRootDir() string { return m.rootDir }

func (m *Manager) Start() error {
	if *m.webRoot != "" {
		m.rootDir = path.Join(*m.webRoot, "templates")
	} else {
		m.rootDir = path.Join(filepath.Dir(os.Args[0]), "../lib/web/templates")
	}

	if m.disabled {
		return nil
	}

	log.Printf("Loading templates in %q", m.rootDir)

	m.rootTemplate = template.New("")

	return walk.NewPathWalker().
		Then(m.addTemplate).
		PathHasSuffix(".html").
		IsFile().
		Walk(m.rootDir)
}
