package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type DashboardList struct {
	Pos        lexer.Position
	Dashboards []*Dashboard `parser:"@@*"`
}

func (c *visitor[T]) DashboardList(d *DashboardList) error {
	var err error
	if d != nil {
		if c.dashboardList != nil {
			err = c.dashboardList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Dashboards {
				err = c.Dashboard(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initDashboardList(v Visitor[*initState], _ *DashboardList) error {
	v.Get().dashboards = make(map[string]*Dashboard)
	return nil
}

func (b *builder[T]) DashboardList(f func(Visitor[T], *DashboardList) error) Builder[T] {
	b.dashboardList = f
	return b
}
