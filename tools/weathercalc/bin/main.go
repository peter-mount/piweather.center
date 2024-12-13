package main

import (
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/tools/weathercalc"
	_ "github.com/peter-mount/piweather.center/weather/forecast"
	_ "github.com/peter-mount/piweather.center/weather/measurement"
	"log"
)

func main() {
	err := kernel.Launch(
		&weathercalc.Service{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
