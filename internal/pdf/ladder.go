package pdf

/**
generate line and ladder page
*/

import (
	"image/color"
	"mynotebook/internal/common"
)

func drawLadder(x, y float64) {
	//fmt.Println(x, y)

	for i := 1; i <= 2; i = i + 2 { // y drop
		ym := float64(i) //+ common.Opt.Spacing

		CreateGC()
		common.Opt.GC.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		common.Opt.GC.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		common.Opt.GC.SetLineWidth(0.1)
		common.Opt.GC.BeginPath()
		bx := x
		by := y * ym
		common.Opt.GC.MoveTo(bx, by)                                            // ul
		common.Opt.GC.LineTo(bx+common.Opt.Spacing, by)                         // ur
		common.Opt.GC.LineTo(bx+common.Opt.Spacing, (by+common.Opt.Spacing)*ym) // lr
		common.Opt.GC.LineTo(bx, (by+common.Opt.Spacing)*ym)                    // lr
		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		CreateGC()
		common.Opt.GC.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		common.Opt.GC.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		common.Opt.GC.SetLineWidth(0.1)
		common.Opt.GC.BeginPath()
		bx = x + common.Opt.Spacing
		by = (y + common.Opt.Spacing) * ym
		common.Opt.GC.MoveTo(bx, by)                                       // ul
		common.Opt.GC.LineTo(bx+common.Opt.Spacing, by)                    // ur
		common.Opt.GC.LineTo(bx+common.Opt.Spacing, by+common.Opt.Spacing) // lr
		common.Opt.GC.LineTo(bx, by+common.Opt.Spacing)                    // lr
		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

	}

}

func drawLadderLines(yPos float64) {

	y := yPos
	// everything is in units of spacing

	// top ascender
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y}, common.Point{X: common.Opt.PageMarginRight, Y: y}, common.Opt.LineWidth, common.Opt.DarkBlack)

	// mid ascender
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y + common.Opt.Spacing}, common.Point{X: common.Opt.PageMarginRight, Y: y + common.Opt.Spacing},
		common.Opt.LineWidth,
		common.Opt.LightGray)
	drawLadder(common.Opt.PageMarginLeft, y)
	y = y + (common.Opt.Spacing * 2)

	// top x-height
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y}, common.Point{X: common.Opt.PageMarginRight, Y: y}, common.Opt.LineWidth, common.Opt.DarkBlack)

	// 1 down from x-height
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y + common.Opt.Spacing}, common.Point{X: common.Opt.PageMarginRight, Y: y + common.Opt.Spacing}, common.Opt.LineWidth, common.Opt.LightGray)
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y + common.Opt.Spacing*2}, common.Point{X: common.Opt.PageMarginRight, Y: y + common.Opt.Spacing*2}, common.Opt.LineWidth, common.Opt.LightGray)
	drawLadder(common.Opt.PageMarginLeft, y)

	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y + common.Opt.Spacing}, common.Point{X: common.Opt.PageMarginRight, Y: y + common.Opt.Spacing}, common.Opt.LineWidth, common.Opt.LightGray)
	y = y + (common.Opt.Spacing * 2)
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y + common.Opt.Spacing}, common.Point{X: common.Opt.PageMarginRight, Y: y + common.Opt.Spacing}, common.Opt.LineWidth, common.Opt.LightGray)
	drawLadder(common.Opt.PageMarginLeft, y)
	y = y + (common.Opt.Spacing * 2)

	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y}, common.Point{X: common.Opt.PageMarginRight, Y: y}, common.Opt.LineWidth, common.Opt.DarkBlack)
	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y + common.Opt.Spacing}, common.Point{X: common.Opt.PageMarginRight, Y: y + common.Opt.Spacing}, common.Opt.LineWidth, common.Opt.LightGray)
	drawLadder(common.Opt.PageMarginLeft, y)
	y = y + (common.Opt.Spacing * 2)

	DrawLine(common.Point{X: common.Opt.PageMarginLeft, Y: y}, common.Point{X: common.Opt.PageMarginRight, Y: y}, common.Opt.LineWidth, common.Opt.DarkBlack)
}

func drawLadderLineGroup() {

	for y := common.Opt.PageMarginTop; y <= common.Opt.PageMarginBottom; y += (common.Opt.Spacing * 12) {
		if y+(common.Opt.Spacing*8) > common.Opt.PageMarginBottom {
			return
		}
		drawLadderLines(y)
	}

}

// func drawLadderBars() {

// 	for yMaster := common.Opt.PageMarginTop; yMaster <= common.Opt.PageMarginBottom; yMaster += (common.Opt.Spacing * 8) {

// 		var racount int = 0

// 		for y := yMaster; y <= yMaster+(common.Opt.Spacing*9); y += common.Opt.Spacing {

// 			switch racount {
// 			case 0:
// 				DrawLine(common.Point{common.Opt.PageMarginLeft, y}, common.Point{common.Opt.PageMarginRight, y}, common.Opt.LineWidth, common.Opt.DarkBlack)
// 				drawLadder(common.Opt.PageMarginLeft, y)
// 				/*
// 					common.Opt.GC.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
// 					common.Opt.GC.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})

// 					common.Opt.GC.BeginPath()
// 					common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, y)
// 					common.Opt.GC.LineTo(common.Opt.PageMarginLeft+4.5, y)
// 					common.Opt.GC.LineTo(common.Opt.PageMarginLeft+4.5, y+4.5)
// 					common.Opt.GC.LineTo(common.Opt.PageMarginLeft, y+4.5)
// 					common.Opt.GC.Close()
// 					common.Opt.GC.FillStroke()
// 				*/
// 				racount++
// 			case 1:
// 				racount++
// 			case 2:
// 				DrawLine(common.Point{common.Opt.PageMarginLeft, y}, common.Point{common.Opt.PageMarginRight, y}, common.Opt.LineWidth, common.Opt.DarkBlack)
// 				drawLadder(common.Opt.PageMarginLeft, y)
// 				racount++
// 			case 3:
// 				racount++
// 			case 4:
// 				drawLadder(common.Opt.PageMarginLeft, y)
// 				racount++
// 			case 5:
// 				racount++
// 			case 6:
// 				DrawLine(common.Point{common.Opt.PageMarginLeft, y}, common.Point{common.Opt.PageMarginRight, y}, common.Opt.LineWidth, common.Opt.DarkBlack)
// 				drawLadder(common.Opt.PageMarginLeft, y)
// 				racount++
// 			case 7:
// 				racount++
// 			case 8:
// 				DrawLine(common.Point{common.Opt.PageMarginLeft, y}, common.Point{common.Opt.PageMarginRight, y}, common.Opt.LineWidth, common.Opt.DarkBlack)
// 				//drawLadder(common.Opt.PageMarginLeft, y)
// 				//y = y + (common.Opt.Spacing * 3)
// 				racount = 0
// 				continue

// 			}

// 		}
// 		yMaster = yMaster + (4 * common.Opt.Spacing)
// 	}
// }
