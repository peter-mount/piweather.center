package main

import (
	"github.com/peter-mount/go-kernel/v2"
	videointro "github.com/peter-mount/piweather.center/tools/videoident"
	"log"
)

func main() {
	err := kernel.Launch(
		&videointro.VideoIntro{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
