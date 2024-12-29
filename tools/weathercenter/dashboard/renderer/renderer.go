package renderer

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	station2 "github.com/peter-mount/piweather.center/station"
	"net/http"
)

type Renderer struct {
	Stations      *station2.Stations `kernel:"inject"`
	renderVisitor station.Visitor[*State]
}

func (r *Renderer) Start() error {
	r.renderVisitor = station.NewBuilder[*State]().
		Container(Container).
		Dashboard(Dashboard).
		Gauge(Gauge).
		MultiValue(MultiValue).
		Text(Text).
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

	if err != nil && !errors.IsVisitorStop(err) {
		log.Printf("render %s:%s got %v", stationId, dashboardId, err)
		return "", http.StatusInternalServerError
	}

	return s.String(), http.StatusOK
}
