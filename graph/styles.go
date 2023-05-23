package graph

import (
	"fmt"
	svg2 "github.com/peter-mount/piweather.center/graph/svg"
	"strings"
)

const (
	StrokeBlack     = "stroke='black'"
	StrokeGrey      = "stroke='grey'"
	StrokeLightGrey = "stroke='lightgrey'"
	StrokeRed       = "stroke='red'"
	StrokeWidth1    = "stroke-width='1px'"
	FillNone        = "fill='none'"
	FillWhite       = "fill='white'"
	FillBlack       = "fill='black'"
	FillRed         = "fill='red'"
	TitleSize       = 16 // Y-axis main title size
	SubTitleSize    = 12 // Y-axis sub title size
	LabelSize       = 8  // X & Y axis tick label size
	DateSize        = 10 // Date size vertical on x-axis
)

var (
	css = strings.ReplaceAll(fmt.Sprintf(`<style type="text/css"><![CDATA[
text{font-family:sans-serif;}
.t5{font-size:5px;}
.t10{font-size:10px;}
.t20{font-size:20px;}
.leftVert{transform:"rotate(-90)";}
.graphId{font-size:16px;fill:black;fill-opacity:0.5;}
.titleX,.titleY{font-size:%dpx;text-anchor:middle;dominant-baseline:central;}
.subTitleX,.subTitleY{font-size:%dpx;text-anchor:middle;dominant-baseline:central;}
.labelX,.labelY{font-size:%dpx;text-anchor:middle;dominant-baseline:central;}
.labelDate{font-size:%dpx;fill:grey;}
.grid0{stroke:grey;stroke-width:1px}
.grid1{stroke:lightgrey;stroke-width:1px}
.txt{fill:black}
]]></style>`,
		TitleSize,    // .titleX,.titleY
		SubTitleSize, // .subTitleX,.subTitleY
		LabelSize,    // .labelX,.labelY
		DateSize,     // .labelDate
	), "\n", "")
)

// CSS writes the default stylesheet to the SVG
func CSS(s svg2.SVG) {
	s.WriteString(css)
}
