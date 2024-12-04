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
	var sep string
	for i, col := range t.columns {
		cs = append(cs, pad(col.width, col.name))
		if i > 0 {
			sep = sep + "-+-"
		}
		sep = sep + strings.Repeat("-", col.width)
	}

	colSep := " | "
	s = append(s, sep, strings.Join(cs, colSep), sep)

	for _, row := range t.rows {
		cs = nil
		for c, cell := range row.cells {
			cs = append(cs, pad(t.columns[c].width, cell.value))
		}
		s = append(s, strings.Join(cs, colSep))
	}

	s = append(s, sep)

	return strings.Join(s, "\n")
}

func pad(w int, s string) string {
	return s + strings.Repeat(" ", w-len(s))
}
