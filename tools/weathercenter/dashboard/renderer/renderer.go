package renderer

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/state"
	"net/http"
)

type Renderer struct {
	Stations      *state.Stations `kernel:"inject"`
	renderVisitor station.Visitor[*State]
}

func (r *Renderer) Start() error {
	r.renderVisitor = station.NewBuilder[*State]().
		Container(Container).
		Gauge(Gauge).
		Dashboard(Dashboard).
		Value(Value).
		Build()
	return nil
}

func (r *Renderer) Render(stationId, dashboardId string) (string, int) {
	st := r.Stations.GetStation(stationId)
	if st == nil {
		return "", http.StatusNotFound
	}

	dash := st.GetDashboard(dashboardId)
	if dash == nil {
		return "", http.StatusNotFound
	}

	s := NewState(dash)

	err := r.renderVisitor.Clone().
		Set(s).
		Dashboard(dash.Dashboard())

	if err != nil {
		log.Printf("render %s:%s got %v", stationId, dashboardId, err)
		return "", http.StatusInternalServerError
	}

	return s.String(), http.StatusOK
}