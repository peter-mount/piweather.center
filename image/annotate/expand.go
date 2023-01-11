package annotate

import (
	graph "github.com/peter-mount/go-graphics"
	"image"
	"image/draw"
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

	newBounds := image.Rectangle{
		Min: oldBounds.Min,
		Max: image.Point{X: oldBounds.Max.X + left + right, Y: oldBounds.Max.Y + top + bottom},
	}

	// Rectangle to draw old image into new one
	topLeft := image.Point{X: left, Y: top}
	drawBounds := oldBounds.Add(topLeft)

	dstImage := image.NewRGBA(newBounds)
	draw.Draw(dstImage, drawBounds, srcImage, image.Point{}, draw.Src)

	return dstImage
}
