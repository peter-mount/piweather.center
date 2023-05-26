package graph

import (
	"bytes"
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/chart/line"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/server/api"
	"github.com/peter-mount/piweather.center/server/store"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"net/http"
	"path"
	"sort"
	"strings"
	"time"
)

func init() {
	kernel.Register(&SVG{})
}

const (
	svgWidth  = 1024
	svgHeight = 132
)

// SVG provides the /api/svg endpoint which displays svg graphs for a metric
type SVG struct {
	Inbound *api.EndpointManager `kernel:"inject"`
	Store   *store.Store         `kernel:"inject"`
	Config  station.Config       `kernel:"inject"`
}

func (s *SVG) Start() error {
	return s.Config.Accept(station.NewVisitor().
		Sensors(s.registerSensors).
		Graph(s.registerGraph).
		WithContext(context.Background()))
}

// registerSensors adds endpoints for a Sensors object
func (s *SVG) registerSensors(ctx context.Context) error {
	sensors := station.SensorsFromContext(ctx)
	sensorsPath := path.Join("/svg", path.Join(strings.Split(sensors.ID, ".")...))

	// Work out what indices we want available
	lineCount := 0
	_ = station.NewVisitor().
		Graph(func(ctx context.Context) error {
			g := station.GraphFromContext(ctx)
			switch {
			case g.Line != nil:
				lineCount++
			}
			return nil
		}).
		WithContext(context.Background()).
		VisitSensors(sensors)

	if lineCount > 0 {
		err := s.Inbound.RegisterHttpEndpoint(
			"svg "+sensors.Name,
			path.Join(sensorsPath, "day"),
			sensors.ID,
			"Index last 24 hours",
			http.MethodGet,
			"html",
			task.Of(s.serveSensorDay).Using(sensors.WithContext),
		)
		if err != nil {
			return err
		}

		err = s.Inbound.RegisterHttpEndpoint(
			"svg "+sensors.Name,
			path.Join(sensorsPath, "today"),
			sensors.ID,
			"Index since midnight",
			http.MethodGet,
			"html",
			task.Of(s.serveSensorToday).Using(sensors.WithContext),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// registerGraph adds endpoints for a Graph object
func (s *SVG) registerGraph(ctx context.Context) error {
	sensors := station.SensorsFromContext(ctx)
	reading := station.ReadingFromContext(ctx)
	calc := station.CalculatedValueFromContext(ctx)
	g := station.GraphFromContext(ctx)

	id := ""
	if reading != nil {
		id = reading.ID
	} else if calc != nil {
		id = calc.ID
	}
	g.Path = path.Join("/svg", path.Join(strings.Split(id, ".")...))

	switch {
	case g.Line != nil:
		err := s.Inbound.RegisterHttpEndpoint(
			"svg "+sensors.Name,
			path.Join(g.Path, "day.svg"),
			id,
			"Line graph for last 24 hours",
			http.MethodGet,
			"svg",
			task.Of(s.serveDay).
				WithValue("id", id).
				//Using(reading.WithContext).
				//Using(calc.WithContext).
				Using(g.WithContext),
		)
		if err != nil {
			return err
		}

		err = s.Inbound.RegisterHttpEndpoint(
			"svg "+sensors.Name,
			path.Join(g.Path, "today.svg"),
			id,
			"Line graph since midnight",
			http.MethodGet,
			"svg",
			task.Of(s.serveToday).
				WithValue("id", id).
				//Using(reading.WithContext).
				//Using(calc.WithContext).
				Using(g.WithContext),
		)
		if err != nil {
			return err
		}

	default:
		// No Chart defined so remove path so we don't use it elsewhere
		g.Path = ""
	}
	return nil
}

func (s *SVG) serveSensorDay(ctx context.Context) error {
	return s.serveSensor(ctx, "day.svg")
}

func (s *SVG) serveSensorToday(ctx context.Context) error {
	return s.serveSensor(ctx, "today.svg")
}

func (s *SVG) serveSensor(ctx context.Context, img string) error {
	r := rest.GetRest(ctx)

	sensors := station.SensorsFromContext(ctx)

	// Sort keys so we have some sort of order with the results
	keys := append(sensors.ReadingsKeys(), sensors.CalculationsKeys()...)
	sort.SliceStable(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	var buf bytes.Buffer

	// Only include Graph with a path defined
	buf.WriteString("<html><body>")
	for _, k := range keys {
		graphs := sensors.GetGraph(k)
		if graphs != nil {
			for _, g := range graphs {
				if g.Path != "" {
					buf.WriteString("<object type=\"image/svg+xml\" data=\"" + g.Path + "/" + img + "\"></object>")
				}
			}
		}
	}
	buf.WriteString("</body></html>")

	r.Status(http.StatusOK).
		HTML().
		Value(buf.Bytes())

	return nil
}

func (s *SVG) serveHour(ctx context.Context) error {
	now := time.Now()
	start := now.Truncate(time.Hour)
	end := start.Add(time.Hour)
	return s.serve(start, end, ctx)
}

func (s *SVG) serveDay(ctx context.Context) error {
	// End is end of current hour
	end := time.Now().Truncate(time.Hour).Add(time.Hour)
	start := end.Add(-24 * time.Hour)
	return s.serve(start, end, ctx)
}

func (s *SVG) serveToday(ctx context.Context) error {
	now := time.Now()
	// Start at beginning of the current local day
	//
	// Note: truncate to hour then subtract hours to get the start.
	// It might look weird when you could truncate to day, but that truncate
	// seems to set it to 0h UTC, so if we are in BST (UTC+1) then the day
	// starts at 0100 and not 0000 midnight.
	//
	// TODO check this works for other timezones
	start := now.Truncate(time.Hour)
	start = start.Add(time.Hour * time.Duration(-start.Hour()))

	end := start.Add(time.Hour * 24)
	return s.serve(start, end, ctx)
}

func (s *SVG) serve(start, end time.Time, ctx context.Context) error {
	r := rest.GetRest(ctx)

	id := ctx.Value("id").(string)
	//reading := station.ReadingFromContext(ctx)
	//id := reading.ID

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
	if calc == nil {
		r.Status(http.StatusExpectationFailed)
		return nil
	}

	if calc.IsPseudo() {
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