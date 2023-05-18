package template

import (
	"fmt"
	"github.com/peter-mount/piweather.center/util"
)

func valLeftPad(v interface{}) float64 {
	f, _ := util.ToFloat64(v)
	lp := len(fmt.Sprintf("%.0f", f))
	return float64(lp)
}

// Width of a decimal column in ch
func decimalWidth(w interface{}) string {
	wf, _ := util.ToFloat64(w)
	return fmt.Sprintf("%.0fch", wf)
}

// The max width in ch a decimal column will use.
//
// lp = number of characters to left of decimal point
// rp = number of decimal places
func decimalWidthMax(lp, rp float64) float64 {
	// 2 = the '.' and 0.5ch on either side to allow for overlap in the browser
	return 2.0 + lp + rp
}

// Pad a decimal column to align the decimal place.
//
// vp = max number of characters to the left of decimal point
// v = number to be aligned, can be float64, int etc
// style="--pad: {{decimalPad $leftPad $reading.Value.Float}};"
func decimalPad(lp, v interface{}) string {
	lpf, _ := util.ToFloat64(lp)
	vf := valLeftPad(v)
	return fmt.Sprintf("%.0f", lpf-vf)
}
