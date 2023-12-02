package model

import "github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/registry"

func init() {
	registry.Register("text", func() registry.Component { return &Text{} })
}

type Text struct {
	Component `yaml:",inline"`
	Text      string `yaml:"text"`
}
