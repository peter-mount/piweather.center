package weatheregress

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-build/version"
	egress2 "github.com/peter-mount/piweather.center/config/egress"
	amqp2 "github.com/peter-mount/piweather.center/config/util/amqp"
	"github.com/peter-mount/piweather.center/mq/amqp"
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

	return egress2.NewBuilder[mqSetup]().
		Amqp(s.initAmqp).
		Build().
		SetData(m).
		Script(s.script)
}

func (s *Processor) initAmqp(v egress2.Visitor[mqSetup], a *amqp2.Amqp) error {
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

func (s *Processor) publishAmqp(v egress2.Visitor[*action], p *egress2.Publish, msg []byte) error {
	a := s.script.State().GetAmqp(p.Amqp)
	if a == nil {
		return participle.Errorf(p.Pos, "Unknown amqp %q", p.Amqp)
	}

	act := v.GetData()
	return a.Publisher.Publish(act.routingKey, msg)
}
