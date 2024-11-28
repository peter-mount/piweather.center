package view

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	"github.com/peter-mount/piweather.center/tools/weathercenter"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/model"
	"github.com/peter-mount/piweather.center/tools/weathercenter/template"
	"github.com/peter-mount/piweather.center/util/config"
	cron2 "gopkg.in/robfig/cron.v2"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type Service struct {
	Cron       *cron.CronService     `kernel:"inject"`
	Rest       *rest.Server          `kernel:"inject"`
	Config     config.Manager        `kernel:"inject"`
	Template   *template.Manager     `kernel:"inject"`
	Server     *weathercenter.Server `kernel:"inject"`
	mutex      sync.Mutex
	dashDir    string
	dashboards map[string]*Live // loaded dashboard instances
	cronIds    map[string]int   // Map of cron ids
}

const (
	dashDir          = "dashboards"
	fileSuffix       = ".yaml"
	webPath          = "/dash/{dash:.{1,}}"
	stationHome      = "/s/{stationId}"
	stationHomeS     = "/s/{stationId}/"
	stationDashboard = "/s/{stationId}/{dash:.{1,}}"
)

func (s *Service) Start() error {
	s.dashboards = make(map[string]*Live)
	s.cronIds = make(map[string]int)

	s.dashDir = filepath.Join(s.Config.EtcDir(), dashDir)

	// Load existing dashboards
	err := walk.NewPathWalker().
		Then(func(path string, info os.FileInfo) error {
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			o := model.DashboardFactory()
			err = model.UnmarshalDashboard(b, o)
			if err == nil {
				s.setDashboard(s.stripPath(path), o.(*model.Dashboard))
			}
			return err
		}).
		PathHasSuffix(fileSuffix).
		IsFile().
		Walk(s.dashDir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// start watching for changes
	s.Config.WatchDirectory(dashDir, model.DashboardFactory, s.updateDashboard, model.UnmarshalDashboard)

	// Old static dashboard TODO remove later
	s.Rest.HandleFunc(webPath, func(writer http.ResponseWriter, request *http.Request) {
		dash := mux.Vars(request)["dash"]
		http.Redirect(writer, request, "/s/home/"+dash, http.StatusSeeOther)
	})

	// Station dashboards
	s.Rest.Handle(stationHome, s.showStationHome).Methods("GET")
	s.Rest.Handle(stationHomeS, s.showStationHome).Methods("GET")
	s.Rest.Handle(stationDashboard, s.showDashboard).Methods("GET")

	return nil
}

func (s *Service) stripPath(n string) string {
	n = strings.TrimPrefix(n, s.dashDir)
	return strings.TrimPrefix(strings.TrimSuffix(n, fileSuffix), "/")
}

func (s *Service) getLive(n string) *Live {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.dashboards[n]
}

func (s *Service) getDashboard(n string) *model.Dashboard {
	l := s.getLive(n)
	if l == nil {
		return nil
	}
	return l.dashboard
}

func (s *Service) setDashboard(n string, d *model.Dashboard) {
	if d == nil {
		return
	}

	// Force the type field, needed for template resolution
	d.Type = "dashboard"

	d.Init(*s.Server.DBServer)

	var useCron bool
	if d.Update != "" {
		id, err := s.Cron.AddFunc(d.Update, func() {
			d.Init(*s.Server.DBServer)
			// Make a new Uid so client refreshes
			d.CronSeq++
			uid := strings.Split(d.Uid, "-")
			d.Uid = uid[0] + "-" + strconv.Itoa(d.CronSeq)
		})
		if err == nil {
			d.CronId = int(id)
			useCron = true
			log.Printf("Cron: Adding %q %d", n, d.CronId)
		}
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Remove any existing cron job
	if oid, exists := s.cronIds[n]; exists {
		delete(s.cronIds, n)
		defer func() {
			log.Printf("Cron: Removing %q %d", n, oid)
			s.Cron.Remove(cron2.EntryID(oid))
		}()
	}
	if useCron {
		// record ID only if we are creating a new one.
		// Note: cronId can be 0 so we can't do this blindly,
		//hence the bool indicating we have actually created one
		s.cronIds[n] = d.CronId
	}

	l := s.dashboards[n]
	if l == nil {
		l = s.newLiveServer(n, d)
		s.dashboards[n] = l
	} else {
		l.newDashboard(d)
	}
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

// showStationHome handles /s/{stationId} which shows the "home" dashboard
func (s *Service) showStationHome(r *rest.Rest) error {
	return s.showDashboardImpl(r, "home")
}

// showDashboard handles /s/{stationId}/{dash} to allow a custom dashboard to be shown
func (s *Service) showDashboard(r *rest.Rest) error {
	return s.showDashboardImpl(r, r.Var("dash"))
}

func (s *Service) showDashboardImpl(r *rest.Rest, dash string) error {
	serverId := r.Var("stationId")

	live := s.getLive(dash)

	if live == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	if live.getDashboard().Refresh > 0 {
		r.AddHeader("Refresh", strconv.Itoa(live.getDashboard().Refresh))
	}

	data := live.getData()

	if serverId != "" {
		data["serverId"] = serverId
	}

	return s.Template.ExecuteTemplate(r, "dash/main.html", data)
}
