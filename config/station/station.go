package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/location"
)

type Stations struct {
	Pos      lexer.Position
	Stations []*Station `parser:"@@+"`
}

type Station struct {
	Pos        lexer.Position
	Name       string             `parser:"'station' '(' @String"`
	Location   *location.Location `parser:"@@?"`
	Dashboards *DashboardList     `parser:"@@ ')'"`
}

type DashboardList struct {
	Pos        lexer.Position
	Dashboards []*Dashboard `parser:"@@*"`
}
