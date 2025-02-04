package config

import (
	"github.com/peter-mount/piweather.center/station"
	"slices"
	"strings"
)

var (
	suffix map[station.LoadOption]string
	types  map[string]station.LoadOption
)

func init() {
	// The usual file suffix for specific options
	suffix = map[station.LoadOption]string{
		station.CalculationOption: ".calc",
		station.DashboardOption:   ".station",
		station.SensorOption:      ".sensor",
		station.JobOption:         ".cron",
	}

	// The aliases for options the tool accepts
	types = map[string]station.LoadOption{
		// file suffix names (the . is stripped if provided)
		"calc":    station.CalculationOption,
		"station": station.DashboardOption,
		"sensor":  station.SensorOption,
		"cron":    station.JobOption,
		// Aliases for types
		"dash":      station.DashboardOption,
		"dashboard": station.DashboardOption,
		"job":       station.JobOption,
		"sensors":   station.SensorOption,
		// Dummy entries to allow for all types, "" happens when "." is passed
		"all": station.AllOption,
		"":    station.AllOption,
	}

}

// loadOption returns the LoadOption for a specific name.
// Any leading '.' is stripped.
// Returns station.NoOption if not recognised.
func loadOption(suffix string) station.LoadOption {
	// Ensure the file suffix starts with a single .
	for suffix != "" && suffix[0] == '.' {
		suffix = suffix[1:]
	}

	return types[strings.ToLower(suffix)]
}

// suffixes returns the possible suffixes for the given LoadOption.
func suffixes(l station.LoadOption) []string {
	m := make(map[string]bool)
	for i := 0; i < 8; i++ {
		v := station.LoadOption(1 << i)
		if l.Is(v) {
			if s, e := suffix[v]; e {
				m[s] = true
			}
		}
	}

	var a []string
	for k, _ := range m {
		a = append(a, k)
	}
	slices.Sort(a)
	return a
}
