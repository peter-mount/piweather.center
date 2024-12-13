package view

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	station2 "github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/renderer"
	"github.com/peter-mount/piweather.center/tools/weathercenter/ws"
	"sync"
)

// Live handles sending metrics to clients over a websocket
type Live struct {
	mutex     sync.Mutex
	server    *Service            // Parent service
	dashboard *station2.Dashboard // Attached dashboard
	websocket *ws.Server          // websocket server for a dashboard
	//js        map[string]interface{} // Javascript templates for this dashboard
	live bool // always matches Dashboard.Live
}

func (s *Service) UpdateJS(stations *station.Stations) error {
	type state struct {
		station *station.Station
	}

	return station.NewBuilder[*state]().
		Station(func(v station.Visitor[*state], s *station.Station) error {
			v.Get().station = s
			return nil
		}).
		Dashboard(func(v station.Visitor[*state], d *station.Dashboard) error {
			st := v.Get()

			if d.Live {
				// reset any existing Live, or create a new one if the first time
				s.getOrCreateLive(st.station.Name, d.Name).
					reset(true)
			} else {
				// reset and disable any existing Live if it was previously active
				s.GetLive(st.station.Name, d.Name).
					reset(false)
			}

			return util.VisitorStop
		}).
		//ComponentListEntry(func(v station.Visitor[*state], c *station.ComponentListEntry) error {
		//	t := c.GetType()
		//	if renderer.HasJavaScript(t) {
		//		v.Get().js[t] = true
		//	}
		//	return nil
		//}).
		Build().
		Set(&state{}).
		Stations(stations)
}

func (s *Service) getDashKey(stN, dashN string) string {
	if dash := s.Stations.GetStation(stN).GetDashboard(dashN); dash == nil {
		return ""
	}
	return stN + ":" + dashN
}

func (s *Service) GetLive(stN, dashN string) *Live {
	k := s.getDashKey(stN, dashN)
	if k == "" {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.dashboards[k]
}

func (s *Service) getOrCreateLive(stN, dashN string) *Live {
	k := s.getDashKey(stN, dashN)
	if k == "" {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	l, exists := s.dashboards[k]

	if !exists {
		l = &Live{server: s, websocket: ws.NewServer()}
		s.dashboards[k] = l

		s.Rest.HandleFunc(renderer.LiveWsPath(stN, dashN), l.websocket.Handle)

		go l.websocket.Run()
	}

	return l
}

// reset sets Live so it's either disabled or active depending on the new Dashboard state
func (s *Live) reset(live bool /*, js map[string]interface{}*/) {
	if s != nil {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.live = live
		//s.js = js
	}
}

func (s *Live) IsLive() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.live
}

func (s *Live) Notify(resp *station2.Response) {
	if s.IsLive() && resp.IsValid() {
		// Only send if we have something to update
		if b, valid := resp.Json(); valid {
			s.websocket.Send(b)
		}
	}
}
