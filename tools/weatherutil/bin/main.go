package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/weatherutil"
	"log"
)

func main() {
	err := kernel.Launch(
		&weatherutil.Rename{},
		&kernel.MemUsage{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
