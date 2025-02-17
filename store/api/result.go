package api

import (
	"fmt"
	"io"
	"strings"
)

type Result struct {
	// Status of the result
	Status int `json:"status"`
	// Additional metadata from the query
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Optional message from status
	Message string `json:"message,omitempty"`
	// Range of the data
	Range *Range `json:"range,omitempty"`
	// Results
	Table    []*Table    `json:"table,omitempty"`
	WindRose []*WindRose `json:"windRose,omitempty"`
}

func (r *Result) write(w *writer) error {
	err := w.int16(int16(r.Status))

	if err == nil {
		err = w.string(r.Message)
	}

	// TODO Meta

	if err == nil {
		err = w.bool(r.Range != nil)
		if err == nil && r.Range != nil {
			err = r.Range.write(w)
		}
	}

	if err == nil {
		err = w.uint16(uint16(len(r.Table)))
		if err == nil {
			for _, t := range r.Table {
				err = t.write(w)
				if err != nil {
					break
				}
			}
		}
	}

	return err
}

func (r *Result) read(rd *reader) error {
	v, err := rd.uint16()
	if err == nil {
		r.Status = int(v)
	}

	if err == nil {
		r.Message, err = rd.string()
	}

	// TODO meta

	if err == nil {
		b, err1 := rd.bool()
		if err1 != nil {
			return err1
		}
		if b {
			r.Range = &Range{}
			err = r.Range.read(rd)
		}
	}

	if err == nil {
		// Number of tables
		v, err = rd.uint16()
		if err == nil && v > 0 {
			for i := 0; i < int(v); i++ {
				t := &Table{}
				err = t.read(rd)
				if err != nil {
					break
				}
				r.Table = append(r.Table, t)
			}
		}
	}

	return err
}

// Close the Result
func (r *Result) Close() error {
	if r != nil {
		r.Meta = nil
		r.Range = nil
		r.Table = nil
		r.WindRose = nil
	}
	return nil
}

func (r *Result) Init() {
	for _, t := range r.Table {
		for _, r := range t.Rows {
			for i, c := range r.GetCells() {
				if c == nil {
					r.SetCell(i, NewNullCell())
				} else if c.Type == CellNumeric {
					v, _ := t.Columns[i].Value(c.Float())
					r.SetCell(i, NewValueCell(c.Time, v))
				}
			}
		}
	}
}

func (r *Result) AddMeta(k string, v interface{}) {
	if r.Meta == nil {
		r.Meta = make(map[string]interface{})
	}
	r.Meta[k] = v
}

func (r *Result) Finalise() {
	for _, t := range r.Table {
		t.Finalise()
	}

	for _, wr := range r.WindRose {
		wr.Finalise()
	}
}

func (r *Result) NewTable() *Table {
	t := NewTable()
	r.Table = append(r.Table, t)
	return t
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

func (r *Result) Read(src io.Reader) error {
	return r.read(newReader(src))
}

func (r *Result) Write(wr io.Writer) error {
	return r.write(newWriter(wr))
}
