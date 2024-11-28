package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type ComponentList struct {
	Pos     lexer.Position
	Entries []*ComponentListEntry `parser:"@@*"`
}

type ComponentListEntry struct {
	Pos        lexer.Position
	Container  *Container  `parser:"( @@"`
	Gauge      *Gauge      `parser:"| @@"`
	MultiValue *MultiValue `parser:"| @@"`
	Value      *Value      `parser:"| @@ )"`
}

type Container struct {
	Pos        lexer.Position
	Type       string         `parser:"@('container' | 'col' | 'row') '('"`
	Component  *Component     `parser:"@@"`
	Components *ComponentList `parser:"@@ ')'"`
}

type Dashboard struct {
	Pos        lexer.Position
	Name       string              `parser:"'dashboard' '(' @String?"`
	Live       bool                `parser:"@'live'?"`
	Update     *time.CronTab       `parser:"('update' @@)?"`
	Component  *Component          `parser:"@@"`
	Components *ComponentListEntry `parser:"@@? ')'"`
}
