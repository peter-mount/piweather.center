package weatherscript

import (
	"github.com/peter-mount/go-script/packages"
	_ "github.com/peter-mount/go-script/stdlib"
	_ "github.com/peter-mount/go-script/stdlib/fmt"
	_ "github.com/peter-mount/go-script/stdlib/io"
	_ "github.com/peter-mount/go-script/stdlib/math"
	_ "github.com/peter-mount/go-script/stdlib/time"
	"github.com/peter-mount/piweather.center/astro/calculator"
	_ "github.com/peter-mount/piweather.center/script/astro"
	_ "github.com/peter-mount/piweather.center/script/geo"
	_ "github.com/peter-mount/piweather.center/script/value"
)

type Script struct {
	Calculator calculator.Calculator `kernel:"inject"`
}

func (s *Script) PostInit() error {
	// These are actual services which we will expose to go-script
	packages.Register("calculator", s.Calculator)
	return nil
}
