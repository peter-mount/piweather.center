package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type HttpFormatType uint8

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
			if util.IsVisitorStop(err) {
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
