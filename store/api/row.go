package api

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type RowType uint8

const (
	// RowTypeData represents a normal row containing data
	RowTypeData = iota
	// RowTypeSummary represents a row containing summary information
	RowTypeSummary
)

// Row Holds details about an individual
type Row struct {
	RowType RowType `json:"rowType" xml:"rowType,attr" yaml:"rowType"`
	Cells   []*Cell `json:"cells" xml:"cells" yaml:"cells"` // Individual columns
}

func (r *Row) IsEmpty() bool {
	return len(r.Cells) == 0
}

func (r *Row) CellCount() int {
	return len(r.Cells)
}

func (r *Row) GetCells() []*Cell {
	return r.Cells
}

func (r *Row) Cell(i int) *Cell {
	if i < 0 || i >= len(r.Cells) {
		return nil
	}
	return r.Cells[i]
}

func (r *Row) write(w *writer) error {
	err := w.uint8(uint8(r.RowType))

	if err == nil {
		err = w.uint16(uint16(len(r.Cells)))
	}

	if err == nil {
		for _, c := range r.Cells {
			err = c.write(w)
			if err != nil {
				break
			}
		}
	}
	return err
}

func (r *Row) read(rd *reader) error {
	rt, err := rd.uint8()
	if err == nil {
		r.RowType = RowType(rt)

		l, err := rd.uint16()
		if err == nil && l > 0 {
			for i := 0; i < int(l); i++ {
				c := &Cell{}
				err = c.read(rd)
				if err != nil {
					break
				}
				r.Cells = append(r.Cells, c)
			}
		}
	}
	return err
}

func (r *Row) add(c *Cell) *Row {
	r.Cells = append(r.Cells, c)
	return r
}

func (r *Row) SetCell(i int, c *Cell) {
	for len(r.Cells) <= i {
		r.AddNull()
	}
	r.Cells[i] = c
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
	for _, c := range r.Cells {
		if !(c.Type == CellNull || c.Type == CellDynamic) {
			return true
		}
	}

	return false
}

func (r *Row) Size() int {
	return len(r.Cells)
}
