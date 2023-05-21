package chart

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Source interface {
	Name() string
	DataSource() DataSource
	Styles() []string
}

type basicSource struct {
	name       string
	dataSource DataSource
	styles     []string
}

type DataSource interface {
	Size() int
	Get(int) (time.Time, value.Value)
	GetXRange() (time.Time, time.Time)
	GetYRange() (value.Value, value.Value)
	GetUnit() *value.Unit
	ForEach(func(int, time.Time, value.Value))
}

func NewSource(name string, dataSource DataSource, styles ...string) Source {
	return &basicSource{
		name:       name,
		dataSource: dataSource,
		styles:     styles,
	}
}

func (s *basicSource) Name() string { return s.name }

func (s *basicSource) DataSource() DataSource { return s.dataSource }

func (s *basicSource) Styles() []string { return s.styles }
