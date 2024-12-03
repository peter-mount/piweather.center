package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Visitor[T any] interface {
	Axis(*Axis) error
	Calculation(*Calculation) error
	CalculationList(*CalculationList) error
	Component(*Component) error
	ComponentList(*ComponentList) error
	ComponentListEntry(*ComponentListEntry) error
	Container(*Container) error
	CronTab(*time.CronTab) error
	Current(*Current) error
	Dashboard(*Dashboard) error
	DashboardList(*DashboardList) error
	Expression(*Expression) error
	Forecast(*Forecast) error
	Function(*Function) error
	Gauge(*Gauge) error
	Load(*Load) error
	Location(*location.Location) error
	LocationExpression(*LocationExpression) error
	Metric(*Metric) error
	MetricList(*MetricList) error
	MetricPattern(*MetricPattern) error
	MultiValue(*MultiValue) error
	Station(*Station) error
	Stations(*Stations) error
	Text(*Text) error
	Unit(*units.Unit) error
	UseFirst(*UseFirst) error
	Value(*Value) error

	// Get returns the attached value T
	Get() T
	// Set the attached value T
	Set(T) Visitor[T]
	// Clone the visitor - used when you need to pass an alternate data object in multiple goroutines.
	// Note: The attached value T is not included in the clone
	Clone() Visitor[T]
}

type visitor[T any] struct {
	common[T]
	data T
}

type common[T any] struct {
	axis               func(Visitor[T], *Axis) error
	calculation        func(Visitor[T], *Calculation) error
	calculationList    func(Visitor[T], *CalculationList) error
	component          func(Visitor[T], *Component) error
	componentList      func(Visitor[T], *ComponentList) error
	componentListEntry func(Visitor[T], *ComponentListEntry) error
	container          func(Visitor[T], *Container) error
	crontab            func(Visitor[T], *time.CronTab) error
	current            func(Visitor[T], *Current) error
	dashboard          func(Visitor[T], *Dashboard) error
	dashboardList      func(Visitor[T], *DashboardList) error
	expression         func(Visitor[T], *Expression) error
	forecast           func(Visitor[T], *Forecast) error
	function           func(Visitor[T], *Function) error
	gauge              func(Visitor[T], *Gauge) error
	load               func(Visitor[T], *Load) error
	location           func(Visitor[T], *location.Location) error
	locationExpression func(Visitor[T], *LocationExpression) error
	metric             func(Visitor[T], *Metric) error
	metricList         func(Visitor[T], *MetricList) error
	metricPattern      func(Visitor[T], *MetricPattern) error
	multiValue         func(Visitor[T], *MultiValue) error
	station            func(Visitor[T], *Station) error
	stations           func(Visitor[T], *Stations) error
	text               func(Visitor[T], *Text) error
	unit               func(Visitor[T], *units.Unit) error
	useFirst           func(Visitor[T], *UseFirst) error
	value              func(Visitor[T], *Value) error
}

func (c *visitor[T]) Get() T {
	return c.data
}

func (c *visitor[T]) Set(data T) Visitor[T] {
	c.data = data
	return c
}

func (c *visitor[T]) Clone() Visitor[T] {
	return &visitor[T]{common: c.common}
}

func (c *visitor[T]) Axis(d *Axis) error {
	var err error
	if d != nil {
		if c.axis != nil {
			err = c.axis(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Calculation(d *Calculation) error {
	var err error
	if d != nil {
		if c.calculation != nil {
			err = c.calculation(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.CronTab(d.Every)
		}

		if err == nil {
			err = c.CronTab(d.ResetEvery)
		}

		if err == nil {
			err = c.Load(d.Load)
		}

		if err == nil {
			err = c.UseFirst(d.UseFirst)
		}

		if err == nil {
			err = c.Expression(d.Expression)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) CalculationList(d *CalculationList) error {
	var err error
	if d != nil {
		if c.calculationList != nil {
			err = c.calculationList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		for _, e := range d.Calculations {
			err = c.Calculation(e)
			if err != nil {
				break
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Component(d *Component) error {
	var err error
	if d != nil {
		if c.component != nil {
			err = c.component(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) ComponentList(d *ComponentList) error {
	var err error
	if d != nil {
		if c.componentList != nil {
			err = c.componentList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Entries {
				err = c.ComponentListEntry(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) ComponentListEntry(d *ComponentListEntry) error {
	var err error
	if d != nil {
		if c.componentListEntry != nil {
			err = c.componentListEntry(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Container(d.Container)
		}

		if err == nil {
			err = c.Forecast(d.Forecast)
		}

		if err == nil {
			err = c.Gauge(d.Gauge)
		}

		if err == nil {
			err = c.MultiValue(d.MultiValue)
		}

		if err == nil {
			err = c.Text(d.Text)
		}

		if err == nil {
			err = c.Value(d.Value)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Container(d *Container) error {
	var err error
	if d != nil {
		if c.container != nil {
			err = c.container(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.ComponentList(d.Components)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) CronTab(d *time.CronTab) error {
	var err error
	if d != nil {
		if c.crontab != nil {
			err = c.crontab(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Current(b *Current) error {
	var err error
	if b != nil {
		if c.current != nil {
			err = c.current(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (c *visitor[T]) Dashboard(d *Dashboard) error {
	var err error
	if d != nil {
		if c.dashboard != nil {
			err = c.dashboard(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.ComponentListEntry(d.Components)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) DashboardList(d *DashboardList) error {
	var err error
	if d != nil {
		if c.dashboardList != nil {
			err = c.dashboardList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Dashboards {
				err = c.Dashboard(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Expression(b *Expression) error {
	var err error
	if b != nil {
		if c.expression != nil {
			err = c.expression(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil && b.Current != nil {
			err = c.Current(b.Current)
		}

		if err == nil && b.Function != nil {
			err = c.Function(b.Function)
		}

		if err == nil && b.Metric != nil {
			err = c.Metric(b.Metric)
		}

		if err == nil && b.Location != nil {
			err = c.LocationExpression(b.Location)
		}

		if err == nil && b.Using != nil {
			err = c.Unit(b.Using)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (c *visitor[T]) Function(d *Function) error {
	var err error
	if d != nil {
		if c.function != nil {
			err = c.function(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Expressions {
				err = c.Expression(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}
func (c *visitor[T]) Forecast(d *Forecast) error {
	var err error
	if d != nil {
		if c.forecast != nil {
			err = c.forecast(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.Metric(d.Temperature)
		}

		if err == nil {
			err = c.Metric(d.Pressure)
		}

		if err == nil {
			err = c.Metric(d.WindDirection)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Gauge(d *Gauge) error {
	var err error
	if d != nil {
		if c.gauge != nil {
			err = c.gauge(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		if err == nil {
			err = c.Axis(d.Axis)
		}

		if err == nil {
			err = c.MetricList(d.Metrics)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Load(b *Load) error {
	var err error
	if b != nil {
		if c.load != nil {
			err = c.load(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (c *visitor[T]) Location(d *location.Location) error {
	var err error
	if d != nil && c.location != nil {
		err = c.location(c, d)
		if util.IsVisitorStop(err) {
			return nil
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) LocationExpression(d *LocationExpression) error {
	var err error
	if d != nil && c.locationExpression != nil {
		err = c.locationExpression(c, d)
		if util.IsVisitorStop(err) {
			return nil
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Metric(d *Metric) error {
	var err error
	if d != nil {
		if c.metric != nil {
			err = c.metric(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) MetricList(d *MetricList) error {
	var err error
	if d != nil {
		if c.metricList != nil {
			err = c.metricList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		for _, s := range d.Metrics {
			err = c.Metric(s)
			if err != nil {
				break
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) MetricPattern(d *MetricPattern) error {
	var err error
	if d != nil {
		if c.metricPattern != nil {
			err = c.metricPattern(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) MultiValue(d *MultiValue) error {
	var err error
	if d != nil {
		if c.multiValue != nil {
			err = c.multiValue(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.MetricPattern(d.Pattern)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Station(d *Station) error {
	var err error
	if d != nil {
		if c.station != nil {
			err = c.station(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Location(d.Location)
		}

		if err == nil {
			err = c.CalculationList(d.Calculations)
		}

		if err == nil {
			err = c.DashboardList(d.Dashboards)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Stations(d *Stations) error {
	var err error
	if d != nil {
		if c.stations != nil {
			err = c.stations(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, s := range d.Stations {
				err = c.Station(s)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Text(d *Text) error {
	var err error
	if d != nil {
		if c.text != nil {
			err = c.text(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Unit(d *units.Unit) error {
	var err error
	if d != nil {
		if c.unit != nil {
			err = c.unit(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) UseFirst(b *UseFirst) error {
	var err error
	if b != nil {
		if c.useFirst != nil {
			err = c.useFirst(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Metric(b.Metric)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (c *visitor[T]) Value(d *Value) error {
	var err error
	if d != nil {
		if c.value != nil {
			err = c.value(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.MetricList(d.Metrics)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}
