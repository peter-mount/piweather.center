package amqp

import (
	"github.com/alecthomas/participle/v2/lexer"
	amqp2 "github.com/peter-mount/piweather.center/util/mq/amqp"
	"strings"
)

// Amqp broker definition
type Amqp struct {
	Pos       lexer.Position
	Name      string `parser:"'amqp' '(' 'name' @String"`
	Url       string `parser:"'url' @String"`
	Exchange  string `parser:"('exchange' @String)? ')'"`
	MQ        *amqp2.MQ
	Publisher *amqp2.Publisher
}

func (s *Amqp) Init() error {
	if s != nil {
		s.Name = strings.ToLower(strings.TrimSpace(s.Name))

		// Exchange is optional, default to amq.topic
		s.Exchange = strings.TrimSpace(s.Exchange)
		if s.Exchange == "" {
			s.Exchange = "amq.topic"
		}
	}

	return nil
}
