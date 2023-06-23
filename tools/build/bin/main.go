package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/build"
	"os"
)

func main() {
	if err := kernel.Launch(
		&build.Vsop87Encoder{},
		&build.YbscEncoder{},
		&build.WebEncoder{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
