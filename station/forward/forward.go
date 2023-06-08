package forward

// Forward defines how the system forwards readings to a downstream system.
type Forward struct {
}

// Endpoint defines the remote destination for messages
type Endpoint struct {
	Amqp *Amqp     `json:"amqp,omitempty"`
	Post *HttpPost `json:"post,omitempty"`
}

// Amqp destination
type Amqp struct {
	// Broker id to use
	Broker string `json:"broker"`
	// Exchange to publish to, defaults to amq.topic
	Exchange string `json:"exchange,omitempty"`
	// Routing key for published messages
	RoutingKey string `json:"routingKey"`
}

// HttpPost destination
type HttpPost struct {
	Url string `json:"url"`
}
