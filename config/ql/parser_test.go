package ql

import (
	"testing"
)

func Test_Parser(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		wantErr bool
	}{
		{
			name:  "at now",
			query: `at "now" select garden.cps`,
		},
		{
			name:  "at 12:00",
			query: `at "2023-12-13T12:00" select garden.cps`,
		},
		{
			name:  "at today",
			query: `at "today" select garden.cps`,
		},
		{
			name:  "at tomorrow",
			query: `at "tomorrow" select garden.cps`,
		},
		{
			name:  "at yesterday",
			query: `at "yesterday" select garden.cps`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBuilder().
				ParseString(tt.name, tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ParseString() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
