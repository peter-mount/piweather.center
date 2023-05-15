package server

import "github.com/peter-mount/go-kernel/v2/rest"

// Server represents the primary service running the weather station.
type Server struct {
	Rest *rest.Server `kernel:"inject"`
}
