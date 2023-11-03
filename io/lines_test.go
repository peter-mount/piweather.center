package io

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriter_WriteLines(t *testing.T) {
	tests := []struct {
		name    string
		src     []string
		want    []byte
		wantErr bool
	}{
		{
			name: "line1",
			src:  []string{"line1"},
			want: []byte("line1\n"),
		},
		{
			name: "line2",
			src:  []string{"line1", "line2"},
			want: []byte("line1\nline2\n"),
		},
		{
			name: "line3",
			src:  []string{"line1", "line2", "line3"},
			want: []byte("line1\nline2\nline3\n"),
		},
		{
			name: "line4",
			src:  []string{"line1", "line2", "line3"},
			want: []byte("line1\nline2\nline3\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewWriter().
				WriteLines(func(w LineWriter) error {
					for _, s := range tt.src {
						if err := w(s); err != nil {
							return err
						}
					}
					return nil
				}).
				CreateBytes()
			switch {

			case tt.wantErr && err != nil:
				// pass, do not do any other tests

			case !tt.wantErr && err != nil:
				t.Errorf("WriteLines() error = %v, wantErr %v", err, tt.wantErr)

			case bytes.Compare(tt.want, b) != 0:
				t.Errorf("WriteLines() want %X got %X", tt.want, b)

			default:
				// pass
			}
		})
	}
}

func TestWriter_WriteRecords(t *testing.T) {
	tests := []struct {
		name    string
		src     []string
		want    []byte
		wantErr bool
	}{
		{
			name: "record1",
			src:  []string{"record1"},
			want: []byte("\002record1\003"),
		},
		{
			name: "record2",
			src:  []string{"record1", "rec2"},
			want: []byte("\002record1\003\002rec2\003"),
		},
		{
			name: "record3",
			src:  []string{"record1", "rec2", "entry3"},
			want: []byte("\002record1\003\002rec2\003\002entry3\003"),
		},
		{
			name: "record4",
			src:  []string{"record1", "rec2", "rec3", "end4"},
			want: []byte("\002record1\003\002rec2\003\002rec3\u0003\u0002end4\u0003"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewWriter().
				WriteRecords(func(w LineWriter) error {
					for _, s := range tt.src {
						if err := w(s); err != nil {
							return err
						}
					}
					return nil
				}).
				CreateBytes()
			switch {

			case tt.wantErr && err != nil:
				// pass, do not do any other tests

			case !tt.wantErr && err != nil:
				t.Errorf("WriteLines() error = %v, wantErr %v", err, tt.wantErr)

			case bytes.Compare(tt.want, b) != 0:
				t.Errorf("WriteLines() want %X got %X", tt.want, b)

			default:
				// pass
			}
		})
	}
}

func TestReader_ForEachLine(t *testing.T) {
	tests := []struct {
		name    string
		src     []byte
		want    []string
		wantErr bool
	}{
		{
			name: "rec1",
			src:  []byte("rec1\n"),
			want: []string{"rec1"},
		},
		{
			name: "rec2",
			src:  []byte("rec1\nrec2\n"),
			want: []string{"rec1", "rec2"},
		},
		{
			name: "rec2 nt",
			src:  []byte("rec1\nrec2"),
			want: []string{"rec1", "rec2"},
		},
		{
			name: "rec3",
			src:  []byte("rec1\nrec2\nrec3\n"),
			want: []string{"rec1", "rec2", "rec3"},
		},
		{
			name: "rec3 nt",
			src:  []byte("rec1\nrec2\nrec3"),
			want: []string{"rec1", "rec2", "rec3"},
		},
		{
			name: "rec4",
			src:  []byte("rec1\nrec2\nrec3\nrec4\n"),
			want: []string{"rec1", "rec2", "rec3", "rec4"},
		},
		{
			name: "rec4 nt",
			src:  []byte("rec1\nrec2\nrec3\nrec4"),
			want: []string{"rec1", "rec2", "rec3", "rec4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			err := NewReader().
				ForEachLine(func(s string) error {
					got = append(got, s)
					return nil
				}).
				FromBytes(tt.src)

			// get string arrays
			equal := len(tt.want) == len(got)
			if equal {
				for i, s := range tt.want {
					equal = s == got[i]
				}
			}

			switch {

			case tt.wantErr && err != nil:
				// pass, do not do any other tests

			case !tt.wantErr && err != nil:
				t.Errorf("WriteLines() error = %v, wantErr %v", err, tt.wantErr)

			case !equal:
				t.Errorf(
					"WriteLines() want %q got %q",
					strings.Join(tt.want, "\\n"),
					strings.Join(got, "\\n"),
				)

			default:
				// pass
			}
		})
	}
}

func TestReader_ForEachRecord(t *testing.T) {
	tests := []struct {
		name    string
		src     []byte
		want    []string
		wantErr bool
	}{
		{
			name: "rec1",
			src:  []byte("\002rec1\003"),
			want: []string{"rec1"},
		},
		{
			name: "rec2",
			src:  []byte("\002rec1\003\002rec2\003"),
			want: []string{"rec1", "rec2"},
		},
		{
			name: "rec2 nt",
			src:  []byte("\002rec1\003\002rec2\003"),
			want: []string{"rec1", "rec2"},
		},
		{
			name: "rec3",
			src:  []byte("\002rec1\003\002rec2\003\002rec3\003"),
			want: []string{"rec1", "rec2", "rec3"},
		},
		{
			name: "rec3 nt",
			src:  []byte("\002rec1\003\002rec2\003\002rec3\003"),
			want: []string{"rec1", "rec2", "rec3"},
		},
		{
			name: "rec4",
			src:  []byte("\002rec1\003\002rec2\003\002rec3\003\002rec4\003"),
			want: []string{"rec1", "rec2", "rec3", "rec4"},
		},
		{
			name: "rec4 nt",
			src:  []byte("\002rec1\003\002rec2\003\002rec3\003\002rec4\003"),
			want: []string{"rec1", "rec2", "rec3", "rec4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			err := NewReader().
				ForEachRecord(func(s string) error {
					got = append(got, s)
					return nil
				}).
				FromBytes(tt.src)

			// get string arrays
			equal := len(tt.want) == len(got)
			if equal {
				for i, s := range tt.want {
					equal = s == got[i]
				}
			}

			switch {

			case tt.wantErr && err != nil:
				// pass, do not do any other tests

			case !tt.wantErr && err != nil:
				t.Errorf("WriteLines() error = %v, wantErr %v", err, tt.wantErr)

			case !equal:
				t.Errorf(
					"WriteLines() want %q got %q",
					strings.Join(tt.want, "\\n"),
					strings.Join(got, "\\n"),
				)

			default:
				// pass
			}
		})
	}
}
