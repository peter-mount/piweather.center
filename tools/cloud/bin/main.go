package main

import (
	"github.com/peter-mount/go-kernel"
	"github.com/peter-mount/piweather.center/tools/cloud"
	"log"
)

func main() {
	err := kernel.Launch(
		&cloud.Service{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
