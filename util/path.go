package util

import (
	"path"
	"strings"
)

// BaseName splits the filename and file type from a path.
// if the file does not contain a '.' then this returns file,""
func BaseName(fileName string) (string, string) {
	n := path.Base(fileName)

	i := strings.LastIndex(n, ".")
	if i > 0 {
		return n[:i], n[i+1:]
	}

	return n, ""
}
