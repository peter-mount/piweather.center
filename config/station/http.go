package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Http struct {
	Pos              lexer.Position
	Method           string               `parser:"'http' @('get'|'post'|'put'|'patch') '('"` // Http method to accept
	Format           *HttpFormat          `parser:"@@?"`                                      // Format of the data, default is json
	Timestamp        *SourcePath          `parser:"('timestamp' '(' @@ ')')?"`                // Timestamp source parameter, time.Now() if not defined
	SourceParameters *SourceParameterList `parser:"@@ ')'"`                                   // Parameters to read from source
}

func (c *visitor[T]) Http(d *Http) error {
	var err error
	if d != nil {
		if c.http != nil {
			err = c.http(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.HttpFormat(d.Format)
		}

		if err == nil {
			err = c.SourcePath(d.Timestamp)
		}

		if err == nil {
			err = c.SourceParameterList(d.SourceParameters)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initHttp(v Visitor[*initState], _ *Http) error {
	s := v.Get()
	s.sensorParameters = make(map[string]*SourceParameter)
	s.sourcePath = nil
	return nil
}

func (b *builder[T]) Http(f func(Visitor[T], *Http) error) Builder[T] {
	b.http = f
	return b
}
