package api

import "time"

// Cell represents an individual cell within a Table's Row
type Cell struct {
	Type   CellType  `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"` // Type of cell
	Time   time.Time `json:"time,omitempty" xml:"time,attr,omitempty" yaml:"time,omitempty"` // Time of value in cell, IsZero()==true if unknown or text
	String string    `json:"string" xml:",cdata" yaml:"string"`                              // String value, always present as formatted by Unit if Float or Int
}

// CellType defines the type of cell
type CellType uint8

const (
	CellString  CellType = iota // Cell is a String
	CellNumeric                 // Cell is a numeric value
	CellNull                    // Cell is not present, String="" but treat like a SQL Null
	CellDynamic                 // Same as CellString but acts like CellNull, e.g. when determining if a row is empty
)

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
