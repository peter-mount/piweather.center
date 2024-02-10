package weatherimport

import (
	"context"
	"flag"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/tools/weatheringress/model"
	"github.com/peter-mount/piweather.center/tools/weatheringress/payload"
	"github.com/peter-mount/piweather.center/util"
	"os"
	"sort"
	"time"
)

type Importer struct {
	Daemon    *kernel.Daemon `kernel:"inject"`
	Config    model.Loader   `kernel:"inject"`
	Store     file.Store     `kernel:"inject"`
	DBServer  *string        `kernel:"flag,import,Import data to DB url"`
	MaxBuffer *int           `kernel:"flag,buffer-size,Number of metrics before submitting to DB,1000"`
	BaseDir   *string        `kernel:"flag,dir,Directory to find files"`
	client    client.Client
	visitor   model.VisitorBuilder
	metrics   []api.Metric
	lineCount int       // number of lines read
	count     int       // number of metrics imported in a file
	total     int       // number of metrics imported in total
	rTime     time.Time // time of last report - used in logging submissions
}

const (
	// importTimeout is the time between logging how many metrics have been imported.
	// This allows logging to show progress if the DB or the connection to the DB is slow
	importTimeout = 3 * time.Second
)

func (i *Importer) IsImporting() bool {
	return i.DBServer != nil && *i.DBServer != ""
}

func (i *Importer) PostInit() error {
	if i.MaxBuffer == nil || *i.MaxBuffer < 1 {
		*i.MaxBuffer = 256
	}

	return nil
}

func (i *Importer) Start() error {
	i.Daemon.ClearDaemon()

	log.Println(version.Version)

	log.Println("Starting Import")
	startTime := time.Now()

	i.total = 0
	i.client = client.Client{Url: *i.DBServer}

	i.visitor = model.NewVisitor().
		Reading(i.processReading)

	// Take the files to read from the arguments
	files := flag.Args()

	// if -dir set then look for .log files in there
	if i.BaseDir != nil && *i.BaseDir != "" {
		err := walk.NewPathWalker().
			Then(func(path string, _ os.FileInfo) error {
				files = append(files, path)
				return nil
			}).
			PathHasSuffix(".log").
			IsFile().
			Walk(*i.BaseDir)
		if err != nil {
			return err
		}
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	for _, fileName := range files {
		if err := i.importFile(fileName); err != nil {
			return err
		}
	}

	endTime := time.Now()
	log.Printf("Imported %d metrics in %s", i.total, endTime.Sub(startTime))

	return nil
}

func (i *Importer) importFile(fileName string) error {
	i.lineCount = 0
	i.count = 0
	i.rTime = time.Now()

	n := fileName
	if len(n) > 40 {
		n = "..." + n[len(n)-40:]
	}
	log.Printf("Reading %s", n)

	err := io.NewReader().
		ForEachLine(func(line string) error {
			i.lineCount++

			p, err := payload.FromLog(line)
			if p != nil && err == nil {
				return i.Config.Accept(i.visitor.WithContext(p.AddContext(context.Background())))
			}
			return err
		}).
		Open(fileName)

	if err == nil {
		i.rTime = time.Time{}
		err = i.submitMetrics()
	}

	i.total += i.count

	return err
}

func (i *Importer) processReading(ctx context.Context) error {
	r := model.ReadingFromContext(ctx)
	if r.Unit() == nil {
		return nil
	}

	p := payload.GetPayload(ctx)

	str, ok := p.Get(r.Source)
	if !ok || str == "" {
		return nil
	}

	if f, ok := util.ToFloat64(str); ok {
		// Convert to As unit then transform to Use unit
		v, err := r.Value(f)
		if err != nil {
			// Ignore, should only happen if the result is
			// invalid as we checked the transform previously
			log.Printf("Conversion failed for %.3f -> %s on %s", f, r.Unit().ID(), r.ID)
			return err
		}

		// Submit to the DB directly.
		// Note use v not f here as it may have been transformed so a different value
		return i.record(api.Metric{
			Metric: r.ID,
			Time:   p.Time(),
			Unit:   v.Unit().ID(),
			Value:  v.Float(),
		})
	}

	log.Printf("Failed to parse %s = %q", r.ID, str)
	return nil
}

func (i *Importer) record(m api.Metric) error {
	i.metrics = append(i.metrics, m)
	i.count++

	if len(i.metrics) >= *i.MaxBuffer || time.Now().Sub(i.rTime) >= importTimeout {
		return i.submitMetrics()
	}
	return nil
}

func (i *Importer) submitMetrics() error {
	if len(i.metrics) > 0 {
		_, err := i.client.RecordMetrics(i.metrics)

		// Log if we have gone over a limit
		now := time.Now()
		if err == nil && now.Sub(i.rTime) >= importTimeout {
			log.Printf("Imported %d metrics %d lines", i.count, i.lineCount)
			i.rTime = now
		}

		i.metrics = nil

		return err
	}
	return nil
}
