package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/store/file/server"
	"github.com/peter-mount/piweather.center/store/util"
	_ "github.com/peter-mount/piweather.center/weather/measurement"
	"os"
)

func main() {
	if err := kernel.Launch(
		&server.Server{},
		&util.Importer{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
