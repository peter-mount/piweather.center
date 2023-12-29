package api

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

func (r *Result) Init() {
	for _, t := range r.Table {
		for _, r := range t.Rows {
			for i, c := range *r {
				if c.Type == CellNumeric {
					v, _ := t.Columns[i].Value(c.Float())
					c.Value = v
					(*r)[i] = c
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
