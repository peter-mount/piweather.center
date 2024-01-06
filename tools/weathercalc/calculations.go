package weathercalc

import (
	"flag"
	"github.com/peter-mount/piweather.center/tools/weathercalc/lang"
)

type Calculations struct {
	script *lang.Script
}

func (s *Calculations) Start() error {
	p := lang.NewParser()
	script, err := p.ParseFiles(flag.Args()...)
	if err != nil {
		return err
	}

	s.script = script
	return nil
}

func (s *Calculations) Script() *lang.Script {
	return s.script
}
