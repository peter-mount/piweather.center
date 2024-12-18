package cloud

import (
	"github.com/peter-mount/go-anim/graph"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
)

type Filter interface {
	// Limit sets the limit to detect clouds.
	// This is a value 0...1
	//
	// 0.84 is the default.
	Limit(rbLimit float64) Filter

	// Filter returns a graph.Filter to use against a source image
	Filter() graph.Filter

	Coverage() Coverage
}

type Coverage struct {
	Cloud float64 // percentage clouds
	Sky   float64 // percentage sky
}

type filter struct {
	mask       image.Image // Mask (optional) to isolate the sky from obstructions
	coverage   float64     // cloud coverage as a percentage
	total      int         // total pixels in image
	cloud      int         // pixels detected as cloud
	sky        int         // pixels detected as the sky
	rbLimit    float64     // red/blue limit to indicate cloud
	skyColor   color.Color // colour for clear sky
	cloudColor color.Color // colour for clouds
	noColor    color.Color // colour for nothing detected or outside the mask
}

func NewFilter(mask image.Image) Filter {
	return &filter{
		mask:     mask,
		skyColor: colornames.Blue,
		noColor:  colornames.Black,
		rbLimit:  0.84,
	}
}

func (c *filter) Coverage() Coverage {
	r := Coverage{}
	if c.total > 0 {
		tf := float64(c.total)
		r.Cloud = float64(c.cloud) / tf
		r.Sky = float64(c.sky) / tf
	}
	return r
}

func (c *filter) Limit(rbLimit float64) Filter {
	c.rbLimit = rbLimit
	return c
}

func (c *filter) CloudColor(col color.Color) Filter {
	c.cloudColor = col
	return c
}

func (c *filter) SkyColor(col color.Color) Filter {
	c.skyColor = col
	return c
}

func (c *filter) Filter() graph.Filter {
	return c.filter
}

func (c *filter) filter(x, y int, col color.Color) (color.Color, error) {
	c.total++
	np := c.noColor

	if c.isVisible(x, y) {
		r, _, b, _ := col.RGBA()
		if b > 0 {
			rb := float64(r) / float64(b)
			switch {
			case rb > c.rbLimit:
				c.cloud++
				// Use cloudColor if set otherwise use the original colour.
				// The latter gives realistic results when used as an image
				if c.cloudColor != nil {
					np = c.cloudColor
				} else {
					np = col
				}

			default:
				c.sky++
				np = c.skyColor
			}
		}
	}

	return np, nil
}

func (c *filter) isVisible(x, y int) bool {
	if c.mask == nil {
		return true
	}
	r, g, b, _ := c.mask.At(x, y).RGBA()
	// Any bright pixel is true, otherwise false
	return r > 0x800 || g > 0x800 || b > 0x800
}
