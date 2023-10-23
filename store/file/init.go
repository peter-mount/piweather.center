package file

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	"os"
	"strings"
	"time"
)

// initLatest ensures that the latest values are read from disk on startup.
// Note: This only works against the current da
func (s *store) initLatest() error {
	prefix := *s.BaseDir + "/"

	day := time.Hour * 24
	now := time.Now().UTC().Add(-day)

	// Run across yesterday and today to ensure we have data in the scenario where we are at UTC midnight with little data
	for d := 0; d < 2; d++ {
		suffix := fmt.Sprintf("/%d/%d/%d.mdb", now.Year(), now.Month(), now.Day())

		err := walk.NewPathWalker().
			Then(func(path string, _ os.FileInfo) error {
				metric := strings.ReplaceAll(strings.TrimSuffix(strings.TrimPrefix(path, prefix), suffix), "/", ".")
				rec, err := s.GetLatestRecord(metric, now)
				if err == nil && rec.IsValid() {
					s.Latest.Append(metric, rec)
				}
				return err
			}).
			PathHasSuffix(suffix).
			IsFile().
			Walk(*s.BaseDir)

		if err != nil {
			return err
		}

		now = now.Add(day)
	}

	return nil
}
