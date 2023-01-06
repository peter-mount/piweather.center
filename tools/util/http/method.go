package http

// Method represents a supported HTTP method
type Method int

const (
	Get Method = iota
	Post
)

func (m Method) String() string {
	switch m {
	case Get:
		return "GET"
	case Post:
		return "POST"
	default:
		return "UNKNOWN"
	}
}
