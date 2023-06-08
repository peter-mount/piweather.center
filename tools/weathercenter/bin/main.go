package main

import (
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/server/api"
	_ "github.com/peter-mount/piweather.center/server/api/graph"
	"github.com/peter-mount/piweather.center/server/ingress"
	"github.com/peter-mount/piweather.center/tools/weathercenter"
	"log"
)

func main() {
	err := kernel.Launch(
		&api.Api{},
		&ingress.Ingress{},
		&weathercenter.Server{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
