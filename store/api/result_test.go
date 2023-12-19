package api

import (
	"encoding/json"
	"testing"
)

func TestTable_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		wantCols int
		wantRows int
	}{
		{
			name:     "Small Table",
			json:     `{"table":[{"columns":[{"name":"time","width":20},{"name":"temp","width":4,"unit":"Celsius"},{"name":"trend","width":5,"unit":"Integer"},{"name":"noise","width":5,"unit":"Volt"},{"name":"trend","width":5,"unit":"Integer"},{"name":"pm2_5","width":5,"unit":"Micrograms Per Cubic Meter"},{"name":"trend","width":5,"unit":"Integer"}],"rows":[["2023-12-19T08:40:00Z",10.9,0,0.340,1,8,0],["2023-12-19T08:50:00Z",11.0,1,0.266,-1,32,1],["2023-12-19T09:00:00Z",11.0,0,0.264,-1,10,null],["2023-12-19T09:10:00Z",11.0,0,0.252,-1,20,1],["2023-12-19T09:20:00Z",11.1,1,0.523,1,42,1],["2023-12-19T09:30:00Z",11.1,0,0.387,-1,22,null]]}]}`,
			wantCols: 7,
			wantRows: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Result{}
			err := json.Unmarshal([]byte(tt.json), r)
			if err != nil {
				t.Errorf("json unmarshal error %v", err)
				return
			}

			if len(r.Table) != 1 {
				t.Errorf("Expected 1 table got %d", len(r.Table))
				return
			}

			table := r.Table[0]

			if len(table.Columns) != tt.wantCols {
				t.Errorf("Expected %d columns got %d", tt.wantCols, len(table.Columns))
			}

			if len(table.Rows) != tt.wantRows {
				t.Errorf("Expected %d rows got %d", tt.wantRows, len(table.Rows))
			}

			for i, row := range table.Rows {
				if len(*row) != tt.wantCols {
					t.Errorf("Row %d expected %d columns got %d", i, tt.wantCols, len(*row))
				}

				/*for j, c := range *row {
					fmt.Printf("Col (%d,%d) type %d %q %f\n",
						j, i,
						c.Type,
						c.String,
						c.Float,
					)
				}*/
			}
		})
	}
}
