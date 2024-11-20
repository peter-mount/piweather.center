package table

import (
	"fmt"
	"strconv"
)

func New(c ...string) *Table {
	t := &Table{}

	for _, s := range c {
		t.AddColumn(s)
	}

	return t
}

type Table struct {
	columns []*Column
	rows    []*Row
}

type Column struct {
	name  string
	width int
}

type Row struct {
	table *Table
	cells []Cell
}

type Cell struct {
	value string
}

func (t *Table) AddColumn(s string) *Table {
	t.columns = append(t.columns, &Column{name: s})
	return t
}

func (t *Table) NewRow() *Row {
	r := &Row{table: t}
	t.rows = append(t.rows, r)
	return r
}

func (r *Row) Add(s string) *Row {
	r.cells = append(r.cells, Cell{value: s})
	if len(r.cells) > len(r.table.columns) {
		r.table.AddColumn("Undefined")
	}
	return r
}

func (r *Row) AddInt(i int) *Row {
	return r.Add(strconv.Itoa(i))
}

func (r *Row) AddFloat(f float64) *Row {
	return r.Add(strconv.FormatFloat(f, 'f', -1, 64))
}

func (r *Row) AddBool(b bool) *Row {
	return r.Add(strconv.FormatBool(b))
}

func (r *Row) AddF(format string, args ...interface{}) *Row {
	return r.Add(fmt.Sprintf(format, args...))
}