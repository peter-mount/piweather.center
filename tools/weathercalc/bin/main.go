package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/tools/weathercalc"
	_ "github.com/peter-mount/piweather.center/weather/forecast"
	_ "github.com/peter-mount/piweather.center/weather/measurement"
	"os"
)

func main() {
	if err := kernel.Launch(
		&weathercalc.Service{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
