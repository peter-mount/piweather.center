package main

import (
	"github.com/peter-mount/go-kernel"
	"github.com/peter-mount/piweather.center/tools/mastodon"
	"log"
)

func main() {
	err := kernel.Launch(
		&mastodon.Mastodon{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
