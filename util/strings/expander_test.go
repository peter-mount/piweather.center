package strings

import (
	"reflect"
	"testing"
)

func TestExpand(t *testing.T) {
	m := map[string]string{
		"name": "a test",
	}

	expander := func(s string) string {
		r, e := m[s]
		if e {
			return r
		}
		return "??no-match??"
	}

	type args struct {
		s string
		e Expander
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple",
			args: args{
				s: "This is a test of ${name} expander",
				e: expander,
			},
			want: "This is a test of a test expander",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Expand(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Expand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpansions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Simple",
			args: args{
				s: "This is a test of ${name} expander",
			},
			want: []string{"name"},
		},
		{
			name: "Multiple",
			args: args{
				s: "This is a test of ${name} expander which gives us ${multiple} entries",
			},
			want: []string{"name", "multiple"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Expansions(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expansions() = %v, want %v", got, tt.want)
			}
		})
	}
}
