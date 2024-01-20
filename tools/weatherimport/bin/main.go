package main

import (
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/tools/weatherimport"
	"log"
)

func main() {
	err := kernel.Launch(
		&weatherimport.Importer{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
