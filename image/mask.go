package image

import (
	"github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/graphics"
	image2 "image"
	"image/draw"
)

// Overlay one image over this one, returning a new composite Image.
func (i *Image) Overlay(overlay *Image) *Image {
	if overlay == nil {
		return i
	}

	g := graphics.New(graph.DuplicateImage(i.Image))
	g.DrawImage(overlay.Bounds(), overlay.Image, image2.Point{}, draw.Over)

	return i.Duplicate(g.Image())
}
