package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type LocationExpression struct {
	Pos       lexer.Position
	Latitude  bool `parser:"( @'latitude'"`
	Longitude bool `parser:"| @'longitude'"`
	Altitude  bool `parser:"| @'altitude' )"`
}

func (c *visitor[T]) LocationExpression(d *LocationExpression) error {
	var err error
	if d != nil && c.locationExpression != nil {
		err = c.locationExpression(c, d)
		if util.IsVisitorStop(err) {
			return nil
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) LocationExpression(f func(Visitor[T], *LocationExpression) error) Builder[T] {
	b.locationExpression = f
	return b
}
