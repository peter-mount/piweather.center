package model

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/util/task"
)

type Visitor interface {
	VisitStations(s *Stations) error
	VisitStation(s *Station) error
	VisitSensors(s *Sensors) error
	VisitReading(s *Reading) error
}

type Visitable interface {
	Accept(v Visitor) error
}

type visitorCommon struct {
	stations task.Task
	station  task.Task
	sensors  task.Task
	reading  task.Task
}

type visitor struct {
	visitorCommon
	ctx context.Context
}

func (v *visitor) VisitStations(s *Stations) error {
	oldCtx := v.ctx
	newCtx := s.WithContext(v.ctx)
	v.ctx = newCtx
	defer func() {
		v.ctx = oldCtx
	}()

	if err := v.stations.Do(v.ctx); err != nil {
		return err
	}

	for id, c := range *s {
		v.ctx = context.WithValue(newCtx, "StationId", id)
		if err := c.Accept(v); err != nil {
			return err
		}
	}

	return nil
}

func (v *visitor) VisitStation(s *Station) error {
	oldCtx := v.ctx
	newCtx := s.WithContext(v.ctx)
	v.ctx = newCtx
	defer func() {
		v.ctx = oldCtx
	}()

	if err := v.station.Do(v.ctx); err != nil {
		return err
	}

	for id, c := range s.Sensors {
		v.ctx = context.WithValue(newCtx, "SensorId", id)
		if err := c.Accept(v); err != nil {
			return err
		}
	}

	return nil
}

func (v *visitor) VisitSensors(s *Sensors) error {
	oldCtx := v.ctx
	newCtx, _ := s.WithContext(v.ctx)
	v.ctx = newCtx
	defer func() {
		v.ctx = oldCtx
	}()

	if err := v.sensors.Do(v.ctx); err != nil {
		return err
	}

	for id, c := range s.Readings {
		v.ctx = context.WithValue(newCtx, "ReadingId", id)
		if err := c.Accept(v); err != nil {
			return err
		}
	}

	return nil
}

func (v *visitor) VisitReading(s *Reading) error {
	oldCtx := v.ctx
	v.ctx, _ = s.WithContext(v.ctx)
	defer func() {
		v.ctx = oldCtx
	}()

	return v.reading.Do(v.ctx)
}

type VisitorBuilder interface {
	Stations(t task.Task) VisitorBuilder
	Station(t task.Task) VisitorBuilder
	Sensors(t task.Task) VisitorBuilder
	Reading(t task.Task) VisitorBuilder
	WithContext(context.Context) Visitor
}

type visitorBuilder struct {
	visitorCommon
}

func NewVisitor() VisitorBuilder {
	return &visitorBuilder{}
}

func (b *visitorBuilder) WithContext(ctx context.Context) Visitor {
	v := &visitor{visitorCommon: b.visitorCommon}
	v.ctx = context.WithValue(ctx, "Visitor", v)
	return v
}

func FromContext(ctx context.Context) Visitor {
	return ctx.Value("Visitor").(Visitor)
}

func (b *visitorBuilder) Stations(t task.Task) VisitorBuilder {
	b.stations = b.stations.Then(t)
	return b
}

func (b *visitorBuilder) Station(t task.Task) VisitorBuilder {
	b.station = b.station.Then(t)
	return b
}

func (b *visitorBuilder) Sensors(t task.Task) VisitorBuilder {
	b.sensors = b.sensors.Then(t)
	return b
}

func (b *visitorBuilder) Reading(t task.Task) VisitorBuilder {
	b.reading = b.reading.Then(t)
	return b
}
