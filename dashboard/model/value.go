package model

import "github.com/peter-mount/piweather.center/dashboard/registry"

func init() {
	f := func() registry.Component { return &Value{} }
	registry.Register("value", f)
	registry.Register("rain-gauge", f)
}

type Value struct {
	Component `yaml:",inline"`
	Label     string   `yaml:"label"`
	Metric    string   `yaml:"metric"`
	Min       *float64 `yaml:"min,omitempty"`
	Max       *float64 `yaml:"max,omitempty"`
}
