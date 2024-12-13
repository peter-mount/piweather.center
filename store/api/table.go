package api

import (
	"fmt"
	"io"
)

type Table struct {
	ColumnGroups []*ColumnGroup `json:"column_groups" xml:"column_groups" yaml:"column_groups"`
	Columns      []*Column      `json:"columns" xml:"columns" yaml:"columns"`
	Rows         []*Row         `json:"rows" xml:"rows" yaml:"rows"`
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Read(r io.Reader) error {
	return t.read(newReader(r))
}

func (t *Table) Write(w io.Writer) error {
	return t.write(newWriter(w))
}

func (t *Table) write(w *writer) error {
	err := w.uint16(uint16(len(t.ColumnGroups)))
	if err == nil {
		for _, c := range t.ColumnGroups {
			err = c.write(w)
			if err != nil {
				break
			}
		}
	}

	err = w.uint16(uint16(len(t.Columns)))
	if err == nil {
		for _, c := range t.Columns {
			err = c.write(w)
			if err != nil {
				break
			}
		}
	}

	if err == nil {
		err = w.uint16(uint16(len(t.Rows)))
	}
	if err == nil {
		for _, r := range t.Rows {
			err = r.write(w)
			if err != nil {
				break
			}
		}
	}

	return err
}

func (t *Table) read(r *reader) error {
	var err error

	v, err := r.uint16()
	if err == nil && v > 0 {
		for i := 0; i < int(v); i++ {
			c := &ColumnGroup{}
			err = c.read(r)
			if err != nil {
				break
			}
			t.ColumnGroups = append(t.ColumnGroups, c)
		}
	}

	v, err = r.uint16()
	if err == nil && v > 0 {
		for i := 0; i < int(v); i++ {
			c := &Column{}
			err = c.read(r)
			if err != nil {
				break
			}
			t.Columns = append(t.Columns, c)
		}
	}

	if err == nil {
		v, err = r.uint16()
	}
	if err == nil && v > 0 {
		for i := 0; i < int(v); i++ {
			row := &Row{}
			err = row.read(r)
			if err != nil {
				break
			}
			t.Rows = append(t.Rows, row)
		}
	}

	return err
}

func (t *Table) IsEmpty() bool {
	return len(t.Rows) == 0
}

// Finalise ensures all columns are defined and filters out rows with no data in them
func (t *Table) Finalise() *Table {
	var tr []*Row

	// Ensure each column width is set to contain all values
	for _, r := range t.Rows {
		if r.IsValid() {
			// If row is shorter than columns add null columns to the table
			for r.Size() < len(t.Columns) {
				r.Cells = append(r.Cells, &Cell{Type: CellNull})
			}

			// now ensure column widths are wide enough
			for i, c := range r.Cells {
				// If we have more entries in the row and columns then add a new one to the table
				for len(t.Columns) <= i {
					t.Columns = append(t.Columns, &Column{Name: fmt.Sprintf("Col%d", len(t.Columns)+1)})
				}

				// Ensure the column width is wide enough - but do not change fixed width columns
				cw := len(c.String())
				col := t.Columns[i]
				if !col.IsFixed() && cw > col.Width {
					col.Width = cw
				}
			}

			tr = append(tr, r)
		}
	}

	for i, c := range t.Columns {
		c.Index = i
	}

	t.Rows = tr

	return t
}

func (t *Table) GetColumn(n string) *Column {
	for _, c := range t.Columns {
		if n == c.Name {
			return c
		}
	}
	return nil
}

func (t *Table) GetCell(n string, r *Row) *Cell {
	for i, c := range t.Columns {
		if n == c.Name {
			if i < len(r.Cells) {
				return r.Cells[i]
			}
			break
		}
	}
	return &Cell{Type: CellNull}
}
