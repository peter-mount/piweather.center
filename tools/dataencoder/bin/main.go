package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/dataencoder"
	"log"
)

func main() {
	err := kernel.Launch(
		&dataencoder.Build{},
		&dataencoder.Vsop87Encoder{},
		&dataencoder.YbscEncoder{},
		&dataencoder.WebEncoder{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
