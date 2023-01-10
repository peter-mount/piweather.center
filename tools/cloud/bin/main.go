package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/cloud"
	"log"
)

func main() {
	err := kernel.Launch(
		&cloud.App{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
