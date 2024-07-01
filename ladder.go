package main

/**
generate line and ladder page
*/

import "image/color"

func drawLadder(x, y float64) {
	//fmt.Println(x, y)

	for i := 1; i <= 2; i = i + 2 { // y drop
		ym := float64(i) //+ Opt.spacing

		createGC()
		Opt.gc.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.gc.SetLineWidth(0.1)
		Opt.gc.BeginPath()
		bx := x
		by := y * ym
		Opt.gc.MoveTo(bx, by)                              // ul
		Opt.gc.LineTo(bx+Opt.spacing, by)                  // ur
		Opt.gc.LineTo(bx+Opt.spacing, (by+Opt.spacing)*ym) // lr
		Opt.gc.LineTo(bx, (by+Opt.spacing)*ym)             // lr
		Opt.gc.Close()
		Opt.gc.FillStroke()

		createGC()
		Opt.gc.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.gc.SetLineWidth(0.1)
		Opt.gc.BeginPath()
		bx = x + Opt.spacing
		by = (y + Opt.spacing) * ym
		Opt.gc.MoveTo(bx, by)                         // ul
		Opt.gc.LineTo(bx+Opt.spacing, by)             // ur
		Opt.gc.LineTo(bx+Opt.spacing, by+Opt.spacing) // lr
		Opt.gc.LineTo(bx, by+Opt.spacing)             // lr
		Opt.gc.Close()
		Opt.gc.FillStroke()

	}

}

func drawLadderLines(yPos float64) {

	y := yPos
	// everything is in units of spacing

	// top ascender
	drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)

	// mid ascender
	drawLine(point{Opt.pageMarginLeft, y + Opt.spacing}, point{Opt.pageMarginRight, y + Opt.spacing}, Opt.lineWidth, Opt.lightGray)
	drawLadder(Opt.pageMarginLeft, y)
	y = y + (Opt.spacing * 2)

	// top x-height
	drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)

	// 1 down from x-height
	drawLine(point{Opt.pageMarginLeft, y + Opt.spacing}, point{Opt.pageMarginRight, y + Opt.spacing}, Opt.lineWidth, Opt.lightGray)
	drawLine(point{Opt.pageMarginLeft, y + Opt.spacing*2}, point{Opt.pageMarginRight, y + Opt.spacing*2}, Opt.lineWidth, Opt.lightGray)
	drawLadder(Opt.pageMarginLeft, y)

	drawLine(point{Opt.pageMarginLeft, y + Opt.spacing}, point{Opt.pageMarginRight, y + Opt.spacing}, Opt.lineWidth, Opt.lightGray)
	y = y + (Opt.spacing * 2)
	drawLine(point{Opt.pageMarginLeft, y + Opt.spacing}, point{Opt.pageMarginRight, y + Opt.spacing}, Opt.lineWidth, Opt.lightGray)
	drawLadder(Opt.pageMarginLeft, y)
	y = y + (Opt.spacing * 2)

	drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
	drawLine(point{Opt.pageMarginLeft, y + Opt.spacing}, point{Opt.pageMarginRight, y + Opt.spacing}, Opt.lineWidth, Opt.lightGray)
	drawLadder(Opt.pageMarginLeft, y)
	y = y + (Opt.spacing * 2)

	drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
}

func drawLadderLineGroup() {

	for y := Opt.pageMarginTop; y <= Opt.pageMarginBottom; y += (Opt.spacing * 12) {
		if y+(Opt.spacing*8) > Opt.pageMarginBottom {
			return
		}
		drawLadderLines(y)
	}

}

// func drawLadderBars() {

// 	for yMaster := Opt.pageMarginTop; yMaster <= Opt.pageMarginBottom; yMaster += (Opt.spacing * 8) {

// 		var racount int = 0

// 		for y := yMaster; y <= yMaster+(Opt.spacing*9); y += Opt.spacing {

// 			switch racount {
// 			case 0:
// 				drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
// 				drawLadder(Opt.pageMarginLeft, y)
// 				/*
// 					Opt.gc.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
// 					Opt.gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})

// 					Opt.gc.BeginPath()
// 					Opt.gc.MoveTo(Opt.pageMarginLeft, y)
// 					Opt.gc.LineTo(Opt.pageMarginLeft+4.5, y)
// 					Opt.gc.LineTo(Opt.pageMarginLeft+4.5, y+4.5)
// 					Opt.gc.LineTo(Opt.pageMarginLeft, y+4.5)
// 					Opt.gc.Close()
// 					Opt.gc.FillStroke()
// 				*/
// 				racount++
// 			case 1:
// 				racount++
// 			case 2:
// 				drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
// 				drawLadder(Opt.pageMarginLeft, y)
// 				racount++
// 			case 3:
// 				racount++
// 			case 4:
// 				drawLadder(Opt.pageMarginLeft, y)
// 				racount++
// 			case 5:
// 				racount++
// 			case 6:
// 				drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
// 				drawLadder(Opt.pageMarginLeft, y)
// 				racount++
// 			case 7:
// 				racount++
// 			case 8:
// 				drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
// 				//drawLadder(Opt.pageMarginLeft, y)
// 				//y = y + (Opt.spacing * 3)
// 				racount = 0
// 				continue

// 			}

// 		}
// 		yMaster = yMaster + (4 * Opt.spacing)
// 	}
// }
