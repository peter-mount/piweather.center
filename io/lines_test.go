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

func forEachTest(tests []struct {
	name string
	src  []byte
	want []string
}, test func([]byte) []string, t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := test(tt.src)

			// get string arrays
			equal := len(tt.want) == len(got)
			if equal {
				for i, s := range tt.want {
					equal = s == got[i]
				}
			}

			if !equal {
				t.Errorf(
					"want %q got %q",
					strings.Join(tt.want, "\",\""),
					strings.Join(got, "\",\""),
				)
			}
		})
	}
}

func TestReader_ForEachLine(t *testing.T) {
	forEachTest(
		[]struct {
			name string
			src  []byte
			want []string
		}{
			{
				name: "rec1",
				src:  []byte("rec1\n"),
				want: []string{"rec1"},
			},
			{
				name: "rec1 nt",
				src:  []byte("rec1"),
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
		},
		func(src []byte) []string {
			var got []string
			err := NewReader().
				ForEachLine(func(s string) error {
					got = append(got, s)
					return nil
				}).
				FromBytes(src)
			if err != nil {
				t.Error(err)
			}
			return got
		},
		t)
}

func TestReader_ForEachRecord(t *testing.T) {
	forEachTest(
		[]struct {
			name string
			src  []byte
			want []string
		}{
			// =========================
			// Test valid records
			// =========================
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
				name: "rec3",
				src:  []byte("\002rec1\003\002rec2\003\002rec3\003"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			{
				name: "rec4",
				src:  []byte("\002rec1\003\002rec2\003\002rec3\003\002rec4\003"),
				want: []string{"rec1", "rec2", "rec3", "rec4"},
			},
			// ==================================================
			// Records which should ignore data outside stx-etx
			// ==================================================
			{
				name: "rec2 prefix 1",
				src:  []byte("pref1\002rec1\003\002rec2\003"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "rec2 prefix 2",
				src:  []byte("pref1\npref2\002rec1\003\002rec2\003"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "rec2 suffix 1",
				src:  []byte("\002rec1\003\002rec2\003suff1"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "rec2 suffix 2",
				src:  []byte("\002rec1\003\002rec2\003suff1\nsuff2"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "rec2 middle 1",
				src:  []byte("\002rec1\003mid1\002rec2\003"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "rec2 middle 2",
				src:  []byte("\002rec1\003mid1\nmid2\002rec2\003"),
				want: []string{"rec1", "rec2"},
			},
			// =========================
			// Test errors
			// =========================
			{
				name: "no etx",
				src:  []byte("\002rec1"),
			},
			{
				name: "rec then no etx",
				src:  []byte("\002rec1\003\002"),
				want: []string{"rec1"},
			},
		},
		func(src []byte) []string {
			var got []string
			err := NewReader().
				ForEachRecord(func(s string) error {
					got = append(got, s)
					return nil
				}).
				FromBytes(src)
			if err != nil {
				t.Error(err)
			}
			return got
		},
		t)
}

func TestReader_ForEach_ScanStxEtxCombiRecord(t *testing.T) {
	forEachTest(
		[]struct {
			name string
			src  []byte
			want []string
		}{
			// =========================
			// Plain text
			// =========================
			{
				name: "singleline",
				src:  []byte("rec1"),
				want: []string{"rec1"},
			},
			{
				name: "multiline_2",
				src:  []byte("rec1\nrec2"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "multiline_3",
				src:  []byte("rec1\nrec2\nrec3"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			{
				name: "multiline_3_crlf",
				src:  []byte("rec1\nrec2\r\nrec3"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			{
				name: "multiline_3_trailing_lf",
				src:  []byte("rec1\nrec2\nrec3\n"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			{
				name: "multiline_3_trailing_crlf",
				src:  []byte("rec1\nrec2\nrec3\r\n"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			// =========================
			// Pure records
			// =========================
			{
				name: "singlerec",
				src:  []byte("\002rec1\003"),
				want: []string{"rec1"},
			},
			{
				name: "multirec2",
				src:  []byte("\002rec1\003\002rec2\003"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "multirec3",
				src:  []byte("\002rec1\003\002rec2\003\002rec3\003"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			// =========================
			// mixed lines and records
			// =========================
			{
				name: "rec with prefix 1",
				src:  []byte("rec1\002rec2\003"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "rec with prefix 2",
				src:  []byte("rec1\nrec2\002rec3\003"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			{
				name: "rec with suffix 1",
				src:  []byte("\002rec1\003\002rec2\003rec3"),
				want: []string{"rec1", "rec2", "rec3"},
			},
			{
				name: "rec with suffix 2",
				src:  []byte("\002rec1\003\002rec2\003rec3\nrec4"),
				want: []string{"rec1", "rec2", "rec3", "rec4"},
			},
			{
				name: "rec with embedded lf",
				src:  []byte("\002rec1\nrec2\003rec3\nrec4"),
				want: []string{"rec1\nrec2", "rec3", "rec4"},
			},
			{
				name: "rec with embedded crlf",
				src:  []byte("\002rec1\r\nrec2\003rec3\nrec4"),
				want: []string{"rec1\r\nrec2", "rec3", "rec4"},
			},
			{
				name: "rec with embedded lf and prefix",
				src:  []byte("rec1\002rec2\nrec3\003rec4"),
				want: []string{"rec1", "rec2\nrec3", "rec4"},
			},
			{
				name: "rec with embedded lf and prefix 2",
				src:  []byte("rec1\nrec2\002rec3\nrec4\003rec5"),
				want: []string{"rec1", "rec2", "rec3\nrec4", "rec5"},
			},
			// ============================================================
			// Test where the last record has no etx - it should be ignored
			// ============================================================
			{
				name: "no etx",
				src:  []byte("\002rec1"),
			},
			{
				name: "rec then no etx",
				src:  []byte("\002rec1\003\002rec2"),
				want: []string{"rec1"},
			},
			{
				name: "prefix 1 rec then no etx",
				src:  []byte("rec1\002rec2\003\002rec3"),
				want: []string{"rec1", "rec2"},
			},
			{
				name: "prefix 2 rec then no etx",
				src:  []byte("rec1\nrec2\002rec3\003\002rec4"),
				want: []string{"rec1", "rec2", "rec3"},
			},
		},
		func(src []byte) []string {
			var got []string
			err := NewReader().
				ForEach(ScanStxEtxCombiRecord, func(s string) error {
					got = append(got, s)
					return nil
				}).
				FromBytes(src)
			if err != nil {
				t.Error(err)
			}
			return got
		},
		t)
}
