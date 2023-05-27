package graph

import (
	"context"
	"fmt"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

func (s *SVG) initChart(start, end time.Time, ctx context.Context, c chart.Chart) (bool, error) {
	id := ctx.Value("id").(string)

	g := station.GraphFromContext(ctx)

	readings := s.Store.GetHistoryBetween(id, start, end)
	if readings == nil {
		return false, nil
	}

	period := time2.PeriodOf(start, end)

	// resolve the datasource
	var dataSource util.DataSource
	calc := g.GetCalculatedValue()
	if calc != nil && calc.IsPseudo() {
		to, ok := value.GetUnit(calc.Use)
		if !ok {
			return false, fmt.Errorf("unit %q not defined", calc.Use)
		}

		t := calc.Sensors().Station().LatLong().Time(period.Start())

		// If we don't set both min and max then use PseudoDataSo
		min, max := g.GetMinMax()
		if min != nil && max != nil {
			dataSource, _ = util.LimitedPseudoDataSource(
				calc.Calculator(),
				period,
				to,
				time.Minute,
				to.Value(*min),
				to.Value(*max),
				t,
			)
		} else {
			dataSource, _ = util.PseudoDataSource(
				calc.Calculator(),
				period,
				to,
				time.Minute,
				t)
		}

	} else {
		dataSource = readings
	}

	c.SetDefinition(g).
		Add(chart.NewUnitSource(id, dataSource)).
		SetPeriod(period)

	return true, nil
}
