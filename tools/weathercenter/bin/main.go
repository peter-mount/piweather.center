package main

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/server"
	_ "github.com/peter-mount/piweather.center/server/api/graph"
	"log"
)

func main() {
	err := kernel.Launch(&server.Server{})
	if err != nil {
		log.Fatal(err)
	}
}
