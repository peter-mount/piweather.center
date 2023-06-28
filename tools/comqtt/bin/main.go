package main

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/comqtt"
	"os"
)

func main() {
	if err := kernel.Launch(&comqtt.CoMQTT{}); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
