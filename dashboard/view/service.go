package view

import (
	"github.com/fsnotify/fsnotify"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	"github.com/peter-mount/piweather.center/dashboard/model"
	"github.com/peter-mount/piweather.center/dashboard/registry"
	"github.com/peter-mount/piweather.center/tools/weathercenter/template"
	"github.com/peter-mount/piweather.center/util/config"
	template2 "html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Service struct {
	Rest       *rest.Server      `kernel:"inject"`
	Config     config.Manager    `kernel:"inject"`
	Template   *template.Manager `kernel:"inject"`
	mutex      sync.Mutex
	dashDir    string
	dashboards map[string]*model.Dashboard // current dashboard
}

const (
	dashDir    = "dashboards"
	fileSuffix = ".yaml"
	webPath    = "/dash/{dash:.{1,}}"
)

func (s *Service) PostInit() error {
	s.Template.AddFunction("showComponent", s.showComponent)
	return nil
}

func (s *Service) Start() error {
	s.dashboards = make(map[string]*model.Dashboard)

	s.dashDir = filepath.Join(s.Config.EtcDir(), dashDir)

	// Load existing dashboards
	err := walk.NewPathWalker().
		Then(func(path string, info os.FileInfo) error {
			log.Printf("path %q n %q", path, s.stripPath(path))
			b, err := os.ReadFile(path)
			if err == nil {
				o := s.factory()
				err = registry.Unmarshal(b, o)
				if err == nil {
					s.setDashboard(s.stripPath(path), o.(*model.Dashboard))
				}
				if err != nil {
					return err
				}
			}
			return nil
		}).
		PathHasSuffix(fileSuffix).
		IsFile().
		Walk(s.dashDir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// start watching for changes
	s.Config.WatchDirectory(dashDir, s.factory, s.updateDashboard, registry.Unmarshal)

	// Service dashboards
	s.Rest.Handle(webPath, s.show).Methods("GET")

	return nil
}

func (s *Service) factory() any {
	return &model.Dashboard{}
}

func (s *Service) stripPath(n string) string {
	n = strings.TrimPrefix(n, s.dashDir)
	return strings.TrimPrefix(strings.TrimSuffix(n, fileSuffix), "/")
}

func (s *Service) getDashboard(n string) *model.Dashboard {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.dashboards[n]
}

func (s *Service) setDashboard(n string, d *model.Dashboard) *model.Dashboard {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	old := s.dashboards[n]

	if d != nil {
		// Force the type field, needed for template resolution
		d.Type = "dashboard"

		s.dashboards[n] = d
		model.Debug(d)
	} else {
		delete(s.dashboards, n)
	}

	return old
}

func (s *Service) updateDashboard(event fsnotify.Event, o any) error {
	if strings.HasSuffix(event.Name, fileSuffix) {
		switch event.Op {
		case fsnotify.Create, fsnotify.Write:
			d := o.(*model.Dashboard)
			s.setDashboard(s.stripPath(event.Name), d)

		case fsnotify.Remove:
			s.setDashboard(s.stripPath(event.Name), nil)

			// Default do nothing
		default:
			return nil
		}
	}

	return nil
}

func (s *Service) show(r *rest.Rest) error {
	dash := r.Var("dash")
	board := s.getDashboard(dash)

	if board == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	return s.Template.ExecuteTemplate(r, "dash/main.html", map[string]interface{}{
		"dash":  dash,
		"board": board,
	})
}

func (s *Service) showComponent(c registry.Component) (template2.HTML, error) {
	return s.Template.Template("dash/"+strings.ToLower(c.GetType())+".html", c)
}
