package template

import (
	"fmt"
	strings2 "github.com/peter-mount/piweather.center/util/strings"
	"strconv"
	"strings"
)

// DecimalAlign provides a means to align a table column against the decimal point.
//
// For this to work you need to scan the column before you render it.
//
// Here we will parse a column of reading keys:
//
// 1  {{ $valPad := decimalAlign }}
//
// 2  {{ range $reading := $keys }}
//
// 3  {{ $reading := getReading $reading }}
//
// 4  {{ $valPad.Add $reading.Value.Float $reading.Precision }}
//
// 5  {{ end }}
//
// Line 1 creates $valPad which will be used to format the column.
// Line 4 adds a value to $valPad as well as the precision (decimal points)
//
// Then in the table header:
//
//	<th class="alignCenter" style="width: {{$valPad.Width}};">Value</th>
//
// and for each cell:
//
//	<td style="--pad: {{$valPad.Pad $reading.Value}};" class="alignDecimal">{{$reading.Value}}</td>
//
// For css:
//
//	.alignDecimal {
//	   text-align: left;
//	   transform: translateX(calc(var(--pad, 0) * 1ch));
//	}
//
// This works best with monospaced fonts.
type DecimalAlign struct {
	lp int
	rp int
}

// Add adds a value to this instance.
// v is the actual value whilst precision is the number of decimal points it will have.
// This returns "" always as templates require a result, and we do not want to write
// any erroneous text to the output.
func (d *DecimalAlign) Add(v, precision interface{}) string {
	vf, _ := strings2.ToFloat64(v)
	lp := len(fmt.Sprintf("%.0f", vf))
	if d.lp < lp {
		d.lp = lp
	}

	pf, _ := strings2.ToFloat64(precision)
	if d.rp < int(pf) {
		d.rp = int(pf)
	}

	// Must return a result as templates will complain at runtime
	// "" means nothing is output from the template
	return ""
}

// Width returns the width of the table column in ch
func (d *DecimalAlign) Width() string {
	// width is number of digits either side of value and
	// 0.5ch on either side to allow for overlap in the browser
	return fmt.Sprintf("%dch", d.lp+d.rp+1)
}

// Pad returns the value to include with the --pad css variable for a
// specific table cell.
func (d *DecimalAlign) Pad(v interface{}) string {
	var lp int
	if s, ok := v.(string); ok {
		// If a string then look for ".".
		// If found then lp=num chars before but excluding it.
		// If not found then use length of string
		lp = strings.Index(s, ".")
		if lp < 0 {
			lp = len(s)
		}
	} else {
		vf, _ := strings2.ToFloat64(v)
		lp = len(fmt.Sprintf("%.0f", vf))
	}
	return strconv.Itoa(d.lp - lp)
}

// NewDecimalAlign returns a new instance.
// It takes 0, 1 or 2 parameters.
// the first is the initial number of digits to the left of '.'
// the second is the initial number of digits to the right of '.'.
// The defaults are 1 to the left and 0 to the right.
func NewDecimalAlign(a ...interface{}) *DecimalAlign {
	// lp=1 as we always have 1 digit in a number. rp minimum is 0
	lp, rp := 1, 0

	if len(a) > 0 {
		f, _ := strings2.ToFloat64(a[0])
		if lp < int(f) {
			lp = int(f)
		}
	}

	if len(a) > 1 {
		f, _ := strings2.ToFloat64(a[1])
		if rp < int(f) {
			rp = int(f)
		}
	}

	return &DecimalAlign{lp: lp, rp: rp}
}
