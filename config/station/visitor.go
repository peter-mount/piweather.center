package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Visitor[T any] interface {
	Component(*Component) error
	ComponentList(*ComponentList) error
	ComponentListEntry(*ComponentListEntry) error
	Container(*Container) error
	CronTab(*time.CronTab) error
	Dashboard(*Dashboard) error
	DashboardList(*DashboardList) error
	Gauge(*Gauge) error
	Location(*location.Location) error
	Metric(*Metric) error
	MetricList(*MetricList) error
	MetricPattern(*MetricPattern) error
	MultiValue(*MultiValue) error
	Station(*Station) error
	Stations(*Stations) error
	Value(*Value) error

	// Get returns the attached value T
	Get() T
	// Set the attached value T
	Set(T) Visitor[T]
}

type visitor[T any] struct {
	common[T]
	data T
}

type common[T any] struct {
	component          func(Visitor[T], *Component) error
	componentList      func(Visitor[T], *ComponentList) error
	componentListEntry func(Visitor[T], *ComponentListEntry) error
	container          func(Visitor[T], *Container) error
	crontab            func(Visitor[T], *time.CronTab) error
	dashboard          func(Visitor[T], *Dashboard) error
	dashboardList      func(Visitor[T], *DashboardList) error
	gauge              func(Visitor[T], *Gauge) error
	location           func(Visitor[T], *location.Location) error
	metric             func(Visitor[T], *Metric) error
	metricList         func(Visitor[T], *MetricList) error
	metricPattern      func(Visitor[T], *MetricPattern) error
	multiValue         func(Visitor[T], *MultiValue) error
	station            func(Visitor[T], *Station) error
	stations           func(Visitor[T], *Stations) error
	unit               func(Visitor[T], *units.Unit) error
	value              func(Visitor[T], *Value) error
}

func (c *visitor[T]) Get() T {
	return c.data
}

func (c *visitor[T]) Set(data T) Visitor[T] {
	c.data = data
	return c
}

func (c *visitor[T]) Component(d *Component) error {
	var err error
	if d != nil && c.component != nil {
		err = c.component(c, d)
	}
	return err
}

func (c *visitor[T]) ComponentList(d *ComponentList) error {
	var err error
	if d != nil {
		if c.componentList != nil {
			err = c.componentList(c, d)
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
		}
		if util.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = c.Container(d.Container)
		}

		if err == nil {
			err = c.Gauge(d.Gauge)
		}

		if err == nil {
			err = c.MultiValue(d.MultiValue)
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
		}
		if util.IsVisitorStop(err) {
			return nil
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
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Dashboard(d *Dashboard) error {
	var err error
	if d != nil {
		if c.dashboard != nil {
			err = c.dashboard(c, d)
		}
		if util.IsVisitorStop(err) {
			return nil
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
		}
		if util.IsVisitorStop(err) {
			return nil
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

func (c *visitor[T]) Gauge(d *Gauge) error {
	var err error
	if d != nil {
		if c.gauge != nil {
			err = c.gauge(c, d)
		}
		if util.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		if err == nil {
			err = c.MetricList(d.Metrics)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Location(d *location.Location) error {
	var err error
	if d != nil && c.location != nil {
		err = c.location(c, d)

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Metric(d *Metric) error {
	var err error
	if d != nil {
		if c.metric != nil {
			err = c.metric(c, d)
		}
		if util.IsVisitorStop(err) {
			return nil
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
		}
		if util.IsVisitorStop(err) {
			return nil
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
		}
		if util.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = c.Location(d.Location)
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
		}
		if util.IsVisitorStop(err) {
			return nil
		}

		for _, s := range d.Stations {
			err = c.Station(s)
			if err != nil {
				break
			}
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
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Value(d *Value) error {
	var err error
	if d != nil {
		err = c.Component(d.Component)

		if err == nil && c.value != nil {
			err = c.value(c, d)
		}
		if util.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = c.MetricList(d.Metrics)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}
