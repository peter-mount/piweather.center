package model

import (
	"github.com/peter-mount/piweather.center/util/mq/amqp"
	"github.com/peter-mount/piweather.center/util/mq/mqtt"
)

// Source defines the source of data for a Sensors collection.
// You must define one of these, otherwise no data will be received.
// You can define multiple entries here
type Source struct {
	WUnderground string      `yaml:"wunderground,omitempty"`
	Http         *Http       `yaml:"ecowitt,omitempty"`
	Amqp         *amqp.Queue `yaml:"amqp,omitempty"`
	Mqtt         *mqtt.Queue `yaml:"mqtt,omitempty"`
}

type Http struct {
	// Path under /api/http/
	Path string `json:"path" xml:"path,attr" yaml:"path"`
	// PassKey unique to the unit. "" to allow all
	PassKey string `json:"passKey,omitempty" xml:"passKey,attr,omitempty" yaml:"passKey,omitempty"`
}
