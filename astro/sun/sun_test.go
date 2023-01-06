package sun

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"testing"
)

func TestSunCoordLow(t *testing.T) {
	type args struct {
		jd julian.Day
	}
	tests := []struct {
		name string
		args args
		want coord.Equatorial
	}{
		{
			name: "1992 10 13",
			args: args{jd: 2448908.5},
			want: coord.Equatorial{RA: "13:13:31", Dec: "S07:47:06"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Low(tt.args.jd)
			if got.Equatorial.RA != tt.want.RA || got.Equatorial.Dec != tt.want.Dec {
				fmt.Println(got.String())
				t.Errorf("Low() %q,%q want: %q,%q\n%s", got.Equatorial.RA, got.Equatorial.Dec, tt.want.RA, tt.want.Dec, got.String())
			}
		})
	}
}
