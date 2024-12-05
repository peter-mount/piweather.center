package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type WindRose struct {
	Pos lexer.Position

	Degrees *Expression      `parser:"'windrose' @@"`
	Speed   *Expression      `parser:"',' @@"`
	Options []WindRoseOption `parser:"('as' @@ (',' @@)* )?"`
}

type WindRoseOption struct {
	Pos   lexer.Position
	Rose  bool `parser:"( @'rose'"`
	Count bool `parser:"| @'count'"`
	Max   bool `parser:"| @'max')"`
}

func (v *visitor) WindRose(b *WindRose) error {
	var err error
	if b != nil {
		if v.windRose != nil {
			err = v.windRose(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}
		if err == nil {
			err = v.Expression(b.Degrees)
		}
		if err == nil {
			err = v.Expression(b.Speed)
		}
	}
	return err
}

func (b *builder) WindRose(f func(Visitor, *WindRose) error) Builder {
	b.common.windRose = f
	return b
}
