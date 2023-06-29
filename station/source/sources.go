package source

import (
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
)

// Source defines the source of data for a Sensors collection.
// You must define one of these, otherwise no data will be received.
// You can define multiple entries here
type Source struct {
	WUnderground string      `yaml:"wunderground,omitempty"`
	EcoWitt      *EcoWitt    `yaml:"ecowitt,omitempty"`
	Amqp         *amqp.Queue `yaml:"amqp,omitempty"`
	Mqtt         *mqtt.Queue `yaml:"mqtt,omitempty"`
}

type EcoWitt struct {
	// Path under /api/ecowitt/
	Path string `json:"path" xml:"path,attr" yaml:"path"`
	// PassKey unique to the unit. "" to allow all
	PassKey string `json:"passKey,omitempty" xml:"passKey,attr,omitempty" yaml:"passKey,omitempty"`
}

type Amqp struct {
	// Url of the broker
	Url   string      `json:"url" xml:"url" yaml:"url"`
	Queue *amqp.Queue `json:"queue" xml:"queue" yaml:"queue"`
	// Exchange for publishing, defaults to amq.topic
	Exchange string `json:"exchange,omitempty" xml:"exchange,omitempty" yaml:"exchange,omitempty"`
	// Connection name that appears in the management plugin
	ConnectionName string `json:"connectionName,omitempty" xml:"connectionName,omitempty" yaml:"connectionName,omitempty"`
	// HeartBeat in seconds. Defaults to 10
	HeartBeat int `json:"heartBeat,omitempty" xml:"heartBeat,omitempty" yaml:"heartBeat,omitempty"`
	// Product name that appears in the management plugin (optional)
	Product string `json:"product,omitempty" xml:"product,omitempty" yaml:"product,omitempty"`
	// Version that appears in the management plugin (optional)
	Version string `json:"version,omitempty" xml:"version,omitempty" yaml:"version,omitempty"`
}
