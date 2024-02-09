package main

import (
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/go-script/stdlib"
	_ "github.com/peter-mount/go-script/stdlib/fmt"
	_ "github.com/peter-mount/go-script/stdlib/io"
	_ "github.com/peter-mount/go-script/stdlib/math"
	_ "github.com/peter-mount/go-script/stdlib/time"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/tools/weatheregress"
	"log"
)

func main() {
	err := kernel.Launch(
		&weatheregress.Egress{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
