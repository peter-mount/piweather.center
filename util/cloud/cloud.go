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

	// GreenLimit sets the limit when red dominates blue.
	// If green is below this level then we do not detect that pixel as cloud.
	//
	// This can detect some obstructions like tree foliage which would otherwise get detected as cloud.
	//
	// The default is 50000
	GreenLimit(greenLimit uint32) Filter

	// Filter returns a graph.Filter to use against a source image
	Filter() graph.Filter

	Coverage() Coverage
}

type Coverage struct {
	Cloud    float64 // percentage clouds
	Sky      float64 // percentage sky
	Obscured float64 // percentage obscured
}

type filter struct {
	privMask      image.Image // Privacy Mask (optional) to isolate areas not to be shown
	mask          image.Image // Mask (optional) to isolate the sky from obstructions
	coverage      float64     // cloud coverage as a percentage
	total         int         // total pixels in image
	cloud         int         // pixels detected as cloud
	sky           int         // pixels detected as the sky
	obscured      int         // pixels detected as obscured
	rbLimit       float64     // red/blue limit to indicate cloud
	greenLimit    uint32      // green limit for when (r/b)>1. If below this limit we do not detect cloud
	skyColor      color.Color // colour for clear sky
	cloudColor    color.Color // colour for clouds, default original if nil
	obscuredColor color.Color // colour for when obscured, default original if nil
	noColor       color.Color // colour for nothing detected or outside the mask
}

func NewFilter(privMask, mask image.Image) Filter {
	return &filter{
		privMask:   privMask,
		mask:       mask,
		skyColor:   colornames.Blue,
		noColor:    colornames.Black,
		rbLimit:    0.84,
		greenLimit: 50000,
	}
}

func (c *filter) Coverage() Coverage {
	r := Coverage{}
	if c.total > 0 {
		f := 100.0 / float64(c.total)
		r.Cloud = f * float64(c.cloud)
		r.Sky = f * float64(c.sky)
		r.Obscured = f * float64(c.obscured)
	}
	return r
}

func (c *filter) Limit(rbLimit float64) Filter {
	c.rbLimit = rbLimit
	return c
}

func (c *filter) GreenLimit(greenLimit uint32) Filter {
	c.greenLimit = greenLimit
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
	np := c.noColor

	if c.isSky(x, y) {
		// Count only sky pixels in the total
		c.total++

		r, g, b, _ := col.RGBA()

		rb := 0.0
		if b > 0 {
			rb = float64(r) / float64(b)
		}

		switch {
		// This handles when red dominates blue but green is not as high.
		// This can detect some obstructions like tree foliage which would
		// otherwise get detected as clouds
		case rb > 1 && g < 50000:
			c.obscured++

			if c.obscuredColor != nil {
				np = c.obscuredColor
			} else {
				np = col
			}

		// Detect cloud
		case rb > c.rbLimit:
			c.cloud++

			if c.cloudColor != nil {
				np = c.cloudColor
			} else {
				np = col
			}

		// Default is sky
		default:
			c.sky++
			np = c.skyColor
		}
	} else if c.isVisible(x, y) {
		// If not sky but visible then keep the original colour
		// This then shows the rest of the image
		np = col
	}

	return np, nil
}

func (c *filter) isVisible(x, y int) bool {
	// If privacyMask is available then stop if it's masked out
	if c.privMask != nil && graph.IsBlack(c.privMask.At(x, y)) {
		return false
	}
	return true
}

func (c *filter) isSky(x, y int) bool {
	if c.mask != nil {
		return graph.IsNotBlack(c.mask.At(x, y))
	}

	return true
}
