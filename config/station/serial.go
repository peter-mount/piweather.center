package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type Serial struct {
	Pos      lexer.Position
	Driver   string `parser:"'serial' '(' @String"` // device driver id
	Port     string `parser:" @String"`             // serial port
	BaudRate int    `parser:" @Number ')'"`         // Baud rate
	//DataBits int    `parser:"('data' @('5'|'6'|'7'|'8'))?"`
	//Parity   string `parser:"('parity' @('no'|'none'|'odd'|'even'))?"`
	//StopBits string `parser:"('stop' @('1'|'1.5'|'2'))?"`
}

func (c *visitor[T]) Serial(d *Serial) error {
	var err error
	if d != nil {
		if c.serial != nil {
			err = c.serial(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Serial(f func(Visitor[T], *Serial) error) Builder[T] {
	b.serial = f
	return b
}
