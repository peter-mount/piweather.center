package chart

import (
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

// Source represents a collection of values provided by a DataSource.
// It also provides metadata specific for the Chart being generated.
type Source interface {
	// Name of this Source
	Name() string
	// Title or "" if unused, used in Axes
	Title() string
	// SubTitle or "" if unused, used in Axes. e.g. value.Unit string
	SubTitle() string
	// DataSource containing the data for this Source
	DataSource() DataSource
	// Styles for use in SVG
	Styles() []string
}

// DataSource represents a collection of values to be plotted
type DataSource interface {
	// Size of the DataSource
	Size() int
	// Get a specific entry in the DataSource
	Get(int) (time.Time, value.Value)
	// Period returns the Period of the entries within the DataSource
	Period() time2.Period
	// GetYRange returns the Range of values in the DataSource
	GetYRange() *value.Range
	// GetUnit returns the Unit of the values in the DataSource
	GetUnit() *value.Unit
	// ForEach calls a function for each entry in the DataSource
	ForEach(func(int, time.Time, value.Value))
}

type basicSource struct {
	name       string
	title      string
	subTitle   string
	dataSource DataSource
	styles     []string
}

// NewSource creates a Source
func NewSource(name, title, subTitle string, dataSource DataSource, styles ...string) Source {
	return &basicSource{
		name:       name,
		title:      title,
		subTitle:   subTitle,
		dataSource: dataSource,
		styles:     styles,
	}
}

// NewUnitSource creates a Source using the DataSource's Unit's Name and Unit as the Title and SubTitle.
func NewUnitSource(name string, dataSource DataSource, styles ...string) Source {
	u := dataSource.GetUnit()
	return NewSource(name, u.Name(), u.Unit(), dataSource, styles...)
}

func (s *basicSource) Name() string { return s.name }

func (s *basicSource) Title() string { return s.title }

func (s *basicSource) SubTitle() string { return s.subTitle }

func (s *basicSource) DataSource() DataSource { return s.dataSource }

func (s *basicSource) Styles() []string { return s.styles }
