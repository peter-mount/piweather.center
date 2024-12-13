package api

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

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
