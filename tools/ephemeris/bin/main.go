package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/ephemeris"
	"log"
)

func main() {
	err := kernel.Launch(
		&ephemeris.Ephemeris{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
