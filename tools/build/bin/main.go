package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/build"
	"log"
)

func main() {
	err := kernel.Launch(
		&build.Vsop87Encoder{},
		&build.YbscEncoder{},
		&build.WebEncoder{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
