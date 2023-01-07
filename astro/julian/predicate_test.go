package julian

import "testing"

func TestAfter(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: true},
		{name: "equal", args: args{a: 10, b: 10}, want: false},
		{name: "before", args: args{a: 5, b: 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := After(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("After() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfterEqual(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: true},
		{name: "equal", args: args{a: 10, b: 10}, want: true},
		{name: "before", args: args{a: 5, b: 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AfterEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AfterEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBefore(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: false},
		{name: "equal", args: args{a: 10, b: 10}, want: false},
		{name: "before", args: args{a: 5, b: 10}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Before(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Before() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeforeEqual(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: false},
		{name: "equal", args: args{a: 10, b: 10}, want: true},
		{name: "before", args: args{a: 5, b: 10}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BeforeEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("BeforeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: false},
		{name: "equal", args: args{a: 10, b: 10}, want: true},
		{name: "before", args: args{a: 5, b: 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFalse(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: false},
		{name: "equal", args: args{a: 10, b: 10}, want: false},
		{name: "before", args: args{a: 5, b: 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := False(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("False() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrue(t *testing.T) {
	type args struct {
		a Day
		b Day
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "after", args: args{a: 10, b: 5}, want: true},
		{name: "equal", args: args{a: 10, b: 10}, want: true},
		{name: "before", args: args{a: 5, b: 10}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := True(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("True() = %v, want %v", got, tt.want)
			}
		})
	}
}
