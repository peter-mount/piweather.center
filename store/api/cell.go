package api

import (
	"strconv"
	"strings"
	"time"
)

// Cell represents an individual cell within a Table's Row
type Cell struct {
	Type   CellType  // Type of cell
	Time   time.Time // Time of value in cell, IsZero()==true if unknown or text
	String string    // String value, always present as formatted by Unit if Float or Int
	Float  float64   // float64 value, only set when unmarshalling from JSON
}

// CellType defines the type of cell
type CellType uint8

const (
	CellString  CellType = iota // Cell is a String
	CellNumeric                 // Cell is a numeric value
	CellNull                    // Cell is not present, String="" but treat like a SQL Null
	CellDynamic                 // Same as CellString but acts like CellNull, e.g. when determining if a row is empty
)

// MarshalJSON simplifies the JSON output of a cell to a single value, be it null, float or string.
func (c *Cell) MarshalJSON() ([]byte, error) {
	var s string
	switch {
	case c == nil, c.Type == CellNull:
		s = "null"
	case c.Type == CellNumeric:
		s = c.String
	default:
		s = `"` + c.String + `"`
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
		c.String = strings.Trim(s, `"`)

	default:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		c.Type = CellNumeric
		c.String = s
		c.Float = f
	}
	return nil
}
