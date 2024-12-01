package view

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	station3 "github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/renderer"
	"github.com/peter-mount/piweather.center/tools/weathercenter/template"
	"github.com/peter-mount/piweather.center/util/config"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Service struct {
	Cron     *cron.CronService `kernel:"inject"`
	Rest     *rest.Server      `kernel:"inject"`
	Config   config.Manager    `kernel:"inject"`
	Template *template.Manager `kernel:"inject"`
	//Server     *weathercenter.Server `kernel:"inject"`
	Stations   *station3.Stations `kernel:"inject"`
	Renderer   *renderer.Renderer `kernel:"inject"`
	DBServer   *string            `kernel:"flag,metric-db,DB url"`
	mutex      sync.Mutex
	dashDir    string
	parser     util.Parser[station.Stations]
	dashboards map[string]*Live // loaded dashboard instances
}

const (
	dashDir          = "stations"
	fileSuffix       = ".station"
	webPath          = "/dash/{dash:.{1,}}"
	stationHome      = "/s/{stationId}"
	stationHomeS     = "/s/{stationId}/"
	stationDashboard = "/s/{stationId}/{dash:.{1,}}"
)

func (s *Service) Start() error {
	s.parser = station.NewParser()

	s.dashboards = make(map[string]*Live)

	s.dashDir = filepath.Join(s.Config.EtcDir(), dashDir)

	// Load existing dashboards
	var files []string
	if err := walk.NewPathWalker().
		Then(func(path string, _ os.FileInfo) error {
			files = append(files, path)
			log.Printf("Found %q", path)
			return nil
		}).
		PathHasSuffix(fileSuffix).
		IsFile().
		Walk(s.dashDir); err != nil && !os.IsNotExist(err) {
		return err
	}

	// Load all the loadedStations
	if stations, err := station.NewParser().ParseFiles(files...); err != nil {
		return err
	} else {
		s.Stations.AddStations(stations)
	}

	// start watching for changes
	//s.Config.WatchDirectory(s.dashDir, model.DashboardFactory, s.updateDashboard, model.UnmarshalDashboard)
	//s.Config.WatchDirectoryParser(s.dashDir, model.DashboardFactory, s.updateDashboard, s.unmarshaller)

	// Old static dashboard TODO remove later
	s.Rest.HandleFunc(webPath, func(writer http.ResponseWriter, request *http.Request) {
		dash := mux.Vars(request)["dash"]
		http.Redirect(writer, request, "/s/home/"+dash, http.StatusSeeOther)
	})

	// Station dashboards
	s.Rest.Handle(stationHome, s.showStationHome).Methods("GET")
	s.Rest.Handle(stationHomeS, s.showStationHome).Methods("GET")
	s.Rest.Handle(stationDashboard, s.showDashboard).Methods("GET")

	return s.loadLatestMetrics()
}

func (s *Service) unmarshaller(name string) (any, error) {
	return s.parser.ParseFile(name)
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

func (s *Service) updateDashboard(event fsnotify.Event, o any) error {
	if strings.HasSuffix(event.Name, fileSuffix) {
		switch event.Op {
		case fsnotify.Create, fsnotify.Write:
			// FIXME need to work on this
			//d := o.(*station.Stations)

			//s.Stations.AddStations(d)

		case fsnotify.Remove:
			// FIXME need to work on this
			//return s.OnStations(func(stations *station.Stations) error {
			//	s.setStations(stations.Remove(event.Name))
			//	return nil
			//})

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

func (s *Service) showDashboardImpl(r *rest.Rest, dashName string) error {
	serverId := r.Var("stationId")

	content, status := s.Renderer.Render(serverId, dashName)

	//if live.getDashboard().Refresh > 0 {
	//	r.AddHeader("Refresh", strconv.Itoa(live.getDashboard().Refresh))
	//}
	r.AddHeader("Refresh", "10")

	content = `<html><head><title>test</title>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" name="viewport"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <link href="/` + station3.UID() + `/css/dash.css" rel="stylesheet"/>
</head><body>` + content + `</div></div>
</body></html>`
	r.Status(status).
		ContentType("text/html").
		Value([]byte(content))

	//live := s.getLive(serverId + "." + dash)
	//if live == nil {
	//	r.Status(http.StatusNotFound)
	//	return nil
	//}

	//data := dash.GetData()
	//
	//return s.Template.ExecuteTemplate(r, "dash/main.html", data)

	return nil
}

// loadLatestMetrics retrieves the current metrics from the DB server
func (s *Service) loadLatestMetrics() error {
	if *s.DBServer != "" {
		c := &client.Client{Url: *s.DBServer}
		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}
		if r != nil {
			s.Stations.Load(r.Metrics)
		}
	}
	return nil
}
