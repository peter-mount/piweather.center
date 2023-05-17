package station

import (
	"github.com/peter-mount/piweather.center/util/mq"
)

// Config is the root and consists of one or more Station's
type Config struct {
	// Stations one or more Weather Stations supported by this instance
	Stations map[string]Station `json:"stations" xml:"stations" yaml:"stations"`
}

// Station defines a Weather Station at a specific location.
// It consists of one or more Sensor's
type Station struct {
	// Name of the station
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Location of the station
	Location Location `json:"location" xml:"location,omitempty" yaml:"location,omitempty"`
	// One or more Sensors collection
	Sensors map[string]Sensors `json:"sensors" xml:"sensors" yaml:"sensors"`
}

// Sensors define a Sensor collection within the Station.
// A Sensor collection is
type Sensors struct {
	// Name of the Sensors collection
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Source of data for this collection
	Source Source `json:"source" xml:"source" yaml:"source"`
	// Sensor's provided by this collection
	Sensors map[string]Sensor `json:"sensors" xml:"sensors" yaml:"sensors"`
}

// Source defines the source of data for a Sensors collection.
// You must define one of these, otherwise no data will be received.
// You can define multiple entries here
type Source struct {
	// WUnderground protocol endpoint
	WUnderground string `json:"wunderground,omitempty" xml:"wunderground,attr,omitempty" yaml:"wunderground,omitempty"`
	// EcoWitt protocol endpoint
	EcoWitt string `json:"ecowitt,omitempty" xml:"ecowitt,attr,omitempty" yaml:"ecowitt,omitempty"`
	// Amqp RabbitMQ broker
	Amqp *mq.Queue `json:"amqp,omitempty" xml:"amqp,omitempty" yaml:"amqp,omitempty"`
	// TODO add MQTT broker here
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
