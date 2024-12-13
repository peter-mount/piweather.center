package source

import (
	"fmt"
	"path"
	"runtime"
)

const (
	maxStackLen = 50
)

// File representing a location within a source file
type File struct {
	// File name of the .go file
	File string
	// Line number within File
	Line int
}

func (f File) String() string {
	return fmt.Sprintf("%s:%d", f.File, f.Line)
}

// SourceFile returns a File representing the line which called this function.
// The returned File will be just the file name of the .go file, not the full path.
func SourceFile() File {
	// 1 to skip this function
	return SourceFileN(1)
}

// SourceFileN returns a File representing the line which called this function.
//
// skip is the number of calls to skip to get the required result
//
// The returned File will be just the file name of the .go file, not the full path.
func SourceFileN(skip int) File {
	// +1 to skip this function
	file := SourceFileFullN(skip + 1)
	file.File = path.Base(file.File)
	return file
}

// SourceFileFull returns a File representing the line which called this function.
// The returned File will be the full file name including the path to that file.
func SourceFileFull() File {
	// 1 to skip this function
	return SourceFileFullN(1)
}

// SourceFileFullN returns a File representing the line which called this function.
//
// skip is the number of calls to skip to get the required result
//
// The returned File will be the full file name including the path to that file.
func SourceFileFullN(skip int) File {
	if skip < 0 {
		skip = 0
	}

	var pc [maxStackLen]uintptr
	// Skip two extra frames to account for this function
	// and runtime.Callers itself.
	n := runtime.Callers(skip+2, pc[:])
	if n == 0 {
		panic("testing: zero callers found")
	}

	frames := runtime.CallersFrames(pc[:n])
	var frame runtime.Frame
	for more := true; more; {
		frame, more = frames.Next()
		if frame.Function == "runtime.gopanic" {
			continue
		}
		return File{File: frame.File, Line: frame.Line}
	}

	return File{}
}
