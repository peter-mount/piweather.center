package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type Http struct {
	Pos              lexer.Position
	Method           string               `parser:"'http' @('get'|'post') '('"`                // Http method to accept
	Format           *HttpFormat          `parser:"@@?"`                                       // Format of the data, default is json
	Timestamp        *string              `parser:"('timestamp' '(' ('now' | @String)? ')')?"` // Timestamp source parameter, time.Now() if not defined
	SourceParameters *SourceParameterList `parser:"@@ ')'"`                                    // Parameters to read from source
}

func (c *visitor[T]) Http(d *Http) error {
	var err error
	if d != nil {
		if c.http != nil {
			err = c.http(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.HttpFormat(d.Format)
		}

		if err == nil {
			err = c.SourceParameterList(d.SourceParameters)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Http(f func(Visitor[T], *Http) error) Builder[T] {
	b.http = f
	return b
}
