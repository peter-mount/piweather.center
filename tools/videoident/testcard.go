package videointro

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

// TestCard adds a Test Card layer to the frame, loosely based on the BBC <HD> Test Card 1080p
func (v *VideoIntro) TestCard(frame int) {
	frate := float64(v.Config.Format.FrameRate)
	sec := float64(v.Config.Start) - (float64(frame) / frate)
	_, secF := math.Modf(sec)
	secF *= float64(v.Config.Format.FrameRate)

	gc := v.gc

	w, h := float64(v.Config.Format.Width), float64(v.Config.Format.Height)
	w2, h2 := w/2, h/2
	// 7 squares, triangle, 7 squares along the border
	dW, dH := w/30.0, h/27.0
	dW2, dHd2, dH2, dH23 := dW*2, dH/2, dH*2, dH*2/3

	// Fill in background
	fillRectangle(gc, 0, 0, w, h, colornames.Grey)

	// Grid lines
	dx, dy := gridLines(gc, w, h, dW, dH, dW2, dH2)

	// frame slider border
	frameSlider(gc, true, frate, secF, dW, dH, dx, dy)

	// Black lines
	blackLineDecorations(gc, w, h, dW, dH, dx, dy)

	// frame slider
	frameSlider(gc, false, frate, secF, dW, dH, dx, dy)

	// corner gradients
	cornerCalibrationMarks(gc, w, h, dW, dH, dx, dy)

	// Clear bottom border
	fillRectangle(gc, 0, h-dH, w, dH, color.Black)

	cornerWhiteSquares(gc, w, h, dW, dH, dW2)

	// Top border
	topBorder(gc, w, dx, dW2, dH, colornames.Yellow, colornames.Cyan, colornames.Green, colornames.Magenta, colornames.Red, colornames.Blue, color.Black)

	// Left & Right borders
	sideBorders(gc, 0, w, h, dW, dH, dH2, colornames.Lightgreen, colornames.Red, color.Black, colornames.Lightblue, colornames.Darkgreen)
	sideBorders(gc, 1, w, h, dW, dH, dH2, colornames.Yellow, colornames.Cyan, colornames.Green, colornames.Magenta, colornames.Red, colornames.Blue, color.Black)

	bottomBorders(gc, w2, h, dW, dH, dHd2)

	// Inner border
	gc.SetStrokeColor(color.White)
	gc.SetLineWidth(10)
	gc.BeginPath()
	rectangle(gc, dW, dH, w-(dW*2), h-(dH*2))
	gc.Stroke()

	fillPoly(gc, color.White, w2, 0, (w-dH)/2, dH, (w+dH)/2, dH)
	fillPoly(gc, color.White, w2, h-1, (w-dH)/2, h-dH, (w+dH)/2, h-dH)
	fillPoly(gc, color.White, 0, h2, dW, h2-dH23, dW2, h2, dW, h2+dH23)

}

func cornerWhiteSquares(gc *draw2dimg.GraphicContext, w, h, dW, dH, dW2 float64) {
	// White squares on each corner
	fillRectangle(gc, 0, 0, dW2, dH, color.White)
	fillRectangle(gc, w-dW2, 0, dW2, dH, color.White)
	fillRectangle(gc, w-dW, h-dH, dW, dH, color.White)
	//fillRectangle(gc, 0, h-dH, dW, dH, color.White)
}

func gridLines(gc *draw2dimg.GraphicContext, w, h, dW, dH, dW2, dH2 float64) (float64, float64) {
	gc.SetStrokeColor(color.White)
	gc.SetLineWidth(10)
	gc.BeginPath()
	rectangle(gc, dW, dH, w-(dW*2), h-(dH*2))
	gc.Stroke()

	gc.BeginPath()

	y := dH
	dy := (h - dH2) / 8
	for i := 0; i < 8; i++ {
		gc.MoveTo(0, y)
		gc.LineTo(w, y)
		y = y + dy
	}

	x := dW
	dx := (w - dW2) / 15
	for i := 0; i < 15; i++ {
		y1 := h
		if i > 3 && i < 12 {
			y1 = y - dy
		}
		gc.MoveTo(x, 0)
		gc.LineTo(x, y1)
		x = x + dx
	}

	gc.Stroke()

	return dx, dy
}

func topBorder(gc *draw2dimg.GraphicContext, w, dx, dW2, dH float64, cols ...color.Color) {
	x := dW2
	dx = (w - x - x) / float64(len(cols))
	for _, c := range cols {
		x, _ = fillRectangle(gc, x, 0, dx, dH, c)
	}
}

func sideBorders(gc *draw2dimg.GraphicContext, side, w, h, dW, dH, dH2 float64, cols ...color.Color) {
	y := dH
	dy := (h - dH2) / float64(len(cols))
	for _, col := range cols {
		_, y = fillRectangle(gc, side*(w-dW), y, dW, dy, col)
	}
}

func bottomBorders(gc *draw2dimg.GraphicContext, w2, h, dW, dH, dHd2 float64) {
	// Bottom left border
	x := 0.0
	dx := (w2 - dW) / 256
	for i := 255; i > 0; i-- {
		_, _ = fillRectangle(gc, x, h-dH, dx, dHd2, color.Gray{Y: uint8(32 + ((255 - i) >> 1))})
		x, _ = fillRectangle(gc, x, h-dHd2, dx, dHd2, color.Gray{Y: uint8(i)})
	}

	// Bottom right border
	x = 0
	dx = (w2/2 - dW) / 256
	for i := 255; i > 0; i-- {
		fillRectangle(gc, w2+x, h-dH, dx, dHd2, color.Gray{Y: uint8(32 + ((255 - i) >> 1))})
		fillRectangle(gc, x, h-dHd2, dx, dHd2, color.Gray{Y: uint8(i)})
	}
}

func blackLineDecorations(gc *draw2dimg.GraphicContext, w, h, dW, dH, dx, dy float64) {
	dd := 7.0
	gc.SetStrokeColor(color.Black)
	gc.SetLineWidth(dd)
	gc.BeginPath()

	// Black lines top row
	relLine(gc,
		dW, dH+(2*dy)-dd,
		(3*dx)-dd, 0,
		0, dd-(2*dy))

	a := func(l, r float64) {
		relLine(gc,
			dW+(l*dx)+dd, dH,
			0, (2*dy)-dd,
			(r*dx)-dd-dd, 0,
			0, dd-(2*dy))
	}
	a(3, 4)
	a(7, 1)
	a(8, 4)

	relLine(gc,
		dW+(12*dx)+dd, dH,
		0, (2*dy)-dd,
		3*dx+dd, 0)

	// Bottom row
	relLine(gc,
		dW, h-dH-(2*dy)+dd,
		(3*dx)-dd, 0,
		0, 2*dy)

	relLine(gc,
		dW+(3*dx)+dd, h-dH+dd+dd,
		0, (-2*dy)-dd,
		(4*dx)-dd-dd, 0,
		0, dy-dd-dd)

	relLine(gc,
		dW+(7*dx)+dd, h-dH-dy-dd,
		0, (-dy)+dd+dd,
		(dx)-dd-dd, 0,
		0, dy-dd)

	relLine(gc,
		dW+(8*dx)+dd, h-dH-dy-dd,
		0, (-dy)+dd+dd,
		(4*dx)-dd-dd, 0,
		0, (2*dy)-dd)

	relLine(gc,
		dW+(12*dx)+dd, h-dH-dd,
		0, (-2*dy)+dd+dd,
		(3*dx)-dd-dd, 0)

	// side black lines
	relLine(gc,
		dW, dH+(2*dy)+dd,
		(3*dx)-dd, 0,
		0, (2*dy)-dd-dd,
		-dx+dd+dd, 0,
		0, -dy/2)

	relLine(gc,
		dW, h-(dH+(2*dy)+dd),
		(3*dx)-dd, 0,
		0, -((2 * dy) - dd - dd),
		-dx+dd+dd, 0,
		0, dy/2)

	relLine(gc,
		w-dW, dH+(2*dy)+dd,
		-((3 * dx) - dd), 0,
		0, (2*dy)-dd-dd,
		-(-dx + dd + dd), 0,
		0, -dy/2)

	relLine(gc,
		w-dW, h-(dH+(2*dy)+dd),
		-((3 * dx) - dd), 0,
		0, -((2 * dy) - dd - dd),
		-(-dx + dd + dd), 0,
		0, dy/2)

	// Draw black lines
	gc.Stroke()
}

func cornerCalibrationMarks(gc *draw2dimg.GraphicContext, w, h, dW, dH, dx, dy float64) {
	for sy := 0; sy < 2; sy++ {
		for sx := 0; sx < 2; sx++ {
			cornerCalibrationMark(gc, sx, sy, w, h, dW, dH, dx, dy)
		}
	}
}

func cornerCalibrationMark(gc *draw2dimg.GraphicContext, sx, sy int, w, h, dW, dH, dx, dy float64) {
	gc.Save()
	defer gc.Restore()

	var tx, ty, rot float64
	switch sx + (sy << 1) {
	case 0:
		tx, ty, rot = dW, -dH, 45.0
	case 1:
		tx, ty, rot = w+dW*.25+20, dH+10, 135
	case 2:
		tx, ty, rot = dW-dx*.75, h-dH-15, -45
	case 3:
		tx, ty, rot = w+dW*.25-dx*.7+10, h+dH, -135
	}
	gc.Translate(tx, ty)
	gc.Rotate(rot * ToRad)

	bw, bh := 2.45*dx, dy
	fillRectangle(gc, 0, 0, bw, bh, color.White)

	gc.BeginPath()
	gc.SetStrokeColor(color.Black)
	gc.SetLineWidth(10)
	for i := 0.1; i < 0.9; i += 0.1 {
		relLine(gc, 0, dy*i, bw-(dy/10), 0)
	}
	gc.Stroke()
}

func frameSlider(gc *draw2dimg.GraphicContext, background bool, frate, secF, dW, dH, dx, dy float64) {
	if background {
		fillRectangle(gc,
			dW+(5.8*dx), dH+(.6*dy),
			3.4*dx, dy*.8,
			color.White)
	} else {
		fillRectangle(gc,
			dW+(6.4*dx), dH+(.8*dy),
			2.2*dx, dy*.4,
			color.Black)

		fillRectangle(gc,
			dW+(6.4*dx)+10, dH+(.8*dy+10),
			(2.2*dx*(frate-secF)/frate)-20, dy*.4-20,
			colornames.Lightgreen)
	}
}

func fillPoly(gc *draw2dimg.GraphicContext, c color.Color, v ...float64) {
	gc.SetFillColor(c)
	gc.BeginPath()
	for i := 0; i < len(v); i += 2 {
		if i == 0 {
			gc.MoveTo(v[0], v[1])
		} else {
			gc.LineTo(v[i], v[i+1])
		}
	}
	gc.Close()
	gc.Fill()
}

func fillPolyRel(gc *draw2dimg.GraphicContext, c color.Color, v ...float64) {
	gc.SetFillColor(c)
	gc.BeginPath()
	var x, y float64
	for i := 0; i < len(v); i += 2 {
		x, y = x+v[i], y+v[i+1]
		if i == 0 {
			gc.MoveTo(x, y)
		} else {
			gc.LineTo(x, y)
		}
	}
	gc.Close()
	gc.Fill()
}

func fillRectangle(gc *draw2dimg.GraphicContext, x, y, w, h float64, c color.Color) (float64, float64) {
	gc.SetFillColor(c)
	gc.BeginPath()
	rectangle(gc, x, y, w, h)
	gc.Fill()
	return x + w, y + h
}

func rectangle(gc *draw2dimg.GraphicContext, x, y, w, h float64) {
	gc.MoveTo(x, y)
	gc.LineTo(x+w, y)
	gc.LineTo(x+w, y+h)
	gc.LineTo(x, y+h)
	gc.Close()
}

func relLine(gc *draw2dimg.GraphicContext, x, y float64, v ...float64) {
	gc.MoveTo(x, y)
	for i := 0; i < len(v); i += 2 {
		x += v[i]
		y += v[i+1]
		gc.LineTo(x, y)
	}
}
