package weather

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

func init() {
	packages.Register("weatherdb", &DB{})
}

type DB struct{}

func (_ DB) Connect(url string) *client.Client {
	return &client.Client{Url: url, Internal: true}
}

// Query short for Connect(url).Query(query)
func (d DB) Query(url, query string) (*api.Result, error) {
	return d.Connect(url).Query(query)
}

// LatestMetrics short for Connect(url).LatestMetrics()
func (d DB) LatestMetrics(url string) (*api.Response, error) {
	return d.Connect(url).LatestMetrics()
}

func (_ DB) NewMetric(m string, t time.Time, val float64, unit string) api.Metric {
	var v value.Value
	if u, exists := value.GetUnit(unit); exists {
		v = u.Value(val)
	}

	return api.Metric{
		Metric:    m,
		Time:      t,
		Unit:      unit,
		Value:     val,
		Formatted: v.String(),
		Unix:      t.Unix(),
	}
}
