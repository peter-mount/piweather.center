package value

type Range struct {
	min Value
	max Value
	u   *Unit
}

func NewRange(u *Unit) *Range {
	// A new empty range.
	// Note: we do set min to max and max to min at first, so that when
	// we add the first value to the range it will then set the min/max correctly.
	return &Range{
		min: u.Value(u.Max()),
		max: u.Value(u.Min()),
		u:   u,
	}
}

// IsValid returns true if the Range is valid.
// Specifically that it has had a value entered into it so that min<=max
func (r *Range) IsValid() bool {
	return LessThanEqual(r.min.Float(), r.max.Float())
}

// Add a value to this Range.
// This will return an error if the supplied value cannot be transformed to the unit of this Range.
func (r *Range) Add(v Value) error {
	if r.IsValid() {
		b, err := v.LessThan(r.min)
		if err != nil {
			return err
		}
		if b {
			r.min = v
		}

		b, err = v.GreaterThan(r.max)
		if err != nil {
			return err
		}
		if b {
			r.max = v
		}
	} else {
		r.min = v
		r.max = v
	}

	return nil
}

func (r *Range) Unit() *Unit { return r.u }

func (r *Range) Min() Value { return r.min }

func (r *Range) Max() Value { return r.max }

func (r *Range) Range() (Value, Value) { return r.min, r.max }

func (r *Range) Include(b *Range) error {
	if b.IsValid() {
		if err := r.Add(b.Min()); err != nil {
			return err
		}
		return r.Add(b.Max())
	}
	return nil
}

func (r *Range) Clone() *Range {
	return &Range{
		min: r.min,
		max: r.max,
		u:   r.u,
	}
}
