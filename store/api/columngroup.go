package api

// ColumnGroup represents a group of columns in a table
type ColumnGroup struct {
	Width int    `json:"width" xml:"width,attr" yaml:"width"`
	Name  string `json:"name" xml:"name" yaml:"name"`
}

func (c *ColumnGroup) write(w *writer) error {
	err := w.uint16(uint16(c.Width))
	if err == nil {
		err = w.string(c.Name)
	}
	return err
}

func (c *ColumnGroup) read(r *reader) error {
	ui, err := r.uint16()
	if err == nil {
		c.Width = int(ui)
		c.Name, err = r.string()
	}
	return err
}

func (t *Table) ColumnGroupCount() int {
	return len(t.ColumnGroups)
}

func (t *Table) ColumnGroupTotalWidth() int {
	w := 0
	for _, g := range t.ColumnGroups {
		w = w + g.Width
	}
	return w
}

func (t *Table) LastColumnGroup() *ColumnGroup {
	l := len(t.ColumnGroups)
	if l == 0 {
		return nil
	}
	return t.ColumnGroups[l-1]
}

// AddColumnGroup adds a column group with the specified width
func (t *Table) AddColumnGroup(name string, width int) *Table {
	// If we have columns already but no groups then add them to an untitled group
	if t.ColumnGroupCount() == 0 && t.ColumnCount() > 0 {
		t.ColumnGroups = append(t.ColumnGroups, &ColumnGroup{Name: "", Width: t.ColumnCount()})
	}

	/*
		// Check the last group is full
		lg := t.LastColumnGroup()
		if lg != nil {
			// If we have less than the expected number of columns in a group then reduce its width
			gw := t.ColumnGroupTotalWidth()
			cc := t.ColumnCount()
			//fmt.Printf("cc %d ctw %d\n", cc, gw)
			if cc < gw {
				lg.Width = lg.Width - (gw - cc)
				if lg.Width > 0 {
					//fmt.Printf("add cg 0 %v\n", t.ColumnGroups)
					return
				}

				// If the last group is empty then remove it
				if len(t.ColumnGroups) > 0 {
					t.ColumnGroups = t.ColumnGroups[:len(t.ColumnGroups)-1]
				} else {
					t.ColumnGroups = nil
				}
			}
		}*/

	// Add the new group
	t.ColumnGroups = append(t.ColumnGroups, &ColumnGroup{Name: name, Width: width})
	//fmt.Printf("add cg 1 %v\n", t.ColumnGroups)
	return t
}
