package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
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

func (v *visitor[T]) TableSelect(b *TableSelect) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v.tableSelect(v, b)
			if errors.IsVisitorStop(err) || errors.IsVisitorExit(err) {
				return nil
			}
		}

		if err == nil {
			err = v.Expression(b.Time)
		}
		if err == nil {
			err = v.Expression(b.Metric)
		}
		if err == nil {
			err = v.Unit(b.Unit)
		}
	}
	return err
}

func initTableSelect(v Visitor[*parserState], t *TableSelect) error {
	var err error
	if t.Unit != nil {
		err = v.Unit(t.Unit)
	}
	return err
}

func (b *builder[T]) TableSelect(f func(Visitor[T], *TableSelect) error) Builder[T] {
	b.common.tableSelect = f
	return b
}
