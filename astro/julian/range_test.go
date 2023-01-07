package julian

import (
	"reflect"
	"testing"
)

func TestRange_Duration(t *testing.T) {
	type fields struct {
		invalid bool // true then ignore start/end
		start   Day
		end     Day
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{name: "positive", fields: fields{start: 10, end: 20}, want: 10},
		{name: "negative", fields: fields{start: 10, end: 20}, want: 10},
		{name: "empty_valid", fields: fields{start: 0, end: 0}, want: 0},
		{name: "empty_invalid", fields: fields{invalid: true}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			if !tt.fields.invalid {
				r = r.Include(tt.fields.start).Include(tt.fields.end)
			}
			if got := r.Duration(); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_End(t *testing.T) {
	type fields struct {
		invalid bool // true then ignore start/end
		empty   bool // true then ignore start,end
		start   Day
		end     Day
	}
	tests := []struct {
		name   string
		fields fields
		want   Day
	}{
		{name: "after", fields: fields{start: 10, end: 20}, want: 20},
		{name: "before", fields: fields{start: 20, end: 10}, want: 20},
		{name: "empty_valid", fields: fields{empty: true}, want: 0},
		{name: "empty_invalid", fields: fields{invalid: true}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			if !tt.fields.invalid {
				if !tt.fields.empty {
					r = r.IncludePeriod(tt.fields.start, tt.fields.end)
				}
			}
			if got := r.End(); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_ForEach(t *testing.T) {
	type args struct {
		start   Day
		end     Day
		step    float64
		empty   bool
		invalid bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []Day
	}{
		{name: "ascending", args: args{start: 10, end: 20, step: 2}, want: []Day{10, 12, 14, 16, 18, 20}},
		{name: "descending", args: args{start: 10, end: 20, step: -2}, want: []Day{20, 18, 16, 14, 12, 10}},
		{name: "single", args: args{start: 10, end: 20, step: 0}, want: []Day{10}},
		{name: "single_swap", args: args{start: 20, end: 10, step: 0}, want: []Day{10}},
		{name: "empty", args: args{empty: true}, want: nil},
		{name: "invalid", args: args{invalid: true}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			var got []Day

			// If invalid then leave r nil
			if !tt.args.invalid {
				// otherwise create an empty instance
				r = &Range{}

				// add period if not required to be empty
				if !tt.args.empty {
					r = r.IncludePeriod(tt.args.start, tt.args.end)
				}
			}

			err := r.ForEach(tt.args.step, func(d Day) error {
				got = append(got, d)
				return nil
			})

			switch {
			case tt.wantErr && err == nil:
				t.Errorf("expected an error got none")
			case !tt.wantErr && err != nil:
				t.Errorf("unexpected error %v", err)
			case !reflect.DeepEqual(got, tt.want):
				t.Errorf("got %v, want %v", got, tt.want)
			default:
				// test passed
			}
		})
	}
}

func TestRange_Include(t *testing.T) {
	type fields struct {
		invalid bool // true then ignore start/end
		start   Day
		step    float64
	}
	tests := []struct {
		name   string
		fields fields
		want   []Day
	}{
		{name: "after", fields: fields{start: 10}, want: []Day{10}},
		{name: "before", fields: fields{start: 20}, want: []Day{20}},
		{name: "empty_valid", fields: fields{start: 10}, want: []Day{10}},
		{name: "empty_invalid", fields: fields{invalid: true}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			if !tt.fields.invalid {
				r = r.Include(tt.fields.start)
			}
			if got := r.Slice(tt.fields.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_IncludeDays(t *testing.T) {
	type args struct {
		days []Day
		step float64
	}
	tests := []struct {
		name string
		args args
		want []Day
	}{
		{
			name: "range1",
			args: args{step: 2, days: []Day{4, 21, 9, 28, 5}},
			want: []Day{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28},
		},
		{
			// same as range1 but with fewer entries
			name: "range2",
			args: args{step: 2, days: []Day{4, 5, 28}},
			want: []Day{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			r = r.IncludeDays(tt.args.days...)
			got := r.Slice(tt.args.step)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncludeDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_IncludePeriod(t *testing.T) {
	type args struct {
		a    Day
		b    Day
		step float64
	}
	tests := []struct {
		name string
		args args
		want []Day
	}{
		{name: "a_b_ascending", args: args{a: 4, b: 10, step: 2}, want: []Day{4, 6, 8, 10}},
		{name: "a_b_descending", args: args{a: 4, b: 10, step: -2}, want: []Day{10, 8, 6, 4}},
		{name: "b_a_ascending", args: args{a: 10, b: 4, step: 2}, want: []Day{4, 6, 8, 10}},
		{name: "b_a_descending", args: args{a: 10, b: 4, step: -2}, want: []Day{10, 8, 6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			r = r.IncludePeriod(tt.args.a, tt.args.b)

			got := r.Slice(tt.args.step)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncludePeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_IncludeRange(t *testing.T) {
	type args struct {
		a1   Day
		b1   Day
		a2   Day
		b2   Day
		step float64
	}
	tests := []struct {
		name string
		args args
		want []Day
	}{
		{name: "a_b_ascending", args: args{a1: 4, b1: 8, a2: 6, b2: 10, step: 2}, want: []Day{4, 6, 8, 10}},
		{name: "a_b_descending", args: args{a1: 4, b1: 8, a2: 6, b2: 10, step: -2}, want: []Day{10, 8, 6, 4}},
		{name: "b_a_ascending", args: args{a1: 8, b1: 4, a2: 6, b2: 10, step: 2}, want: []Day{4, 6, 8, 10}},
		{name: "b_a_descending", args: args{a1: 8, b1: 4, a2: 10, b2: 6, step: -2}, want: []Day{10, 8, 6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r1, r2 *Range
			r1 = r1.IncludePeriod(tt.args.a1, tt.args.b1)
			r2 = r2.IncludePeriod(tt.args.a2, tt.args.b2)

			r1 = r1.IncludeRange(r2)
			got := r1.Slice(tt.args.step)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncludeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Start(t *testing.T) {
	type fields struct {
		invalid bool // true then ignore start/end
		start   Day
		end     Day
	}
	tests := []struct {
		name   string
		fields fields
		want   Day
	}{
		{name: "after", fields: fields{start: 10, end: 20}, want: 10},
		{name: "before", fields: fields{start: 20, end: 10}, want: 10},
		{name: "empty_valid", fields: fields{start: 10, end: 20}, want: 10},
		{name: "empty_invalid", fields: fields{invalid: true}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r *Range
			if !tt.fields.invalid {
				r = r.IncludePeriod(tt.fields.start, tt.fields.end)
			}
			if got := r.Start(); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
