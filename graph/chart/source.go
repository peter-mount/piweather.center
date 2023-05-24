package chart

import "github.com/peter-mount/piweather.center/util"

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
	DataSource() util.DataSource
	// Styles for use in SVG
	Styles() []string
}

type basicSource struct {
	name       string
	title      string
	subTitle   string
	dataSource util.DataSource
	styles     []string
}

// NewSource creates a Source
func NewSource(name, title, subTitle string, dataSource util.DataSource, styles ...string) Source {
	return &basicSource{
		name:       name,
		title:      title,
		subTitle:   subTitle,
		dataSource: dataSource,
		styles:     styles,
	}
}

// NewUnitSource creates a Source using the DataSource's Unit's Name and Unit as the Title and SubTitle.
func NewUnitSource(name string, dataSource util.DataSource, styles ...string) Source {
	u := dataSource.GetUnit()
	return NewSource(name, u.Name(), u.Unit(), dataSource, styles...)
}

func (s *basicSource) Name() string { return s.name }

func (s *basicSource) Title() string { return s.title }

func (s *basicSource) SubTitle() string { return s.subTitle }

func (s *basicSource) DataSource() util.DataSource { return s.dataSource }

func (s *basicSource) Styles() []string { return s.styles }
