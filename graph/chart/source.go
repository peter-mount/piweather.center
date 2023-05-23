package chart

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Source interface {
	Name() string
	Title() string
	SubTitle() string
	DataSource() DataSource
	Styles() []string
}

type DataSource interface {
	Size() int
	Get(int) (time.Time, value.Value)
	GetXRange() (time.Time, time.Time)
	GetYRange() (value.Value, value.Value)
	GetUnit() *value.Unit
	ForEach(func(int, time.Time, value.Value))
}

type basicSource struct {
	name       string
	title      string
	subTitle   string
	dataSource DataSource
	styles     []string
}

func NewSource(name, title, subTitle string, dataSource DataSource, styles ...string) Source {
	return &basicSource{
		name:       name,
		title:      title,
		subTitle:   subTitle,
		dataSource: dataSource,
		styles:     styles,
	}
}

func NewUnitSource(name string, dataSource DataSource, styles ...string) Source {
	u := dataSource.GetUnit()
	return NewSource(name, u.Name(), u.Unit(), dataSource, styles...)
}

func (s *basicSource) Name() string { return s.name }

func (s *basicSource) Title() string { return s.title }

func (s *basicSource) SubTitle() string { return s.subTitle }

func (s *basicSource) DataSource() DataSource { return s.dataSource }

func (s *basicSource) Styles() []string { return s.styles }
