package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Rtl433 struct {
	Pos              lexer.Position
	Frequency        float64              `parser:"'rtl433' '(' ('freq' @Number)?"` // Frequency to receive on in MHz, default 433M
	Model            string               `parser:" 'model' @String"`               // Model to match against
	SubType          string               `parser:" ('subtype' @String)?"`          // Subtype, e.g. Wx080 has 3 types of messages
	Id               string               `parser:" 'id' @String"`                  // ID to match against
	SourceParameters *SourceParameterList `parser:"@@ ')'"`                         // Parameters to read from source
}

func (c *visitor[T]) Rtl433(d *Rtl433) error {
	var err error

	if d != nil {
		if c.rtl433 != nil {
			err = c.rtl433(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.SourceParameterList(d.SourceParameters)
		}

		err = errors.Error(d.Pos, err)
	}

	return err
}

func initRtl433(v Visitor[*initState], _ *Rtl433) error {
	s := v.Get()
	s.sensorParameters = make(map[string]*SourceParameter)
	s.sourcePath = nil
	return nil
}

func (b *builder[T]) Rtl433(f func(Visitor[T], *Rtl433) error) Builder[T] {
	b.rtl433 = f
	return b
}
