package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	_ "github.com/peter-mount/piweather.center/homeassistant"
	"github.com/peter-mount/piweather.center/tools/weathercenter"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/view"
	"os"
)

func main() {
	if err := kernel.Launch(
		&weathercenter.Server{},
		&view.Service{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
