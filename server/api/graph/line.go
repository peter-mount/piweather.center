package graph

import (
	"bytes"
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/chart/line"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"net/http"
	"time"
)

// serveLine generates a line graph
func (s *SVG) serveLine(start, end time.Time, ctx context.Context) error {
	r := rest.GetRest(ctx)

	id := ctx.Value("id").(string)

	readings := s.Store.GetHistoryBetween(id, start, end)
	if readings == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	var buf bytes.Buffer

	period := time2.PeriodOf(start, end)
	l := line.New()

	g := station.GraphFromContext(ctx)

	// resolve the datasource
	var dataSource util.DataSource
	calc := g.GetCalculatedValue()
	if calc != nil && calc.IsPseudo() {
		to, ok := value.GetUnit(calc.Use)
		if !ok {
			return fmt.Errorf("unit %q not defined", calc.Use)
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

	l.SetDefinition(g).
		Add(chart.NewUnitSource(id, dataSource)).
		SetPeriod(period).
		SetBounds(svg.NewRect(0, 0, svgWidth, svgHeight))

	svg.New(&buf, svgWidth, svgHeight, func(s svg.SVG) {
		graph.CSS(s)
		s.Draw(l)
	})

	r.Status(http.StatusOK).
		ContentType("image/svg+xml").
		Value(buf.Bytes())

	return nil
}
