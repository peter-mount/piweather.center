package view

import "github.com/peter-mount/go-kernel/v2"

func init() {
	kernel.Register(&Home{}, &Units{})
}
