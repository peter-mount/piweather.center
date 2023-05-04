package annotate

import (
	graph "github.com/peter-mount/go-graphics"
	"image"
	"image/draw"
	"strconv"
	"strings"
)

func IMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Expand will expand an image with the given dimensions.
// If top>0 then top pixels are added to the top.
// Same for left, bottom or right.
func Expand(srcImage image.Image, top, left, bottom, right int) graph.Image {
	// Ensure parameters are >=0
	left = IMax(left, 0)
	right = IMax(right, 0)
	top = IMax(top, 0)
	bottom = IMax(bottom, 0)

	// Calculate new image size
	oldBounds := srcImage.Bounds()
	newBounds := image.Rectangle{Min: oldBounds.Min, Max: image.Point{X: oldBounds.Max.X + left + right, Y: oldBounds.Max.Y + top + bottom}}

	// Rectangle to draw old image into new one
	topLeft := image.Point{X: left, Y: top}
	drawBounds := oldBounds.Add(topLeft)

	dstImage := image.NewRGBA(newBounds)
	draw.Draw(dstImage, drawBounds, srcImage, image.Point{}, draw.Src)

	return dstImage
}

func Crop(srcImage image.Image, bounds image.Rectangle) image.Image {
	dstImage := image.NewRGBA(bounds.Sub(bounds.Min))
	draw.Draw(dstImage, dstImage.Bounds(), srcImage, bounds.Min, draw.Src)
	return dstImage
}

// ParseCoordinates parses a string into a slice of integers.
// The string will be trimmed of any spaces and ( ) pairs.
// The integers can be delimited by either space, comma or x
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
//
// To split around the first instance of a separator, see Cut.
func ParseCoordinates(s string, n int) ([]int, error) {
	var i []int
	// Slice removing any whitespace, wrapping () and changing , or x to space as delimiters
	for _, a := range strings.SplitN(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.Trim(strings.TrimSpace(s), "()"),
				",", " "),
			"x", " "),
		" ", n) {
		if a != "" {
			v, err := strconv.Atoi(a)
			if err != nil {
				return nil, err
			}
			i = append(i, v)
		}
	}

	return i, nil
}
