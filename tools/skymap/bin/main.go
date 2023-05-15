package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/skymap"
	"log"
)

func main() {
	err := kernel.Launch(
		&skymap.Skymap{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
