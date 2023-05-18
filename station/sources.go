package station

import (
	"github.com/peter-mount/piweather.center/util/mq"
)

// Source defines the source of data for a Sensors collection.
// You must define one of these, otherwise no data will be received.
// You can define multiple entries here
type Source struct {
	// WUnderground protocol endpoint
	WUnderground string `json:"wunderground,omitempty" xml:"wunderground,attr,omitempty" yaml:"wunderground,omitempty"`
	// EcoWitt protocol endpoint under /api/import
	EcoWitt *EcoWitt `json:"ecowitt,omitempty" xml:"ecowitt,attr,omitempty" yaml:"ecowitt,omitempty"`
	// Amqp RabbitMQ broker
	Amqp *mq.Queue `json:"amqp,omitempty" xml:"amqp,omitempty" yaml:"amqp,omitempty"`
	// TODO add MQTT broker here
}

type EcoWitt struct {
	// Path under /api/ecowitt/
	Path string `json:"path" xml:"path,attr" yaml:"path"`
	// PassKey unique to the unit. "" to allow all
	PassKey string `json:"passKey,omitempty" xml:"passKey,attr,omitempty" yaml:"passKey,omitempty"`
}

type Amqp struct {
	// Url of the broker
	Url   string    `json:"url" xml:"url" yaml:"url"`
	Queue *mq.Queue `json:"queue" xml:"queue" yaml:"queue"`
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
