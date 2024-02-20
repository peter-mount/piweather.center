package lang

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/util"
	"strings"
)

func NewParser() util.Parser[Script] {
	return util.NewParser[Script](nil, nil, cronInit)
}

func cronInit(q *Script, err error) (*Script, error) {
	if err == nil {
		state := NewState()

		err = NewBuilder[*State]().
			Every(initEvery).
			Build().
			SetData(state).
			Script(q)

		if err == nil {
			q.state = state.Cleanup()
		}
	}
	return q, err
}

func initEvery(v Visitor[*State], a *Every) error {
	// Convert aliases to actual definitions
	switch strings.ToLower(a.Definition) {
	case "day", "daily", "midnight":
		a.Cron = "0 0 0 * * *"
	case "hour", "hourly":
		a.Cron = "0 0 * * * *"
	case "minute":
		a.Cron = "0 * * * * *"
	case "second":
		a.Cron = "* * * * * *"
	default:
		return participle.Errorf(a.Pos, "unsupported every %q", a.Definition)
	}

	return nil
}
