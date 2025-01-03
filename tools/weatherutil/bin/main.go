package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/weatherutil/query"
	"github.com/peter-mount/piweather.center/tools/weatherutil/rename"
	"github.com/peter-mount/piweather.center/tools/weatherutil/statistics"
	"log"
)

func main() {
	err := kernel.Launch(
		&query.Query{},
		&rename.Rename{},
		&statistics.Stats{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
