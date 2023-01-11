package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/videogenerator"
	"log"
)

func main() {
	err := kernel.Launch(
		&videogenerator.VideoGenerator{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
