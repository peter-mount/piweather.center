package model

import "github.com/peter-mount/piweather.center/dashboard/registry"

func init() {
	registry.Register("value", func() registry.Component { return &Value{} })
}

type Value struct {
	Component `yaml:",inline"`
	Label     string `yaml:"label"`
	Metric    string `yaml:"metric"`
}
