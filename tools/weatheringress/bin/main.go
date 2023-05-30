package main

import (
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	_ "github.com/peter-mount/piweather.center/server/api/graph"
	"github.com/peter-mount/piweather.center/server/ingress"
	"log"
)

func main() {
	err := kernel.Launch(
		&ingress.Ingress{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
