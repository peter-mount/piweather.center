package weatherdb

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/station/service"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"os"
	"path/filepath"
	"strings"
)

type Importer struct {
	Config service.Config `kernel:"inject"`
	Store  file.Store     `kernel:"inject"`
	Import *string        `kernel:"flag,metric-import,Import log archives from directory"`
}

func (i *Importer) Start() error {
	if *i.Import == "" {
		return nil
	}

	log.Println("Scanning log archives")

	ctx := value.WithMap(context.Background())

	return i.Config.Accept(station.NewVisitor().
		Sensors(i.importSensor).
		WithContext(ctx))
}

func (i *Importer) importSensor(ctx context.Context) error {
	visitor := station.NewVisitor().
		Sensors(value.ResetMap).
		Reading(i.processReading)

	sensors := station.SensorsFromContext(ctx)

	dir := filepath.Join(*i.Import, filepath.Join(strings.Split(sensors.ID, ".")...))

	err := walk.NewPathWalker().
		Then(func(path string, _ os.FileInfo) error {
			log.Printf("Reading %s", path)
			return io.NewReader().
				ForEachLine(func(line string) error {
					p, err := payload.FromLog(line)
					if p != nil && err == nil {
						_ = visitor.
							WithContext(p.AddContext(ctx)).
							VisitSensors(sensors)
					}
					return nil
				}).
				Open(path)
		}).
		PathHasSuffix(".log").
		IsFile().
		Walk(dir)

	if os.IsNotExist(err) {
		err = nil
	}

	return err
}

func (i *Importer) processReading(ctx context.Context) error {
	r := station.ReadingFromContext(ctx)
	if r.Unit() != nil {
		p := payload.GetPayload(ctx)

		str, ok := p.Get(r.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			// Convert to Type unit then transform to Use unit
			v, err := r.Value(f)
			if err != nil {
				// Ignore, should only happen if the result is
				// invalid as we checked the transform previously
				return nil
			}

			return i.Store.Append(r.ID, file.Record{
				Time:  p.Time(),
				Value: v,
			})
		}
	}
	return nil
}
