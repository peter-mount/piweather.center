package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/server/api"
	_ "github.com/peter-mount/piweather.center/server/api/graph"
	"github.com/peter-mount/piweather.center/server/ingress"
	"github.com/peter-mount/piweather.center/tools/weathercenter"
	"os"
)

func main() {
	if err := kernel.Launch(
		&api.Api{},
		&ingress.Ingress{},
		&weathercenter.Server{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
