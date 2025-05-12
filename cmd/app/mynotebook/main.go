package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"image/color"

	"mynotebook/internal/common"
	"mynotebook/internal/pdf"
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
	flag.BoolVar(&common.Opt.Dark, "dark", false, "dark lines")
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
	if common.Opt.Dark {
		common.Opt.LightGray = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	} else {
		common.Opt.LightGray = color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff}
	}

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
		pdf.DrawLines()
	case 1:
		pdf.DrawDots()
	case 2:
		pdf.CursiveGrid()
	default:
		fmt.Println("Invalid style")
		os.Exit(1)
	}

}
