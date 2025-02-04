package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/weatherutil/config"
	"github.com/peter-mount/piweather.center/tools/weatherutil/query"
	"github.com/peter-mount/piweather.center/tools/weatherutil/rename"
	"github.com/peter-mount/piweather.center/tools/weatherutil/statistics"
	"os"
)

func main() {
	err := kernel.Launch(
		&config.Config{},
		&query.Query{},
		&rename.Rename{},
		&statistics.Stats{},
	)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
