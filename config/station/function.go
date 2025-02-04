package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

// Function handles function calls
type Function struct {
	Pos         lexer.Position
	Name        string        `parser:"@Ident"`
	Expressions []*Expression `parser:"'(' (@@ (',' @@)*)? ')'"`
}

func (c *visitor[T]) Function(d *Function) error {
	var err error
	if d != nil {
		if c.function != nil {
			err = c.function(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Expressions {
				err = c.Expression(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (s *initState) initFunction(_ Visitor[*initState], l *Function) error {
	l.Name = strings.ToLower(l.Name)

	if !value.CalculatorExists(l.Name) {
		return participle.Errorf(l.Pos, "function %q is undefined", l.Name)
	}

	return nil
}

func (b *builder[T]) Function(f func(Visitor[T], *Function) error) Builder[T] {
	b.function = f
	return b
}

func printFunction(v Visitor[*printState], d *Function) error {
	var err error

	st := v.Get().
		Append("%s(", d.Name)

	for i, e := range d.Expressions {
		if i > 0 {
			st.Append(",")
		}

		err = v.Expression(e)
		if err != nil {
			break
		}
	}
	st.Append(")")

	if err == nil {
		err = errors.VisitorStop
	}
	return err
}
