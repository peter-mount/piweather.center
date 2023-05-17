package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/server"
	"log"
)

func main() {
	err := kernel.Launch(
		&server.Server{},
		//&ecowitt.Server{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
