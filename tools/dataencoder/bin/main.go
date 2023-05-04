package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/dataencoder/ybsc"
	"log"
)

func main() {
	err := kernel.Launch(
		&ybsc.YbscEncoder{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
