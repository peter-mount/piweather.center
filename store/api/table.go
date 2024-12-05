package api

import (
	"fmt"
	"io"
	"strings"
)

type Table struct {
	Columns []*Column `json:"columns" xml:"columns" yaml:"columns"`
	Rows    []*Row    `json:"rows" xml:"rows" yaml:"rows"`
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
	err := w.uint16(uint16(len(t.Columns)))
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

func (t *Table) ColumnCount() int {
	return len(t.Columns)
}

func (t *Table) GetColumnByIndex(i int) *Column {
	return t.Columns[i]
}

func (t *Table) AddColumn(c *Column) *Table {
	if c.Width < len(c.Name) {
		c.Width = len(c.Name)
	}

	t.Columns = append(t.Columns, c)
	return t
}

// NewRow adds a new row to the Table
func (t *Table) NewRow() *Row {
	r := &Row{}
	t.Rows = append(t.Rows, r)
	return r
}

// CurrentRow returns the current (last) Row in the table.
// If the Table is empty then a new row will be returned.
func (t *Table) CurrentRow() *Row {
	// If called before NewRow() then implicitly call it
	if len(t.Rows) == 0 {
		return t.NewRow()
	}
	// Return the last row in the table
	return t.Rows[len(t.Rows)-1]
}

// PruneCurrentRow will remove the last row in the table if it's not valid
func (t *Table) PruneCurrentRow() *Table {
	if t.CurrentRowPrunable() {
		t.Rows = t.Rows[:len(t.Rows)-1]
	}
	return t
}

// CurrentRowPrunable will return true if  the table is not empty and the
// current (last) row is not valid.
func (t *Table) CurrentRowPrunable() bool {
	return len(t.Rows) > 0 && !t.Rows[len(t.Rows)-1].IsValid()
}

func (t *Table) RowCount() int {
	return len(t.Rows)
}

func (t *Table) GetRow(i int) *Row {
	return t.Rows[i]
}

func (t *Table) IsEmpty() bool {
	return len(t.Rows) == 0
}

// Finalise ensures all columns are defined and filters out rows with no data in them
func (t *Table) Finalise() {
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

func (t *Table) String(b []string) []string {
	// Create line separator
	var s0, s1 []string
	for _, col := range t.Columns {
		s0 = append(s0, strings.Repeat("-", col.Width))
		s1 = append(s1, fmt.Sprintf(fmt.Sprintf("%%%d.%ds", col.Width, col.Width), col.Name))
	}
	head := "+" + strings.Join(s0, "+") + "+"
	sep := "|" + strings.Join(s0, "|") + "|"

	// Add table header
	b = append(b, head, "|"+strings.Join(s1, "|")+"|")

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
