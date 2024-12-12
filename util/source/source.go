package source

import (
	"fmt"
	"path"
	"runtime"
)

const (
	maxStackLen = 50
)

type File struct {
	File string
	Line int
}

func (f File) String() string {
	return fmt.Sprintf("%s:%d", f.File, f.Line)
}

func SourceFile() File {
	// 1 to skip this function
	return SourceFileN(1)
}

func SourceFileN(skip int) File {
	// +1 to skip this function
	file := SourceFileFullN(skip + 1)
	file.File = path.Base(file.File)
	return file
}

func SourceFileFull() File {
	// 1 to skip this function
	return SourceFileFullN(1)
}

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
