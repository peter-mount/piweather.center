package julian

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/util"
	"testing"
	"time"
)

func TestIsGregorian(t *testing.T) {
	type args struct {
		d int
		m int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "battle hastings", args: args{d: 14, m: 10, y: 1066}, want: false},
		{name: "early 1582", args: args{d: 30, m: 4, y: 1582}, want: false},
		{name: "end julian", args: args{d: 4, m: 10, y: 1582}, want: false},
		{name: "start gregorian", args: args{d: 15, m: 10, y: 1582}, want: true},
		{name: "late 1582", args: args{d: 25, m: 12, y: 1582}, want: true},
		{name: "sputnik", args: args{d: 4, m: 10, y: 1957}, want: true},
		{name: "21 century", args: args{d: 12, m: 2, y: 2022}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGregorian(tt.args.d, tt.args.m, tt.args.y); got != tt.want {
				t.Errorf("IsGregorian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want Day
	}{
		{name: "Sputnik", args: args{t: time.Date(1957, 10, 4, 19, 26, 24, 0, time.UTC)}, want: 2436116.31},
		{name: "year333", args: args{t: time.Date(333, 1, 27, 12, 0, 0, 0, time.UTC)}, want: 1842713.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromTime(tt.args.t); got != tt.want {
				t.Errorf("FromTime() = %f, want %f", got, tt.want)
			}
		})
	}
}

func TestDay_Time(t *testing.T) {
	tests := []struct {
		t Day
		y int
		m time.Month
		d float64
	}{
		{y: 1957, m: time.October, d: 4.81, t: 2436116.31},
		{y: 2000, m: time.January, d: 1.5, t: 2451545.0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f", tt.t), func(t *testing.T) {
			got := tt.t.Time()
			td, th, tm, ts := util.FdayToHMS(tt.d)
			if got.Year() != tt.y || got.Month() != tt.m ||
				got.Day() != td || got.Hour() != th || got.Minute() != tm || got.Second() != ts {
				t.Errorf("Time() = %v, want %d %d %.4f", got, tt.y, tt.m, tt.d)
			}
		})
	}
}

func TestDay_IsGregorian(t *testing.T) {
	type args struct {
		t Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Sputnik", args: args{t: FromDate(1957, 10, 4, 19, 26, 24)}, want: true},
		{name: "year333", args: args{t: FromDate(333, 1, 27, 12, 0, 0)}, want: false},
		{name: "1582_early", args: args{t: FromDate(1582, 5, 20, 12, 0, 0)}, want: false},
		{name: "1582_10_1", args: args{t: FromDate(1582, 10, 1, 12, 0, 0)}, want: false},
		{name: "1582_10_4", args: args{t: FromDate(1582, 10, 4, 12, 0, 0)}, want: false},
		{name: "1582_10_15", args: args{t: FromDate(1582, 10, 15, 12, 0, 0)}, want: true},
		{name: "1582_10_31", args: args{t: FromDate(1582, 10, 31, 12, 0, 0)}, want: true},
		{name: "1582_late", args: args{t: FromDate(1582, 12, 31, 12, 0, 0)}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.IsGregorian(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
