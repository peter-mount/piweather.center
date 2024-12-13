package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type I2C struct {
	Pos lexer.Position
	// smbus is a subset of i2c so it's an alias here
	Driver string `parser:"('i2c'|'smbus') '(' @String"` // device driver id
	Bus    int    `parser:"    @Number"`                 // i2c bus id in the OS kernel
	Device int    `parser:"':' @Number ')'"`             // i2c address on the specific bus
}

func (c *visitor[T]) I2C(d *I2C) error {
	var err error
	if d != nil {
		if c.i2c != nil {
			err = c.i2c(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) I2C(f func(Visitor[T], *I2C) error) Builder[T] {
	b.i2c = f
	return b
}
