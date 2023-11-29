package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/dashboard/view"
	_ "github.com/peter-mount/piweather.center/homeassistant"
	_ "github.com/peter-mount/piweather.center/server/api/graph"
	"github.com/peter-mount/piweather.center/tools/weathercenter"
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
