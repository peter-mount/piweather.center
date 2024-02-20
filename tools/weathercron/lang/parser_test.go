package lang

import "testing"

func Test_Parser(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		wantErr bool
	}{
		// ============================================================
		// Schedules
		// ============================================================
		{
			name: "cron",
			query: `
task cron "* * * * *"
run "false"
`,
		}, {
			name:  "every day",
			query: `task every "day" run "false"`,
		}, {
			name:  "every hour",
			query: `task every "hour" run "false"`,
		}, {
			name:  "every minute",
			query: `task every "minute" run "false"`,
		}, {
			name:  "every second",
			query: `task every "second" run "false"`,
		}, {
			name:  "every dummy",
			query: `task every "dummy" run "false"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewParser().
				ParseString(tt.name, tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
