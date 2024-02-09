package weatheregress

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
	"strings"
)

type mqSetup struct {
	appName    string
	appVersion string
}

func (s *Processor) initMq() error {
	vs := strings.SplitN(version.Version, " ", 2)
	m := mqSetup{
		appName:    strings.Join([]string{"piweather.center", vs[0]}, " "),
		appVersion: strings.Trim(vs[1], "()"),
	}

	return lang.NewBuilder[mqSetup]().
		Amqp(s.initAmqp).
		Build().
		SetData(m).
		Script(s.script)
}

func (s *Processor) initAmqp(v lang.Visitor[mqSetup], a *lang.Amqp) error {
	m := v.GetData()

	a.MQ = &amqp.MQ{
		Url:            a.Url,
		Exchange:       a.Exchange,
		ConnectionName: m.appName + " Egress",
		Product:        m.appName,
		Version:        m.appVersion,
	}

	a.Publisher = &amqp.Publisher{
		Exchange:  a.Exchange,
		Mandatory: false,
		Immediate: false,
		Tag:       strings.TrimSpace(strings.Join([]string{m.appName, a.Name}, " ")),
	}

	return a.Publisher.Bind(a.MQ)
}

func (s *Processor) publishAmqp(v lang.Visitor[*action], p *lang.Publish) error {
	a := s.script.State().GetAmqp(p.Amqp)
	if a == nil {
		return participle.Errorf(p.Pos, "Unknown amqp %q", p.Amqp)
	}

	act := v.GetData()

	// FIXME this works for strings, but need to handle json as well
	msg, err := calculator.GetString(act.message)
	if err == nil {
		err = a.Publisher.Publish(act.routingKey, []byte(msg))
	}

	return err
}
