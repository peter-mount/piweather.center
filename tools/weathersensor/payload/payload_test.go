package payload

import (
	"github.com/peter-mount/piweather.center/config/station"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"strings"
	"testing"
	"time"
)

func TestPayload_Get(t *testing.T) {
	tests := []struct {
		json    string
		path    station.SourcePath
		want    time.Time
		wantErr bool
	}{
		{
			"{\"timestamp\":\"2023-05-17T12:07:16Z\"}",
			station.SourcePath{Path: []string{"timestamp"}},
			time2.ParseTime("2023-05-17T12:07:16Z"),
			false,
		},
		{
			"{\"results\":{\"timestamp\":\"2023-05-17T12:07:16Z\"}}",
			station.SourcePath{Path: []string{"results", "timestamp"}},
			time2.ParseTime("2023-05-17T12:07:16Z"),
			false,
		},
		{
			// This should fail as timestamp is under results but Payload will
			// have the current UTC time not that in the sample
			"{\"results\":{\"timestamp\":\"2023-05-17T12:07:16Z\"}}",
			station.SourcePath{Path: []string{"timestamp"}},
			time2.ParseTime("2023-05-17T12:07:16Z"),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.path.Path, "."), func(t *testing.T) {
			p, err := FromBytes("", station.HttpFormatJson, &tt.path, []byte(tt.json))
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
