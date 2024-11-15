package table

import "strings"

func (t *Table) String() string {
	var s []string

	for _, col := range t.columns {
		col.width = len(col.name)
	}

	for _, row := range t.rows {
		for c, cell := range row.cells {
			l := len(cell.value)
			if l > t.columns[c].width {
				t.columns[c].width = l
			}
		}
	}

	var cs []string
	l := 0
	for _, col := range t.columns {
		cs = append(cs, pad(col.width, col.name))
		l = l + col.width + 1
	}

	s = append(s,
		strings.Join(cs, " "),
		strings.Repeat("-", l-1))

	for _, row := range t.rows {
		cs = nil
		for c, cell := range row.cells {
			cs = append(cs, pad(t.columns[c].width, cell.value))
		}
		s = append(s, strings.Join(cs, " "))
	}

	return strings.Join(s, "\n")
}

func pad(w int, s string) string {
	return s + strings.Repeat(" ", w-len(s))
}
