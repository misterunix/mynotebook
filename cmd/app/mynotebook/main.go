package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"image/color"

	"mynotebook/internal/common"

	"github.com/llgcode/draw2d/draw2dpdf"
)

func main() {

	style := 0

	flag.IntVar(&style, "style", 0, "page style\n  0 - lines\n  1 - dots\n  2 - cursive grid")

	flag.Float64Var(&common.Opt.Spacing, "sp", 7.0, "spacing between dots or lines in mm")
	flag.BoolVar(&common.Opt.Centermark, "c", false, "draw center dot or line")
	flag.StringVar(&common.Opt.PaperOrientation, "o", "L", "paper orientation. L for landscape, P for portrait")
	flag.StringVar(&common.Opt.PaperSize, "ps", "Letter", "paper size. Letter, A4, etc")
	flag.Float64Var(&common.Opt.Cursiveunits, "u", 5.0, "units for cursive grid, overrides spacing")
	flag.Float64Var(&common.Opt.Angle, "a", 0.0, "angle in degrees offset of center mark")
	flag.BoolVar(&common.Opt.Ladder, "l", false, "for blackletter 2/4/2/4")
	flag.Parse()

	common.Opt.LRmargin = 12.7

	if common.Opt.Centermark {
		common.Opt.CenterSpaceing = common.Opt.Spacing / 2.0
	}

	common.Opt.LineWidth = 0.2 // line width in mm
	if common.Opt.PaperSize != "" {
		common.Opt.PaperSize = strings.ToUpper(common.Opt.PaperSize)
	} else {
		fmt.Println("Invalid paper size")
		os.Exit(1)
	}

	switch common.Opt.PaperSize {
	case "LETTER":
		switch common.Opt.PaperOrientation {
		case "L":
			common.Opt.PageWidth = 279.4
			common.Opt.PageHeight = 215.9
			common.Opt.Margins = common.Opt.LRmargin / 2
		case "P":
			common.Opt.PageWidth = 215.9
			common.Opt.PageHeight = 279.4
			common.Opt.Margins = common.Opt.LRmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}
	case "A4":
		switch common.Opt.PaperOrientation {
		case "L":
			common.Opt.PageWidth = 297
			common.Opt.PageHeight = 210
			common.Opt.Margins = common.Opt.LRmargin
		case "P":
			common.Opt.PageWidth = 210
			common.Opt.PageHeight = 297
			common.Opt.Margins = common.Opt.LRmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}
	case "B5":
		switch common.Opt.PaperOrientation {
		case "L":
			common.Opt.PageWidth = 250
			common.Opt.PageHeight = 176
			common.Opt.Margins = common.Opt.LRmargin
		case "P":
			common.Opt.PageWidth = 176
			common.Opt.PageHeight = 250
			common.Opt.Margins = common.Opt.LRmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid paper size")
		os.Exit(1)
	}

	common.Opt.DarkBlack = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	common.Opt.LightGray = color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff}

	os.Mkdir("pdf", 0755)
	// if common.Opt.dot {
	// 	if common.Opt.Centermark {
	// 		common.Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%d-center.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, int(common.Opt.Spacing))
	// 	} else {
	// 		common.Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%d.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, int(common.Opt.Spacing))
	// 	}
	// } else {
	// 	if common.Opt.Centermark {
	// 		common.Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%d-center.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, int(common.Opt.Spacing))
	// 	} else {
	// 		common.Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%d.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, int(common.Opt.Spacing))
	// 	}
	// }

	common.Opt.PageMarginLeft = common.Opt.Margins
	common.Opt.PageMarginRight = common.Opt.PageWidth - common.Opt.Margins
	common.Opt.PageMarginTop = common.Opt.Margins
	common.Opt.PageMarginBottom = common.Opt.PageHeight - common.Opt.Margins

	fmt.Printf("%v\n", common.Opt)

	switch style {
	case 0:
		drawLines()
	case 1:
		drawDots()
	case 2:
		cursivegrid()
	default:
		fmt.Println("Invalid style")
		os.Exit(1)
	}

}

// createPDFBase creates a new PDF surface with a given orientation and a given unit
func createPDFBase() {
	//common.Opt.NewDest := draw2dpdf.NewPdf(common.Opt.PaperOrientation, "mm", common.Opt.PaperSize)
	common.Opt.Dest = draw2dpdf.NewPdf(common.Opt.PaperOrientation, "mm", common.Opt.PaperSize)
}

// create a new Graphic context
func createGC() {
	common.Opt.GC = draw2dpdf.NewGraphicContext(common.Opt.Dest)
}

func drawLines() {
	if common.Opt.Centermark {
		common.Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%02.3f-center.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, common.Opt.Spacing)
	} else {
		common.Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%02.3f.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, common.Opt.Spacing)
	}

	createPDFBase()

	if common.Opt.Ladder {
		drawLadderLineGroup()
	}

	if common.Opt.Centermark {
		for y := common.Opt.PageMarginTop + common.Opt.CenterSpaceing; y <= common.Opt.PageMarginBottom; y += common.Opt.Spacing {
			drawLine(point{common.Opt.PageMarginLeft, y}, point{common.Opt.PageMarginRight, y}, common.Opt.LineWidth, common.Opt.LightGray)
		}
	}

	/*


		createGC()

		//gc.SetFillColor(color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff})

		// set stroke color
		common.Opt.GC.SetStrokeColor(color.RGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xff})

		// set line width
		common.Opt.GC.SetLineWidth(common.Opt.LineWidth)

		// base line
		for y := common.Opt.PageMarginTop; y < common.Opt.PageMarginBottom; y += common.Opt.Spacing {
			common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, y)
			common.Opt.GC.LineTo(common.Opt.PageMarginRight, y)
			//fmt.Println(common.Opt.PageMarginLeft, y, common.Opt.PageMarginRight, y)
		}

		// close Graphic context
		common.Opt.GC.Close()

		// fill and stroke
		common.Opt.GC.FillStroke()

		// center line if set
		if common.Opt.Centermark {
			// create a new Graphic context
			createGC()

			common.Opt.GC.SetStrokeColor(color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xff})
			common.Opt.GC.SetLineWidth(common.Opt.LineWidth) // set line width
			count := 0
			for y := common.Opt.PageMarginTop; y < common.Opt.PageMarginBottom; y += common.Opt.Spacing {
				if count == 0 {
					count++
					continue
				}
				y1 := y - common.Opt.CenterSpaceing
				common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, y1)
				common.Opt.GC.LineTo(common.Opt.PageMarginRight, y1)
			}

			// close Graphic context
			common.Opt.GC.Close()

			// fill and stroke
			common.Opt.GC.FillStroke()
		}
	*/
	// save to file
	draw2dpdf.SaveToPdfFile(common.Opt.Filename, common.Opt.Dest)

}

func drawDot(a point, radius float64, width float64, linecolor color.RGBA) {
	createGC()
	common.Opt.GC.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	common.Opt.GC.SetLineWidth(width)
	common.Opt.GC.MoveTo(a.x, a.y)
	common.Opt.GC.ArcTo(a.x, a.y, radius, radius, 0, 2*math.Pi)
	common.Opt.GC.Close()
	common.Opt.GC.FillStroke()
}

func drawDots() {
	if common.Opt.Centermark {
		common.Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%f-center.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, common.Opt.Spacing)
	} else {
		common.Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%f.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, common.Opt.Spacing)
	}

	createPDFBase()

	for y := common.Opt.PageMarginTop; y <= common.Opt.PageMarginBottom; y += common.Opt.Spacing {
		for x := common.Opt.PageMarginLeft; x <= common.Opt.PageMarginRight; x += common.Opt.Spacing {
			drawDot(point{x, y}, 0.15, common.Opt.LineWidth, common.Opt.DarkBlack)
		}
		//drawLine(point{common.Opt.PageMarginLeft, y}, point{common.Opt.PageMarginRight, y}, common.Opt.LineWidth, common.Opt.DarkBlack)
	}

	if common.Opt.Centermark {
		for y := common.Opt.PageMarginTop + common.Opt.CenterSpaceing; y <= common.Opt.PageMarginBottom; y += common.Opt.Spacing {
			for x := common.Opt.PageMarginLeft + common.Opt.CenterSpaceing; x <= common.Opt.PageMarginRight; x += common.Opt.Spacing {
				//angle := 0.0 //-12.0
				s := math.Sin(common.Opt.Angle*math.Pi/180) * common.Opt.Spacing
				drawDot(point{x + s, y}, 0.15, common.Opt.LineWidth, common.Opt.LightGray)
			}
		}
	}

	/*

		createGC()
		// set stroke color
		common.Opt.GC.SetStrokeColor(color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff})
		common.Opt.GC.SetLineWidth(common.Opt.LineWidth)

		common.Opt.border = 25 / 2.0
		xb := 25.0 - common.Opt.border
		yb := 25.0 - common.Opt.border

		for x := xb; x < common.Opt.PageWidth-common.Opt.border; x += common.Opt.Spacing {
			for y := yb; y < common.Opt.PageHeight-common.Opt.border; y += common.Opt.Spacing {

				common.Opt.GC.MoveTo(x, y)
				common.Opt.GC.ArcTo(x, y, 0.15, 0.15, 0, 2*math.Pi)
				common.Opt.GC.Close()
			}
		}

		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		// center line if set
		if common.Opt.Centermark {

			common.Opt.GC.SetStrokeColor(color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xff})
			common.Opt.GC.SetLineWidth(common.Opt.LineWidth)
			for x := xb + common.Opt.Spacing; x < common.Opt.PageWidth-common.Opt.border; x += common.Opt.Spacing {
				for y := yb; y < common.Opt.PageHeight-common.Opt.border; y += common.Opt.Spacing {

					common.Opt.GC.MoveTo(x-(common.Opt.Spacing/2), y-(common.Opt.Spacing/2))
					common.Opt.GC.ArcTo(x-(common.Opt.Spacing/2), y-(common.Opt.Spacing/2), 0.15, 0.15, 0, 2*math.Pi)
					common.Opt.GC.Close()
				}
			}

			common.Opt.GC.Close()
			common.Opt.GC.FillStroke()
		}
	*/

	draw2dpdf.SaveToPdfFile(common.Opt.Filename, common.Opt.Dest)

}

func cursivegrid() {

	common.Opt.Spacing = common.Opt.Cursiveunits
	common.Opt.Filename = fmt.Sprintf("pdf/cursive-%s-%s-%f-center.pdf", common.Opt.PaperSize, common.Opt.PaperOrientation, common.Opt.Spacing)

	createPDFBase()
	createGC()

	common.Opt.GC.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	common.Opt.GC.SetLineWidth(common.Opt.LineWidth)

	common.Opt.GC.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	common.Opt.GC.SetLineWidth(common.Opt.LineWidth)

	down := 0.0
	pos := 0.0

	for {

		// ascender line
		common.Opt.GC.SetLineWidth(0.5)
		common.Opt.GC.SetStrokeColor(common.Opt.DarkBlack)
		down = common.Opt.PageMarginTop + pos
		common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, down)
		common.Opt.GC.LineTo(common.Opt.PageMarginRight, down)
		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		// t-d line
		common.Opt.GC.SetLineWidth(0.2)
		common.Opt.GC.SetStrokeColor(common.Opt.LightGray)
		down = common.Opt.PageMarginTop + pos + common.Opt.Spacing

		common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, down)
		common.Opt.GC.LineTo(common.Opt.PageMarginRight, down)

		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		// x-height line
		common.Opt.GC.SetLineWidth(0.2)
		down = common.Opt.PageMarginTop + pos + (common.Opt.Spacing * 2)

		common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, down)
		common.Opt.GC.LineTo(common.Opt.PageMarginRight, down)

		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		// base line
		common.Opt.GC.SetLineWidth(0.5)
		common.Opt.GC.SetStrokeColor(common.Opt.DarkBlack)
		down = common.Opt.PageMarginTop + pos + (common.Opt.Spacing * 3)

		common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, down)
		common.Opt.GC.LineTo(common.Opt.PageMarginRight, down)

		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		// descender line
		common.Opt.GC.SetLineWidth(0.5)
		common.Opt.GC.SetStrokeColor(common.Opt.DarkBlack)
		down = common.Opt.PageMarginTop + pos + (common.Opt.Spacing * 5)

		common.Opt.GC.MoveTo(common.Opt.PageMarginLeft, down)
		common.Opt.GC.LineTo(common.Opt.PageMarginRight, down)

		common.Opt.GC.Close()
		common.Opt.GC.FillStroke()

		pos += (common.Opt.Spacing * 6)

		//fmt.Println(pos, (common.Opt.PageHeight - common.Opt.PageMarginBottom))

		if pos > (common.Opt.PageHeight - common.Opt.Margins) {
			break
		}

	}

	draw2dpdf.SaveToPdfFile(common.Opt.Filename, common.Opt.Dest)
}
