package julian

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		want    Day
		wantErr bool
	}{
		{name: "2459587.5", want: 2459587.5},
		{name: "2022-01-08", want: FromDate(2022, 1, 8, 0, 0, 0)},
		{name: "2022-01-08T12:34:56", want: FromDate(2022, 1, 8, 12, 34, 56)},
		{name: "2022-01-08T12:34", want: FromDate(2022, 1, 8, 12, 34, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
