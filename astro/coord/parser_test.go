package coord

import (
	"testing"
)

func TestParse(t *testing.T) {
	type want struct {
		long float64
		lat  float64
		alt  float64
		name string
	}
	tests := []struct {
		name    string
		want    want
		wantErr bool
	}{
		{name: "12.34,56.78", want: want{long: 12.34, lat: 56.78}},
		{name: "12:45,56:15", want: want{long: 12.75, lat: 56.25}},
		{name: "12.34,56.78,129", want: want{long: 12.34, lat: 56.78, alt: 129}},
		{name: "12:45,56:15,9283", want: want{long: 12.75, lat: 56.25, alt: 9283}},
		{name: "12.34,56.78,129,Home", want: want{long: 12.34, lat: 56.78, alt: 129, name: "Home"}},
		{name: "12:45,56:15,9283,Somewhere Else", want: want{long: 12.75, lat: 56.25, alt: 9283, name: "Somewhere Else"}},
		{name: "12:45,56:15,9283,Commas, are allowed here", want: want{long: 12.75, lat: 56.25, alt: 9283, name: "Commas, are allowed here"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.long != got.Coord().Lon.Deg() ||
				tt.want.lat != got.Coord().Lat.Deg() ||
				tt.want.alt != got.Altitude ||
				tt.want.name != got.Name {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
