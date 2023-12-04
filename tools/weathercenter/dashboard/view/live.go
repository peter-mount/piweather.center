package view

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/model"
	"github.com/peter-mount/piweather.center/tools/weathercenter/ws"
	"sort"
	"sync"
)

type Live struct {
	name      string                 // name of dashboard
	server    *Service               // Parent service
	dashboard *model.Dashboard       // Attached dashboard
	websocket *ws.Server             // websocket server for a dashboard
	js        map[string]interface{} // Javascript templates for this dashboard
	mutex     sync.Mutex
}

func (s *Service) newLiveServer(n string, d *model.Dashboard) *Live {
	l := &Live{
		name:   n,
		server: s,
	}

	l.newDashboard(d)

	// Listen for live metric updates sharing the single queue
	s.Server.Listener().Add(l.notify)

	return l
}

func (s *Live) getDashboard() *model.Dashboard {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.dashboard
}

// return the data required by the templates
func (s *Live) getData() any {
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
		"dash":  s.name,
		"board": s.dashboard,
		"js":    js,
	}
}

// Set the dashboard for this instance and update the state
func (s *Live) newDashboard(d *model.Dashboard) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.dashboard = d

	s.js = make(map[string]interface{})
	s.updateJs(d)

	s.initDashboard()

	// Enable websocket if the dashboard expects live updates
	// Note: Changing Dashboard.Live to false will not remove the websocket,
	// but this allows us to only create them for those that do need them
	if s.websocket == nil && d.Live {
		s.websocket = ws.NewServer()
		s.server.Rest.HandleFunc("/live/dash/"+s.name, s.websocket.Handle)
		go s.websocket.Run()
	}
}

// Populate the dashboard with current data
func (s *Live) initDashboard() {
	latest := s.server.Server.Latest
	r := &model.Response{}
	for _, n := range latest.Metrics() {
		m, exists := latest.Latest(n)
		if exists {
			metric := api.Metric{
				Metric:    n,
				Time:      m.Time,
				Unit:      m.Value.Unit().ID(),
				Value:     m.Value.Float(),
				Formatted: m.Value.String(),
			}
			s.dashboard.Process(metric, r)
		}
	}
}

// Update the list of javascript templates
func (s *Live) updateJs(c model.Instance) {
	t := c.GetType()
	n := "dash/" + t + ".js"
	if s.server.Template.HasTemplate(n) {
		s.js[t] = true
	}

	if d, ok := c.(*model.Dashboard); ok {
		s.updateJsl(d.Components)
	} else if d, ok := c.(*model.Container); ok {
		s.updateJsl(d.Components)
	}
}

func (s *Live) updateJsl(l model.ComponentList) {
	for _, c := range l {
		s.updateJs(c)
	}
}

func (s *Live) notify(m api.Metric) {
	if m.IsValid() {
		d := s.getDashboard()
		if d != nil {
			r := &model.Response{}

			d.Process(m, r)

			// Only send if we have something to update
			if b, valid := r.Json(); valid {
				s.websocket.Send(b)
			}
		}
	}
}
