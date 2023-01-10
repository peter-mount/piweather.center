package dir

import (
	"testing"
)

func Test_directory_Resolve(t *testing.T) {
	tests := []struct {
		name     string
		dirName  string
		fileName string
		want     string
	}{
		{name: "plain", dirName: "/tmp", fileName: "test.jpg", want: "/tmp/test.jpg"},
		{name: "./plain", dirName: "/tmp", fileName: "./test.jpg", want: "/tmp/test.jpg"},
		{name: "/plain", dirName: "/tmp", fileName: "/test.jpg", want: "/test.jpg"},
		{name: "/srv/plain", dirName: "/tmp", fileName: "/srv/test.jpg", want: "/srv/test.jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// need local copy otherwise pointer will be to last test
			dirName := tt.dirName
			d := &directory{
				Dir: &dirName,
			}
			got := d.Resolve(tt.fileName)
			if got != tt.want {
				t.Errorf("Split() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directory_Split(t *testing.T) {
	tests := []struct {
		name     string
		dirName  string
		fileName string
		wantDir  string
		wantFile string
	}{
		{name: "plain", dirName: "/tmp", fileName: "test.jpg", wantDir: "/tmp", wantFile: "test.jpg"},
		{name: "./plain", dirName: "/tmp", fileName: "./test.jpg", wantDir: "/tmp", wantFile: "test.jpg"},
		{name: "/plain", dirName: "/tmp", fileName: "/test.jpg", wantDir: "/", wantFile: "test.jpg"},
		{name: "/srv/plain", dirName: "/tmp", fileName: "/srv/test.jpg", wantDir: "/srv", wantFile: "test.jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// need local copy otherwise pointer will be to last test
			dirName := tt.dirName
			d := &directory{
				Dir: &dirName,
			}
			gotDir, gotFile := d.Split(tt.fileName)
			if gotDir != tt.wantDir {
				t.Errorf("Split() got = %v, want %v", gotDir, tt.wantDir)
			}
			if gotFile != tt.wantFile {
				t.Errorf("Split() got = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
