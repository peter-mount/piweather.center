package homeassistant

import (
	mq "github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/rabbitmq/amqp091-go"
)

type HomeAssistant struct {
	// if set then HomeAssistant is disabled - used in development
	Disabled bool `yaml:"disabled"`

	// Amqp broker to use, "" for none
	Amqp           string        `yaml:"amqp,omitempty"`
	AmqpPublisher  *mq.Publisher `yaml:"amqp_publisher"`
	amqp           *mq.MQ
	amqpConnection *amqp091.Connection

	// DiscoveryPrefix, will default to "homeassistant"
	DiscoveryPrefix string `yaml:"discovery_prefix,omitempty"`
	// Sensors sensor definitions
	Sensors []Sensor `yaml:"sensors"`
}

func (h *HomeAssistant) close() {
	if h != nil {
		if h.amqpConnection != nil {
			_ = h.amqpConnection.Close()
		}

		h.amqp = nil
		h.amqpConnection = nil
	}
}

// Sensor defines a sensor within HomeAssistant.
// This is similar to Sensors within the weather station in that it's a collection of
// Sensors (be it Reading, CalculatedValue etc.) that will appear together within HA.
type Sensor struct {
	// NodeId to use to keep a Sensor's Entity's together
	NodeId string `yaml:"nodeId"`
	// ObjectIdPrefix is a unique prefix which will be used with the Entity to generate a
	// unique object_id that will be sent to HomeAssistant
	ObjectIdPrefix string `yaml:"object_id_prefix"`
	// Entities contained in this Sensor.
	// The key will be appended to ObjectIdPrefix to form the object_id
	Entities map[string]*Entity `yaml:"entities"`
}

type Entity struct {
	// ObjectID for this entity. This is set when the configuration is loaded.
	//ObjectID string `yaml:"-" json:"object_id"`
	// SensorId linking to the value recorded.
	// This is only used by yaml, json ignores this as we don't want it sent to HomeAssistant.
	SensorId   string `yaml:"sensor_id" json:"-"`
	SensorType string `yaml:"sensor_type" json:"-"`
	ObjectId   string `yaml:"object_id,omitempty" json:"object_id"`
	// Name of sensor (optional, if not set then will be set to the key name in the parent Sensors)
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// DeviceClass of device, default "None"
	DeviceClass string `yaml:"device_class,omitempty" json:"device_class,omitempty"`
	// EnabledByDefault flag which defines if the entity should be enabled when first added, default true
	EnabledByDefault bool `yaml:"enabled_by_default,omitempty" json:"enabled_by_default,omitempty"`
	// Encoding The encoding of the payloads received and published messages.
	// Set to "" to disable decoding of incoming payload.
	Encoding *string `yaml:"encoding,omitempty" json:"encoding,omitempty"`
	//The category of the entity.
	EntityCategory string `yaml:"entity_category,omitempty" json:"entity_category,omitempty"`
	// Icon for the entity (optional)
	Icon string `yaml:"icon,omitempty" json:"icon,omitempty"`
	// Min minimum value (Number default 1)
	Min *float64 `yaml:"min,omitempty" json:"min,omitempty"`
	// Max maximum value (Number default 100)
	Max *float64 `yaml:"max,omitempty" json:"max,omitempty"`
	// Mode optional default "auto"
	Mode string `yaml:"mode,omitempty" json:"mode,omitempty"`
	// StateTopic MQTT topic subscribed to receive values (Number)
	StateTopic string `yaml:"state_topic" json:"state_topic,omitempty"`
	// Step value, smallest value 0.001, default 1
	Step *float64 `yaml:"step,omitempty" json:"step,omitempty"`
	// UniqueID An ID that uniquely identifies this Number.
	// If two Numbers have the same unique ID Home Assistant will raise an exception.
	UniqueID string `yaml:"unique_id,omitempty" json:"unique_id,omitempty"`
	// UnitOfMeasurement Defines the unit of measurement of the sensor, if any. The unit_of_measurement can be null.
	UnitOfMeasurement string `yaml:"unit_of_measurement,omitempty" json:"unit_of_measurement,omitempty"`
	//	ValueTemplate	Defines a template to extract the value.
	ValueTemplate             string `yaml:"value_template,omitempty" json:"value_template,omitempty"`
	SuggestedDisplayPrecision *int   `yaml:"suggested_display_precision,omitempty" json:"suggested_display_precision,omitempty"`
}
