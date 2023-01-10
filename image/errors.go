package image

import "errors"

var (
	noFilename      = errors.New("no filename defined")
	unsupportedType = errors.New("unsupported image type")
)

func IsNoFilename(err error) bool { return err == noFilename }

func IsUnsupported(err error) bool { return err == unsupportedType }
