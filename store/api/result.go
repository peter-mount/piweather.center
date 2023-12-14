package api

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	Table []*Table `json:"table" xml:"table" yaml:"table"`
}

func (r *Result) NewTable() *Table {
	t := &Table{}
	r.Table = append(r.Table, t)
	return t
}

func (r *Result) Finalise() {
	for _, t := range r.Table {
		t.Finalise()
	}
}

type Table struct {
	Columns []*Column `json:"columns" xml:"columns" yaml:"columns"`
	Rows    []*Row    `json:"rows" xml:"rows" yaml:"rows"`
}

func (t *Table) AddColumn(c Column) *Table {
	// TODO should we enforce column width to the name length or just default to this if Width==0?
	if c.Width < len(c.Name) {
		c.Width = len(c.Name)
	}

	t.Columns = append(t.Columns, &c)
	return t
}

func (t *Table) NewRow() *Row {
	r := &Row{}
	t.Rows = append(t.Rows, r)
	return r
}

func (t *Table) CurrentRow() *Row {
	// If called before NewRow() then implicitly call it
	if len(t.Rows) == 0 {
		return t.NewRow()
	}
	// Return the last row in the table
	return t.Rows[len(t.Rows)-1]
}

func (t *Table) Finalise() {
	// Ensure each column width is set to contain all values
	for _, r := range t.Rows {

		// If row is shorter than columns add null columns to the table
		for len(r.Columns) < len(t.Columns) {
			r.Columns = append(r.Columns, Cell{Type: CellNull})
		}

		// now ensure column widths are wide enough
		for i, c := range r.Columns {
			// If we have more entries in the row and columns then add a new one to the table
			for len(t.Columns) <= i {
				t.Columns = append(t.Columns, &Column{Name: fmt.Sprintf("Col%d", len(t.Columns)+1)})
			}

			// Ensure the column width is wide enough - but do not change fixed width columns
			cw := len(c.String)
			col := t.Columns[i]
			if !col.IsFixed() && cw > col.Width {
				col.Width = cw
			}
		}
	}
}

type Column struct {
	Name  string      `json:"name" xml:",chardata" yaml:"name"`                                  // Name of column
	Type  ColumnType  `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"`    // Type of column
	Width int         `json:"width,omitempty" xml:"width,omitempty,attr" yaml:"width,omitempty"` // Width in characters, 0=unknown
	Unit  string      `json:"unit,omitempty" xml:"unit,omitempty,attr" yaml:"unit,omitempty"`    // Unit of column, ""=unknown or text
	unit  *value.Unit // resolved unit
}

// IsFixed returns true if the column is of fixed width
func (c *Column) IsFixed() bool { return c != nil && (c.Type&ColumnFixed) == ColumnFixed }

func (c *Column) IsDefault() bool { return c != nil && (c.Type&ColumnCenter) == ColumnDefault }
func (c *Column) IsLeft() bool    { return c != nil && (c.Type&ColumnCenter) == ColumnLeft }
func (c *Column) IsRight() bool   { return c != nil && (c.Type&ColumnCenter) == ColumnRight }
func (c *Column) IsCenter() bool  { return c != nil && (c.Type&ColumnCenter) == ColumnCenter }

func (c *Column) SetUnit(u value.Value) {
	if u.IsValid() && c.unit == nil {
		c.unit = u.Unit()
		c.Unit = c.unit.Name()
	}
}

func (c *Column) Transform(v value.Value) (value.Value, error) {
	if c.unit == nil {
		c.SetUnit(v)
	}
	if c.unit != nil {
		return v.As(c.unit)
	}
	return v, nil
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
type Row struct {
	Columns []Cell `json:"columns" xml:"columns" yaml:"columns"` // Individual columns
}

func (r *Row) Add(c Cell) *Row {
	r.Columns = append(r.Columns, c)
	return r
}

func (r *Row) AddFloat(t time.Time, f float64) *Row {
	return r.Add(Cell{
		Type:   CellFloat,
		Time:   t,
		Float:  f,
		String: strconv.FormatFloat(f, 'f', 3, 64),
	})
}

func (r *Row) AddInt(t time.Time, i int64) *Row {
	return r.Add(Cell{
		Type:   CellInt,
		Time:   t,
		Int:    i,
		String: strconv.FormatInt(i, 64),
	})
}

func (r *Row) AddString(t time.Time, s string) *Row {
	return r.Add(Cell{
		Type:   CellString,
		Time:   t,
		String: s,
	})
}

func (r *Row) AddNull() *Row {
	return r.Add(Cell{Type: CellNull})
}

// Cell represents an individual cell within a Table's Row
type Cell struct {
	Type   CellType  `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"`    // Type of cell
	Time   time.Time `json:"time,omitempty" xml:"time,attr,omitempty" yaml:"time,omitempty"`    // Time of value in cell, IsZero()==true if unknown or text
	Float  float64   `json:"float,omitempty" xml:"float,attr,omitempty" yaml:"float,omitempty"` // Float value
	Int    int64     `json:"int,omitempty" xml:"int,attr,omitempty" yaml:"int,omitempty"`       // Int64 value
	String string    `json:"string" xml:",cdata" yaml:"string"`                                 // String value, always present as formatted by Unit if Float or Int
}

func (c Cell) IsNull() bool { return c.Type == CellNull }

func (c Cell) HasTime() bool { return !(c.IsNull() || c.Time.IsZero()) }

// CellType defines the type of cell
type CellType uint8

const (
	CellString CellType = iota // Cell is a String
	CellFloat                  // Cell is a Float
	CellInt                    // Cell is an Int
	CellNull                   // Cell is not present, String="" but treat like a SQL Null
)

func (r *Result) String() string {
	if r == nil {
		return ""
	}

	var b []string
	for _, t := range r.Table {
		b = t.String(b)
	}
	return strings.Join(b, "\n") + "\n"
}

func (t *Table) String(b []string) []string {
	// Create line break
	var s0, s1 []string
	for _, col := range t.Columns {
		s0 = append(s0, strings.Repeat("-", col.Width))
		s1 = append(s1, fmt.Sprintf(fmt.Sprintf("%%%d.%ds", col.Width, col.Width), col.Name))
	}
	sep := "|" + strings.Join(s0, "|") + "|"

	// Add table header
	b = append(b, sep, "|"+strings.Join(s1, "|")+"|", sep)

	for _, r := range t.Rows {
		s1 = nil
		for i, c := range r.Columns {
			s1 = append(s1, t.Columns[i].String(c.String))
		}
		b = append(b, "|"+strings.Join(s1, "|")+"|")
	}

	return b
}
