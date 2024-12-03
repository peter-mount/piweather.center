package api

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

func (r *Result) NewTable() *Table {
	t := &Table{}
	r.Table = append(r.Table, t)
	return t
}

type Table struct {
	Columns []*Column `json:"columns" xml:"columns" yaml:"columns"`
	Rows    []*Row    `json:"rows" xml:"rows" yaml:"rows"`
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
				*r = append(*r, &Cell{Type: CellNull})
			}

			// now ensure column widths are wide enough
			for i, c := range *r {
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
			if i < len(*r) {
				return (*r)[i]
			}
			break
		}
	}
	return &Cell{Type: CellNull}
}

type Column struct {
	Index int         `json:"index"`           // Index of column in Row
	Name  string      `json:"name"`            // Name of column
	Type  ColumnType  `json:"type,omitempty"`  // Type of column
	Width int         `json:"width,omitempty"` // Width in characters, 0=unknown
	Unit  string      `json:"unit,omitempty"`  // Unit of column, ""=unknown or text
	unit  *value.Unit // resolved unit
}

func (c *Column) write(w *writer) error {
	err := w.int16(int16(c.Index))

	if err == nil {
		err = w.string(c.Name)
	}

	if err == nil {
		err = w.int16(int16(c.Type))
	}

	if err == nil {
		err = w.int16(int16(c.Width))
	}

	if err == nil {
		// Hash of 0 is nil
		var h uint64
		if c.unit != nil {
			h = c.unit.Hash()
		}
		err = w.uint64(h)
	}
	return err
}

func (c *Column) read(r *reader) error {
	v, err := r.int16()
	if err == nil {
		c.Index = int(v)
	}

	if err == nil {
		c.Name, err = r.string()
	}

	if err == nil {
		v, err = r.int16()
		c.Type = ColumnType(v)
	}

	if err == nil {
		v, err = r.int16()
		c.Width = int(v)
	}

	if err == nil {
		h, err1 := r.uint64()
		if err1 != nil {
			return err
		}

		if h > 0 {
			u, ok := value.GetUnitByHash(h)
			if ok {
				c.unit = u
				c.Unit = u.Unit()
			}
		}
	}

	return err
}

// IsFixed returns true if the column is of fixed width
func (c *Column) IsFixed() bool { return c != nil && (c.Type&ColumnFixed) == ColumnFixed }

func (c *Column) IsDefault() bool { return c != nil && (c.Type&ColumnCenter) == ColumnDefault }
func (c *Column) IsLeft() bool    { return c != nil && (c.Type&ColumnCenter) == ColumnLeft }
func (c *Column) IsRight() bool   { return c != nil && (c.Type&ColumnCenter) == ColumnRight }
func (c *Column) IsCenter() bool  { return c != nil && (c.Type&ColumnCenter) == ColumnCenter }

func (c *Column) SetUnit(u *value.Unit) {
	c.unit = u
	if u == nil {
		c.Unit = ""
	} else {
		c.Unit = c.unit.ID()
	}
}

func (c *Column) Transform(v value.Value) (value.Value, error) {
	if c.unit == nil && v.IsValid() {
		c.SetUnit(v.Unit())
	}
	if c.unit != nil {
		return v.As(c.unit)
	}
	return v, nil
}

func (c *Column) Value(f float64) (value.Value, error) {
	if c.unit == nil {
		u, ok := value.GetUnit(c.Unit)
		if !ok {
			return value.Value{}, fmt.Errorf("unknown unit %q", c.Unit)
		}
		c.SetUnit(u)
	}
	if c.unit != nil {
		return c.unit.Value(f), nil
	}
	return value.Value{}, nil
}

// ColumnType defines the Column type. This is a bit field so each value defines the type of each bit.
type ColumnType int16

const (
	ColumnDefault = 0x00 // Default, align left
	ColumnLeft    = 0x01 // Align left
	ColumnRight   = 0x02 // Align right
	ColumnCenter  = 0x03 // Align center
	ColumnFixed   = 0x10 // Fixed width column
)

func (c *Column) String(s string) string {
	if len(s) >= c.Width {
		return s[:c.Width]
	}
	switch {
	case len(s) == c.Width:
		return s
	case len(s) > c.Width:
		return s[:c.Width]
	case c.IsLeft(), c.IsDefault():
		return c.pad(s, 0, c.Width-len(s))
	case c.IsRight():
		return c.pad(s, c.Width-len(s), 0)
	case c.IsCenter():
		p := (c.Width - len(s)) >> 1
		return c.pad(s, p, p)
	}
	return s
}

func (c *Column) pad(s string, l, e int) string {
	var r []byte
	if l > 0 {
		r = append(r, strings.Repeat(" ", l)...)
	}
	r = append(r, s...)
	if e > 0 {
		r = append(r, strings.Repeat(" ", e)...)
	}
	for len(r) < c.Width {
		r = append(r, ' ')
	}
	return string(r[:c.Width])
}

// Row Holds details about an individual
type Row []*Cell // Individual columns

func (r *Row) IsEmpty() bool {
	return len(*r) == 0
}

func (r *Row) CellCount() int {
	return len(*r)
}

func (r *Row) Cell(i int) *Cell {
	return (*r)[i]
}

func (r *Row) write(w *writer) error {
	err := w.uint16(uint16(len(*r)))
	if err == nil {
		for _, c := range *r {
			err = c.write(w)
			if err != nil {
				break
			}
		}
	}
	return err
}

func (r *Row) read(rd *reader) error {
	l, err := rd.uint16()
	if err == nil && l > 0 {
		for i := 0; i < int(l); i++ {
			c := &Cell{}
			err = c.read(rd)
			if err != nil {
				break
			}
			*r = append(*r, c)
		}
	}
	return err
}

func (r *Row) add(c Cell) *Row {
	*r = append(*r, &c)
	return r
}

// AddValue adds a CellString cell to the row based on the value.Value.
// If the value is not valid then a CellNull cell is added instead.
func (r *Row) AddValue(t time.Time, v value.Value) *Row {
	return r.add(NewValueCell(t, v))
}

// AddString adds a CellString cell with the supplied value
func (r *Row) AddString(t time.Time, s string) *Row {
	return r.add(NewStringCell(t, s))
}

// AddDynamic adds a CellDynamic cell with the supplied value.
func (r *Row) AddDynamic(t time.Time, s string) *Row {
	return r.add(NewDynamicCell(t, s))
}

// AddNull adds a CellNull cell to the row.
func (r *Row) AddNull() *Row {
	return r.add(NewNullCell())
}

// IsValid returns true of the row contains at least one cell not CellNull or CellDynamic
func (r *Row) IsValid() bool {
	for _, c := range *r {
		if !(c.Type == CellNull || c.Type == CellDynamic) {
			return true
		}
	}

	return false
}

func (r *Row) Size() int {
	return len(*r)
}

func (r *Result) String() string {
	if r == nil {
		return ""
	}

	var b []string

	b = append(b, fmt.Sprintf("Status: %d", r.Status))
	if r.Message != "" {
		b = append(b, r.Message)
	}

	if r.Meta != nil {
		for k, v := range r.Meta {
			b = append(b, fmt.Sprintf("%q = %v", k, v))
		}
	}

	for _, t := range r.Table {
		b = t.String(b)
	}

	return strings.Join(b, "\n") + "\n"
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
		if i == 0 {
			b = append(b, sep)
		}
		s1 = nil
		for i, c := range *r {
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
