package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"strings"
)

type HttpFormatType uint8

func (t HttpFormatType) String() string {
	return httpFormatName[t]
}

const (
	// HttpFormatJson indicates data will be in JSON
	HttpFormatJson HttpFormatType = iota
	// HttpFormatXml indicates data will be in XML
	HttpFormatXml
	// HttpFormatYaml indicates data will be in YAML
	HttpFormatYaml
	// HttpFormatPost indicates data will be in http post format. Also used for query string etc
	HttpFormatPost
	// HttpFormatCarbon is formatted string: "key value timestamp" where timestamp is in unix seconds
	HttpFormatCarbon
)

var (
	httpFormatName = []string{"json", "xml", "yaml", "post", "carbon"}
)

// HttpFormat defines the format of the incoming data.
// This maps to the formats supported by Payload
type HttpFormat struct {
	Pos    lexer.Position
	Carbon bool `parser:"'format' '(' ( @'carbon'"`
	Json   bool `parser:"| @'json'"`
	Post   bool `parser:"| @('post'|'query'|'query' 'string'|'form')"`
	XML    bool `parser:"| @'xml'"`
	YAML   bool `parser:"| @('yaml'|'yml') ) ')'"`
}

func (c *visitor[T]) HttpFormat(d *HttpFormat) error {
	var err error
	if d != nil {
		if c.httpFormat != nil {
			err = c.httpFormat(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) HttpFormat(f func(Visitor[T], *HttpFormat) error) Builder[T] {
	b.httpFormat = f
	return b
}

// GetType returns the HttpFormatType. Default if not defined is HttpFormatJson
func (d *HttpFormat) GetType() HttpFormatType {
	if d != nil {
		switch {
		case d.Json:
			return HttpFormatJson
		case d.Carbon:
			return HttpFormatCarbon
		case d.Post:
			return HttpFormatPost
		case d.XML:
			return HttpFormatXml
		case d.YAML:
			return HttpFormatYaml
		}
	}
	return HttpFormatJson
}

func ParseHttpFormatType(s string) HttpFormatType {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "carbon":
		return HttpFormatCarbon
	case "json":
		return HttpFormatJson
	case "post", "query", "querystring", "form":
		return HttpFormatPost
	case "xml":
		return HttpFormatXml
	case "yaml", "yml":
		return HttpFormatYaml
	default:
		return HttpFormatJson
	}
}
