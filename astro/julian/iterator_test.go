package julian

import (
	"reflect"
	"testing"
)

func TestIterate(t *testing.T) {
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
		{
			name: "increment",
			args: args{a: 0, b: 10, step: 1},
			want: []Day{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "increment_a_after_b",
			args: args{a: 10, b: 0, step: 1},
			want: []Day{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "decrement",
			args: args{a: 0, b: 10, step: -1},
			want: []Day{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			// Pass a and b in reverse, tests the iterator handles this correctly
			name: "decrement_a_after_b",
			args: args{a: 10, b: 0, step: -1},
			want: []Day{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name: "single",
			args: args{a: 0, b: 0, step: 0},
			want: []Day{0},
		},
		{
			name: "single_ignore_b",
			args: args{a: 0, b: 10, step: 0},
			want: []Day{0},
		},
		{
			// as single_2 but as a>b then it should return b
			name: "single_a_after_b",
			args: args{a: 10, b: 0, step: 0},
			want: []Day{0},
		},
		{
			// as single but check negative-zero is the same as positive-zero
			name: "single_step_neg_zero",
			args: args{a: 0, b: 0, step: -0},
			want: []Day{0},
		},
	}

	// Used to test multiple functions with the same dataset
	const (
		forEach = iota
		slice
		endMarker
	)

	for funcId := 0; funcId < endMarker; funcId++ {
		for _, tt := range tests {

			var name string
			switch funcId {
			case forEach:
				name = "forEach_"
			case slice:
				name = "slice_"
			}

			t.Run(name+tt.name, func(t *testing.T) {
				var got []Day

				it := Iterate(tt.args.a, tt.args.b, tt.args.step)

				switch funcId {

				case forEach:
					// Call for each but append directly into got slice
					_ = it.ForEach(func(day Day) error {
						got = append(got, day)
						return nil
					})

				case slice:
					got = it.Slice()
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Iterate() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
