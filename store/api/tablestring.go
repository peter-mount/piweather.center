package api

import (
	"fmt"
	"strings"
)

func (t *Table) String(b []string) []string {
	// Create line separator
	var s0, s1 []string
	for _, col := range t.Columns {
		s0 = append(s0, strings.Repeat("-", col.Width))
		s1 = append(s1, fmt.Sprintf(fmt.Sprintf("%%%d.%ds", col.Width, col.Width), col.Name))
	}
	head := "+" + strings.Join(s0, "+") + "+"
	sep := "|" + strings.Join(s0, "|") + "|"

	if t.ColumnGroupCount() > 0 {
		var g0, g1 []string
		c := 0
		for _, cg := range t.ColumnGroups {
			w := 0
			for j := 0; j < cg.Width; j++ {
				w = w + t.Columns[c].Width
				c++
			}

			// Add separators within the group
			w = w + cg.Width - 1

			g0 = append(g0, strings.Repeat("-", w))
			g1 = append(g1, fmt.Sprintf(fmt.Sprintf("%%%d.%ds", w, w), cg.Name))
		}
		b = append(b,
			"+"+strings.Join(g0, "+")+"+",
			"|"+strings.Join(g1, "|")+"|",
		)
	}

	// Add table header
	b = append(b, head, "|"+strings.Join(s1, "|")+"|", head)

	for i, r := range t.Rows {
		if i == 0 || r.RowType == RowTypeSummary {
			b = append(b, sep)
		}
		s1 = nil
		for i, c := range r.Cells {
			s1 = append(s1, t.Columns[i].String(c.String()))
		}
		b = append(b, "|"+strings.Join(s1, "|")+"|")
	}

	rc := len(t.Rows)
	if rc > 0 {
		b = append(b, head)
	}
	b = append(b, fmt.Sprintf("Rows: %d", rc))

	return b
}
