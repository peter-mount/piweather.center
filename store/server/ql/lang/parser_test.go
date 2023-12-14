package lang

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
			name:  "AT NOW",
			query: `SELECT garden.cps AT "now"`,
		},
		{
			name:  "AT 12:00",
			query: `SELECT garden.cps AT "2023-12-13T12:00"`,
		},
		{
			name:  "AT Today",
			query: `SELECT garden.cps AT "today"`,
		},
		{
			name:  "AT Tomorrow",
			query: `SELECT garden.cps AT "TOMORROW"`,
		},
		{
			name:  "AT Yesterday",
			query: `SELECT garden.cps AT "YESTERDAY"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New().
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