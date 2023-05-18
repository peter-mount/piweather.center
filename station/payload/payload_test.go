package payload

import (
	"github.com/peter-mount/piweather.center/station"
	"testing"
	"time"
)

func TestPayload_Get(t *testing.T) {
	tests := []struct {
		json    string
		path    string
		want    time.Time
		wantErr bool
	}{
		{
			"{\"timestamp\":\"2023-05-17T12:07:16Z\"}",
			"timestamp",
			time.Date(2023, 5, 17, 12, 7, 16, 0, time.UTC),
			false,
		},
		{
			"{\"results\":{\"timestamp\":\"2023-05-17T12:07:16Z\"}}",
			"results.timestamp",
			time.Date(2023, 5, 17, 12, 7, 16, 0, time.UTC),
			false,
		},
		{
			// This should fail as timestamp is under results but Payload will
			// have the current UTC time not that in the sample
			"{\"results\":{\"timestamp\":\"2023-05-17T12:07:16Z\"}}",
			"timestamp",
			time.Date(2023, 5, 17, 12, 7, 16, 0, time.UTC),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			s1 := station.Sensors{
				Name:      "",
				Source:    station.Source{},
				Format:    "",
				Timestamp: tt.path,
				Readings:  nil,
			}

			p, err := s1.FromBytes([]byte(tt.json))
			if err != nil {
				t.Errorf("FromBytes failed %v", err)
				return
			}

			got := p.Time()
			if got != tt.want && !tt.wantErr {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
