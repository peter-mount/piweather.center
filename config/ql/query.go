package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type Query struct {
	Pos         lexer.Position
	TimeZone    *time.TimeZone    `parser:"@@?"`
	QueryRange  *QueryRange       `parser:"@@"`
	Using       *UsingDefinitions `parser:"(@@)?"`
	Histogram   []*Histogram      `parser:"( ( @@ )+"`
	WindRose    []*WindRose       `parser:"| ( @@ )+"`
	TableSelect *TableSelect      `parser:"| @@ "`
	Limit       int               `parser:"| ( 'limit' @Number )?"`
	Select      []*Select         `parser:"  ( @@ )+ )"`
}

func (v *visitor[T]) Query(b *Query) error {
	var err error
	if b != nil {
		if v.query != nil {
			err = v.query(v, b)
		}
		if errors.IsVisitorStop(err) || errors.IsVisitorExit(err) {
			return nil
		}

		if err == nil {
			err = v.TimeZone(b.TimeZone)
		}

		if err == nil {
			err = v.QueryRange(b.QueryRange)
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
