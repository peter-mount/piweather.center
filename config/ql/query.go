package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type Query struct {
	Pos         lexer.Position
	QueryRange  *QueryRange       `parser:"@@"`
	Using       *UsingDefinitions `parser:"(@@)?"`
	Histogram   []*Histogram      `parser:"( ( @@ )+"`
	WindRose    []*WindRose       `parser:"| ( @@ )+"`
	TableSelect *TableSelect      `parser:"| @@ "`
	Limit       int               `parser:"| ( 'limit' @Int )?"`
	Select      []*Select         `parser:"  ( @@ )+ )"`
}

func (v *visitor[T]) Query(b *Query) error {
	var err error
	if b != nil {
		// Process QueryRange first
		err = v.QueryRange(b.QueryRange)

		if err == nil && v.query != nil {
			err = v.query(v, b)
		}
		if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
			return nil
		}

		if err == nil {
			err = v.UsingDefinitions(b.Using)
		}

		if err == nil {
			for _, sel := range b.Select {
				if err == nil {
					err = v.Select(sel)
				}
			}
		}

		if err == nil {
			for _, sel := range b.WindRose {
				if err == nil {
					err = v.WindRose(sel)
				}
			}
		}

		if err == nil {
			err = v.TableSelect(b.TableSelect)
		}
	}
	return err
}

func initQuery(_ Visitor[*parserState], s *Query) error {
	return assertLimit(s.Pos, s.Limit)
}

func (b *builder[T]) Query(f func(Visitor[T], *Query) error) Builder[T] {
	b.common.query = f
	return b
}
