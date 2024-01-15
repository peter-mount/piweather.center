package api

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"strconv"
	"strings"
	"time"
)

// Cell represents an individual cell within a Table's Row
type Cell struct {
	Type   CellType    // Type of cell
	Time   time.Time   // Time of value in cell, IsZero()==true if unknown or text
	string string      // String value, always present as formatted by Unit if Float or Int
	float  float64     // float64 value, only set when unmarshalling from JSON
	Value  value.Value // Converted value
}

// CellType defines the type of cell
type CellType uint8

const (
	CellString  CellType = iota // Cell is a String
	CellNumeric                 // Cell is a numeric value
	CellNull                    // Cell is not present, String="" but treat like a SQL Null
	CellDynamic                 // Same as CellString but acts like CellNull, e.g. when determining if a row is empty
)

func NewNullCell() Cell {
	return Cell{Type: CellNull}
}

func NewNumericCell(t time.Time, s string, v float64) Cell {
	return Cell{
		Type:   CellNumeric,
		Time:   t,
		string: s,
		float:  v,
	}
}

func NewStringCell(t time.Time, s string) Cell {
	return Cell{
		Type:   CellString,
		Time:   t,
		string: s,
	}
}

func NewDynamicCell(t time.Time, s string) Cell {
	return Cell{
		Type:   CellDynamic,
		Time:   t,
		string: s,
	}
}

func NewValueCell(t time.Time, v value.Value) Cell {
	if v.IsValid() {
		return NewNumericCell(t, v.PlainString(), v.Float())
	}
	return NewNullCell()
}

// MarshalJSON simplifies the JSON output of a cell to a single value, be it null, float or string.
func (c *Cell) MarshalJSON() ([]byte, error) {
	var s string
	switch {
	case c == nil, c.Type == CellNull:
		s = "null"
	case c.Type == CellNumeric:
		s = c.string
	default:
		s = `"` + c.string + `"`
	}
	return []byte(s), nil
}

func (c *Cell) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch {
	case s == "null":
		c.Type = CellNull

	case strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`):
		c.Type = CellString
		c.string = strings.Trim(s, `"`)
		c.Time, _ = time.Parse(time.RFC3339, c.string)

	default:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		c.Type = CellNumeric
		c.string = s
		c.float = f
	}
	return nil
}

func (c *Cell) Float() float64 {
	return c.float
}

func (c *Cell) Int() int {
	return int(c.float)
}

func (c *Cell) String() string {
	return c.string
}
