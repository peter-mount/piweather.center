package graph

// Line graph of one or more properties
type Line struct {
	Common `yaml:",inline"`
	// YTitle of Y axis, default "" means none
	YTitle string `yaml:"YTitle,omitempty"`
	// YSubTitle subtitle of Y axis. default "" means the unit of the first reading.
	// If the value matches a registered Unit then it's Unit will be used.
	// Otherwise, text used on Y axis but the unit as the default will be used for the chart.
	YSubTitle string `yaml:"subtitle,omitempty"`
	// XStep if defined sets the number of minutes for the major markings on the x-axis
	XStep *float64 `yaml:"xstep,omitempty"`
	// XSubStep if defined sets the fraction of XStep for minor markings on the x-axis
	// If not set then the minor markings are not plotted
	XSubStep *float64 `yaml:"xsubstep,omitempty"`
	// YStep if defined sets the number of minutes for the major markins on the y-axis.
	YStep *float64 `yaml:"ystep,omitempty"`
	// YSubStep if defined sets the fraction of YStep for the minor markings on the y-axis.
	// If not set then the minor markings are not plotted
	YSubStep *float64 `yaml:"ysubstep,omitempty"`
	// ZeroY if true will ensure that 0 is present on the Y axis.
	// e.g. if true and Min or the calculated min value is > 0 then 0 will be the minimum.
	// if Max or the calculated max < 0 then 0 will be the maximum.
	ZeroY bool `yaml:"zeroY,omitempty"`
	// UTC if true then the chart will be plotted in UTC. false in the Local timezone
	UTC bool `yaml:"utc,omitempty"`
	// Readings is a list of readings to plot.
	// If this Line is linked to a Reading then it will be ignored and the reading used.
	Readings []string `yaml:"readings,omitempty"`
}
