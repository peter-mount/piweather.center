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
			query: `AT "now" SELECT garden.cps`,
		},
		{
			name:  "AT 12:00",
			query: `AT "2023-12-13T12:00" SELECT garden.cps`,
		},
		{
			name:  "AT Today",
			query: `AT "today" SELECT garden.cps`,
		},
		{
			name:  "AT Tomorrow",
			query: `AT "TOMORROW" SELECT garden.cps`,
		},
		{
			name:  "AT Yesterday",
			query: `AT "YESTERDAY" SELECT garden.cps`,
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
