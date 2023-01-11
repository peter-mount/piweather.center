package cloud

// Payload is a json representation of the detected cloud levels.
// It's available to be written to disk and/or sent to a message broker
type Payload struct {
	Cloud float64 `json:"cloud" xml:"cloud,attr"`
	Sky   float64 `json:"sky" xml:"sky,attr"`
	Okta  string  `json:"okta" xml:"okta,attr"`
}
