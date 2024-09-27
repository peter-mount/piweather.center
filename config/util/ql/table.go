package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/units"
)

// TableSelect allows for a single metric to be returned in multiple rows
// based on a time range, e.g. all values for a day returned in a single row.
// If the select range spans multiple days then multiple days are returned.
//
//	between "2024-09-23" and "2024-09-27" every "1h"
//	declare "hour" AS between "row" truncate "1h" and "row" add "every"
//	table select
//	timeof(last(home.ecowitt.hrain_piezo)),
//	max(home.ecowitt.hrain_piezo)
type TableSelect struct {
	Pos lexer.Position

	Time   *Expression `parser:"'table' 'select' @@"`
	Metric *Expression `parser:"',' @@"`
	Unit   *units.Unit `parser:"( 'unit' @@ )?"`
}
