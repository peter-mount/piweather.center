package station

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Component struct {
	Pos lexer.Position
	//Type      string     `yaml:"type"`            // type of component - required
	Title     string     `parser:"('title' @String)?"` // title, optional based on component
	Class     string     `parser:"('class' @String)?"` // optional CSS class
	Style     string     `parser:"('style' @String)?"` // optional inline CSS
	ID        string     // Unique ID, generated on load
	dashboard *Dashboard // link to dashboard
}

type Value struct {
	Pos       lexer.Position
	Type      string      `parser:"@('gauge' | 'value') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Metrics   *MetricList `parser:"@@ ')'"`
}

type MultiValue struct {
	Pos       lexer.Position
	Component *Component     `parser:"'multiValue' '(' @@"`
	Pattern   *MetricPattern `parser:"@@"`
	Time      bool           `parser:"@'time'? ')'"`
}
