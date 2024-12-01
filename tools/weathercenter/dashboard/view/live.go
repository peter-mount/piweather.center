package view

import (
	"github.com/peter-mount/piweather.center/config/station"
	station2 "github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/tools/weathercenter/ws"
	"sync"
)

type Live struct {
	server    *Service               // Parent service
	dashboard *station2.Dashboard    // Attached dashboard
	websocket *ws.Server             // websocket server for a dashboard
	js        map[string]interface{} // Javascript templates for this dashboard
	mutex     sync.Mutex
}

var (
	updateJsVisitor = station.NewBuilder[*Live]().
		ComponentListEntry(func(v station.Visitor[*Live], c *station.ComponentListEntry) error {
			s := v.Get()
			t := c.GetType()
			if t != "" {
				n := "dash/" + t + ".js"
				if s.server.Template.HasTemplate(n) {
					s.mutex.Lock()
					defer s.mutex.Unlock()
					s.js[t] = true
				}
			}
			return nil
		}).
		Build()
)

/*

func (s *Service) newLiveServer(d *station.Dashboard) *Live {
	l := &Live{
		server: s,
	}

	log.Printf("1")
	l.newDashboard(d)

	// Listen for live metric updates sharing the single queue
	//s.Server.Listener().Add(l.notify)

	return l
}

func (s *Live) getStation() *station.Station {
	return s.getDashboard().Station()
}

func (s *Live) getDashboard() *station.Dashboard {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.dashboard
}

// return the data required by the templates
func (s *Live) getData() map[string]interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var js []string
	for k, _ := range s.js {
		js = append(js, k)
	}
	sort.SliceStable(js, func(i, j int) bool {
		return js[i] < js[j]
	})

	return map[string]interface{}{
		"dash":  s.dashboard.Dashboard().Name,
		"board": s.dashboard,
		"js":    js,
	}
}

// Set the dashboard for this instance and update the station
func (s *Live) newDashboard(d *station.Dashboard) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.dashboard = d

	s.js = make(map[string]interface{})
	s.updateJs()

	// Enable websocket if the dashboard expects live updates
	// Note: Changing Dashboard.Live to false will not remove the websocket,
	// but this allows us to only create them for those that do need them
	if s.websocket == nil && d.Dashboard().Live {
		s.websocket = ws.NewServer()
		s.server.Rest.HandleFunc(strings.Join([]string{
			"/live/dash",
			d.Station().Station().Name,
			d.Dashboard().Name,
		}, "/"),
			s.websocket.Handle)

		go s.websocket.Run()
	}
}

// Update the list of javascript templates
func (s *Live) updateJs() {
	_ = updateJsVisitor.Clone().
		Set(s).
		Dashboard(s.dashboard.Dashboard())
}

func (s *Live) notify(m api.Metric) {
	log.Printf("notify %q %.2f", m.Metric, m.Value)
	if m.IsValid() {
		d := s.getDashboard()
		if d != nil {
			resp := s.server.Stations.Notify(m)

			// Only send if we have something to update
			if b, valid := resp.Json(); valid {
				s.websocket.Send(b)
			}
		}
	}
}
*/
