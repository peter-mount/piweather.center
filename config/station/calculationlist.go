package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type CalculationList struct {
	Pos          lexer.Position
	Calculations []*Calculation `parser:"(@@)*"`
}

func (c *visitor[T]) CalculationList(d *CalculationList) error {
	var err error
	if d != nil {
		if c.calculationList != nil {
			err = c.calculationList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		for _, e := range d.Calculations {
			err = c.Calculation(e)
			if err != nil {
				break
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initCalculationList(v Visitor[*initState], d *CalculationList) error {
	s := v.Get()
	s.calculations = make(map[string]*Calculation)

	// sensorPrefix is not used for calculations
	s.sensorPrefix = ""
	if s.stationPrefix == "" {
		// should never occur
		return errors.Errorf(d.Pos, "stationPrefix not defined")
	}
	return nil
}

func (b *builder[T]) CalculationList(f func(Visitor[T], *CalculationList) error) Builder[T] {
	b.calculationList = f
	return b
}
