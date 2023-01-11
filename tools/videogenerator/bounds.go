package videogenerator

import (
	"image"
)

type Bounds struct {
	X      int     `yaml:"x"`
	Y      int     `yaml:"y"`
	Scale  float64 `yaml:"scale"`  // Scale image
	Width  int     `yaml:"width"`  // Width of output
	Height int     `yaml:"height"` // Height of output
}

func (b Bounds) Rect(b2 image.Rectangle) image.Rectangle {
	// Allow scaling with fixed width
	if b.Height == 0 && b.Width > 0 && b.Scale == 0.0 {
		b.Scale = float64(b.Width) / float64(b2.Dx())
	}

	// Allow scaling with fixed height
	if b.Height > 0 && b.Width == 0 && b.Scale == 0.0 {
		b.Scale = float64(b.Height) / float64(b2.Dy())
	}

	// Scale from original image
	if b.Scale != 0.0 {
		b.Width = int(float64(b2.Dx()) * b.Scale)
		b.Height = int(float64(b2.Dy()) * b.Scale)
	}

	// Failsafe if still no size
	if b.Width == 0 {
		b.Width = b2.Dx()
	}
	if b.Height == 0 {
		b.Height = b2.Dy()
	}

	return image.Rectangle{
		Min: image.Point{X: b.X, Y: b.Y},
		Max: image.Point{X: b.X + b.Width - 1, Y: b.Y + b.Height - 1},
	}
}
