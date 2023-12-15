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
		// ============================================================
		// Standard UPPER case
		// ============================================================
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
		// ============================================================
		// Tests lower case to ensure the parser is case insensitive
		// ============================================================
		{
			name:  "Case at NOW",
			query: `at "now" select garden.cps`,
		},
		{
			name:  "Case at 12:00",
			query: `at "2023-12-13T12:00" select garden.cps`,
		},
		{
			name:  "Case at Today",
			query: `at "today" select garden.cps`,
		},
		{
			name:  "Case at Tomorrow",
			query: `at "TOMORROW" select garden.cps`,
		},
		{
			name:  "Case at Yesterday",
			query: `at "YESTERDAY" select garden.cps`,
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
