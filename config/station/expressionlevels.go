package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type ExpressionLevel1 struct {
	Pos lexer.Position

	Left  *ExpressionLevel2 `parser:"@@"`
	Op    string            `parser:"[ @( '|' '|' )"`
	Right *ExpressionLevel1 `parser:"  @@ ]"`
}

func (c *visitor[T]) ExpressionLevel1(b *ExpressionLevel1) error {
	var err error
	if b != nil {
		if c.expressionLevel1 != nil {
			err = c.expressionLevel1(c, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.ExpressionLevel2(b.Left)
		}

		if err == nil {
			err = c.ExpressionLevel1(b.Right)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) ExpressionLevel1(f func(Visitor[T], *ExpressionLevel1) error) Builder[T] {
	b.expressionLevel1 = f
	return b
}

type ExpressionLevel2 struct {
	Pos lexer.Position

	Left  *ExpressionLevel3 `parser:"@@"`
	Op    string            `parser:"[ @( '&' '&' )"`
	Right *ExpressionLevel2 `parser:"  @@ ]"`
}

func (c *visitor[T]) ExpressionLevel2(b *ExpressionLevel2) error {
	var err error
	if b != nil {
		if c.expressionLevel2 != nil {
			err = c.expressionLevel2(c, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.ExpressionLevel3(b.Left)
		}

		if err == nil {
			err = c.ExpressionLevel2(b.Right)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) ExpressionLevel2(f func(Visitor[T], *ExpressionLevel2) error) Builder[T] {
	b.expressionLevel2 = f
	return b
}

type ExpressionLevel3 struct {
	Pos   lexer.Position
	Left  *ExpressionLevel4 `parser:"@@"`
	Op    string            `parser:"[ @( '=' '=' | '!' '=' | '<' '=' | '<' | '>' '=' | '>' )"`
	Right *ExpressionLevel3 `parser:"  @@ ]?"`
}

func (c *visitor[T]) ExpressionLevel3(b *ExpressionLevel3) error {
	var err error
	if b != nil {
		if c.expressionLevel3 != nil {
			err = c.expressionLevel3(c, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.ExpressionLevel4(b.Left)
		}

		if err == nil {
			err = c.ExpressionLevel3(b.Right)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) ExpressionLevel3(f func(Visitor[T], *ExpressionLevel3) error) Builder[T] {
	b.expressionLevel3 = f
	return b
}

type ExpressionLevel4 struct {
	Pos lexer.Position

	Op    string            `parser:"  ( @( '!' | '-' )"`
	Left  *ExpressionLevel5 `parser:"    @@ )"`
	Right *ExpressionLevel5 `parser:"| @@"`
}

func (c *visitor[T]) ExpressionLevel4(b *ExpressionLevel4) error {
	var err error
	if b != nil {
		if c.expressionLevel4 != nil {
			err = c.expressionLevel4(c, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.ExpressionLevel5(b.Left)
		}

		if err == nil {
			err = c.ExpressionLevel5(b.Right)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) ExpressionLevel4(f func(Visitor[T], *ExpressionLevel4) error) Builder[T] {
	b.expressionLevel4 = f
	return b
}

type ExpressionLevel5 struct {
	Pos lexer.Position

	Float         *float64        `parser:"( @Number"`
	True          bool            `parser:"| @'true'"`
	False         bool            `parser:"| @'false'"`
	SubExpression *Expression     `parser:"| '(' @@ ')' "`
	Atom          *ExpressionAtom `parser:"| @@ )"`
}

func (c *visitor[T]) ExpressionLevel5(b *ExpressionLevel5) error {
	var err error
	if b != nil {
		if c.expressionLevel5 != nil {
			err = c.expressionLevel5(c, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Expression(b.SubExpression)
		}

		if err == nil {
			err = c.ExpressionAtom(b.Atom)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) ExpressionLevel5(f func(Visitor[T], *ExpressionLevel5) error) Builder[T] {
	b.expressionLevel5 = f
	return b
}

func printExpressionLevelX[L any, R any](v Visitor[*printState], l L, op string, r R, lf func(L) error, rf func(R) error) error {
	err := lf(l)
	if err == nil && op != "" {
		v.Get().Append(" %s ", op)
		err = rf(r)
	}

	if err == nil {
		err = errors.VisitorStop
	}
	return err
}

func printExpressionLevel1(v Visitor[*printState], d *ExpressionLevel1) error {
	return printExpressionLevelX(v, d.Left, d.Op, d.Right, v.ExpressionLevel2, v.ExpressionLevel1)
}

func printExpressionLevel2(v Visitor[*printState], d *ExpressionLevel2) error {
	return printExpressionLevelX(v, d.Left, d.Op, d.Right, v.ExpressionLevel3, v.ExpressionLevel2)
}

func printExpressionLevel3(v Visitor[*printState], d *ExpressionLevel3) error {
	return printExpressionLevelX(v, d.Left, d.Op, d.Right, v.ExpressionLevel4, v.ExpressionLevel3)
}

func printExpressionLevel4(v Visitor[*printState], d *ExpressionLevel4) error {
	var err error
	if d.Left != nil {
		v.Get().Append(" %s ", d.Op)
		err = v.ExpressionLevel5(d.Left)
	} else {
		err = v.ExpressionLevel5(d.Right)
	}

	if err == nil {
		err = errors.VisitorStop
	}
	return err
}

func printExpressionLevel5(v Visitor[*printState], d *ExpressionLevel5) error {
	st := v.Get()
	var err error
	switch {
	case d.Float != nil:
		st.Append("%.3f", *d.Float)
	case d.True:
		st.Append("true")
	case d.False:
		st.Append("false")
	case d.SubExpression != nil:
		st.Append("( ")
		err = v.Expression(d.SubExpression)
		st.Append(" )")

	case d.Atom != nil:
		err = v.ExpressionAtom(d.Atom)
	}

	if err == nil {
		err = errors.VisitorStop
	}
	return err
}
