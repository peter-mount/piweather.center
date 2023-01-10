package sidereal

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	"testing"
)

func TestFromJD(t *testing.T) {
	type args struct {
		jd julian.Day
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1987_Apr_10", args: args{jd: 2446895.5}, want: "13:10:46"},
		{name: "1987_Apr_10_1921", args: args{jd: 2446896.30625}, want: "08:34:57"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromJD(tt.args.jd); util.HourDMSString(got) != tt.want {
				t.Errorf("FromJD() = %v, want %v", util.HourDMSString(got), tt.want)
			}
		})
	}
}

func TestAtGreenwichMidnight(t *testing.T) {
	type args struct {
		jd julian.Day
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// These two are for the same day so the want will be the same for both regardless of time of day
		{name: "1987_Apr_10", args: args{jd: 2446895.5}, want: "13:10:46"},
		{name: "1987_Apr_10_1921", args: args{jd: 2446896.30625}, want: "13:10:46"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AtGreenwichMidnight(tt.args.jd); util.HourDMSString(got) != tt.want {
				t.Errorf("AtGreenwichMidnight() = %v, want %v", got, tt.want)
			}
		})
	}
}
