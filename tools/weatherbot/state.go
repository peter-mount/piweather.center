package bot

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/weather/state"
	"strings"
)

// getCurrentState retrieves the current station state from weathercenter
func (t *Bot) getCurrentState() error {
	// Get current state for the station for this post
	stn, err := state.New(*t.Host).GetState(t.post.StationId)
	if err != nil {
		return err
	}
	if stn == nil {
		return fmt.Errorf("StationId %q does not exist", t.post.StationId)
	}

	// For the bot, IDs are case-insensitive
	for _, m := range stn.Measurements {
		m.ID = strings.ToLower(m.ID)
	}

	log.Printf("Station %q data at %v", stn.ID, stn.Meta.Time)
	t.station = stn
	return nil
}

// getMeasurement returns a specific Measurement or nil if not provided by the station
func (t *Bot) getMeasurement(id string) *state.Measurement {
	id = strings.ToLower(id)
	for _, m := range t.station.Measurements {
		if m.ID == id {
			return m
		}
	}
	return nil
}
