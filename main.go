package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"image/color"

	"github.com/jung-kurt/gofpdf"
	"github.com/llgcode/draw2d/draw2dpdf"
)

type point struct {
	x float64
	y float64
}

// going to use bad programming practive and use global variables.
// All tied to this struct.
type options struct {
	ladder         bool    // for blackletter 2/4/2/4
	spacing        float64 // spacing between dots or lines in mm
	centermark     bool    // draw center dot or line
	centerSpaceing float64 // center mark spacing
	dot            bool    // draw dots or lines
	border         float64 // border in mm
	angle          float64 // angle in degrees offset of center mark

	cursiveunits float64 // units for cursive grid

	paperOrientation string // paper orientation

	pageWidth  float64 // page width in mm
	pageHeight float64 // page height in mm

	margins          float64 // page margins in mm
	lrmargin         float64 // left and right margin
	pageMarginLeft   float64 // page margin left in mm
	pageMarginRight  float64 // page margin right in mm
	pageMarginTop    float64 // page margin top in mm
	pageMarginBottom float64 // page margin bottom in mm
	paperSize        string  // paper size

	lineWidth float64 // line width in mm
	//NewDest *gofpdf.Fpdf
	dest *gofpdf.Fpdf              // PDF surface
	gc   *draw2dpdf.GraphicContext // graphic context

	filename string // filename

	file *os.File

	darkBlack color.RGBA
	lightGray color.RGBA
}

var Opt options

func main() {

	style := 0

	flag.IntVar(&style, "style", 0, "page style\n  0 - lines\n  1 - dots\n  2 - cursive grid")

	flag.Float64Var(&Opt.spacing, "sp", 7.0, "spacing between dots or lines in mm")
	flag.BoolVar(&Opt.centermark, "c", false, "draw center dot or line")
	flag.StringVar(&Opt.paperOrientation, "o", "L", "paper orientation. L for landscape, P for portrait")
	flag.StringVar(&Opt.paperSize, "ps", "Letter", "paper size. Letter, A4, etc")
	flag.Float64Var(&Opt.cursiveunits, "u", 5.0, "units for cursive grid, overrides spacing")
	flag.Float64Var(&Opt.angle, "a", 0.0, "angle in degrees offset of center mark")
	flag.BoolVar(&Opt.ladder, "l", false, "for blackletter 2/4/2/4")
	flag.Parse()

	Opt.lrmargin = 12.7

	if Opt.centermark {
		Opt.centerSpaceing = Opt.spacing / 2.0
	}

	Opt.lineWidth = 0.2 // line width in mm
	if Opt.paperSize != "" {
		Opt.paperSize = strings.ToUpper(Opt.paperSize)
	} else {
		fmt.Println("Invalid paper size")
		os.Exit(1)
	}

	switch Opt.paperSize {
	case "LETTER":
		switch Opt.paperOrientation {
		case "L":
			Opt.pageWidth = 279.4
			Opt.pageHeight = 215.9
			Opt.margins = Opt.lrmargin / 2
		case "P":
			Opt.pageWidth = 215.9
			Opt.pageHeight = 279.4
			Opt.margins = Opt.lrmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}
	case "A4":
		switch Opt.paperOrientation {
		case "L":
			Opt.pageWidth = 297
			Opt.pageHeight = 210
			Opt.margins = Opt.lrmargin
		case "P":
			Opt.pageWidth = 210
			Opt.pageHeight = 297
			Opt.margins = Opt.lrmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}
	case "B5":
		switch Opt.paperOrientation {
		case "L":
			Opt.pageWidth = 250
			Opt.pageHeight = 176
			Opt.margins = Opt.lrmargin
		case "P":
			Opt.pageWidth = 176
			Opt.pageHeight = 250
			Opt.margins = Opt.lrmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid paper size")
		os.Exit(1)
	}

	Opt.darkBlack = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	Opt.lightGray = color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff}

	os.Mkdir("pdf", 0755)
	// if Opt.dot {
	// 	if Opt.centermark {
	// 		Opt.filename = fmt.Sprintf("pdf/dots-%s-%s-%d-center.pdf", Opt.paperSize, Opt.paperOrientation, int(Opt.spacing))
	// 	} else {
	// 		Opt.filename = fmt.Sprintf("pdf/dots-%s-%s-%d.pdf", Opt.paperSize, Opt.paperOrientation, int(Opt.spacing))
	// 	}
	// } else {
	// 	if Opt.centermark {
	// 		Opt.filename = fmt.Sprintf("pdf/lines-%s-%s-%d-center.pdf", Opt.paperSize, Opt.paperOrientation, int(Opt.spacing))
	// 	} else {
	// 		Opt.filename = fmt.Sprintf("pdf/lines-%s-%s-%d.pdf", Opt.paperSize, Opt.paperOrientation, int(Opt.spacing))
	// 	}
	// }

	Opt.pageMarginLeft = Opt.margins
	Opt.pageMarginRight = Opt.pageWidth - Opt.margins
	Opt.pageMarginTop = Opt.margins
	Opt.pageMarginBottom = Opt.pageHeight - Opt.margins

	fmt.Printf("%v\n", Opt)

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
	//Opt.NewDest := draw2dpdf.NewPdf(Opt.paperOrientation, "mm", Opt.paperSize)
	Opt.dest = draw2dpdf.NewPdf(Opt.paperOrientation, "mm", Opt.paperSize)
}

// create a new Graphic context
func createGC() {
	Opt.gc = draw2dpdf.NewGraphicContext(Opt.dest)
}

func drawLines() {
	if Opt.centermark {
		Opt.filename = fmt.Sprintf("pdf/lines-%s-%s-%02.3f-center.pdf", Opt.paperSize, Opt.paperOrientation, Opt.spacing)
	} else {
		Opt.filename = fmt.Sprintf("pdf/lines-%s-%s-%02.3f.pdf", Opt.paperSize, Opt.paperOrientation, Opt.spacing)
	}

	createPDFBase()

	if Opt.ladder {
		drawLadderLineGroup()
	}

	if Opt.centermark {
		for y := Opt.pageMarginTop + Opt.centerSpaceing; y <= Opt.pageMarginBottom; y += Opt.spacing {
			drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.lightGray)
		}
	}

	/*


		createGC()

		//gc.SetFillColor(color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff})

		// set stroke color
		Opt.gc.SetStrokeColor(color.RGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xff})

		// set line width
		Opt.gc.SetLineWidth(Opt.lineWidth)

		// base line
		for y := Opt.pageMarginTop; y < Opt.pageMarginBottom; y += Opt.spacing {
			Opt.gc.MoveTo(Opt.pageMarginLeft, y)
			Opt.gc.LineTo(Opt.pageMarginRight, y)
			//fmt.Println(Opt.pageMarginLeft, y, Opt.pageMarginRight, y)
		}

		// close Graphic context
		Opt.gc.Close()

		// fill and stroke
		Opt.gc.FillStroke()

		// center line if set
		if Opt.centermark {
			// create a new Graphic context
			createGC()

			Opt.gc.SetStrokeColor(color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xff})
			Opt.gc.SetLineWidth(Opt.lineWidth) // set line width
			count := 0
			for y := Opt.pageMarginTop; y < Opt.pageMarginBottom; y += Opt.spacing {
				if count == 0 {
					count++
					continue
				}
				y1 := y - Opt.centerSpaceing
				Opt.gc.MoveTo(Opt.pageMarginLeft, y1)
				Opt.gc.LineTo(Opt.pageMarginRight, y1)
			}

			// close Graphic context
			Opt.gc.Close()

			// fill and stroke
			Opt.gc.FillStroke()
		}
	*/
	// save to file
	draw2dpdf.SaveToPdfFile(Opt.filename, Opt.dest)

}

func drawDot(a point, radius float64, width float64, linecolor color.RGBA) {
	createGC()
	Opt.gc.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	Opt.gc.SetLineWidth(width)
	Opt.gc.MoveTo(a.x, a.y)
	Opt.gc.ArcTo(a.x, a.y, radius, radius, 0, 2*math.Pi)
	Opt.gc.Close()
	Opt.gc.FillStroke()
}

func drawDots() {
	if Opt.centermark {
		Opt.filename = fmt.Sprintf("pdf/dots-%s-%s-%f-center.pdf", Opt.paperSize, Opt.paperOrientation, Opt.spacing)
	} else {
		Opt.filename = fmt.Sprintf("pdf/dots-%s-%s-%f.pdf", Opt.paperSize, Opt.paperOrientation, Opt.spacing)
	}

	createPDFBase()

	for y := Opt.pageMarginTop; y <= Opt.pageMarginBottom; y += Opt.spacing {
		for x := Opt.pageMarginLeft; x <= Opt.pageMarginRight; x += Opt.spacing {
			drawDot(point{x, y}, 0.15, Opt.lineWidth, Opt.darkBlack)
		}
		//drawLine(point{Opt.pageMarginLeft, y}, point{Opt.pageMarginRight, y}, Opt.lineWidth, Opt.darkBlack)
	}

	if Opt.centermark {
		for y := Opt.pageMarginTop + Opt.centerSpaceing; y <= Opt.pageMarginBottom; y += Opt.spacing {
			for x := Opt.pageMarginLeft + Opt.centerSpaceing; x <= Opt.pageMarginRight; x += Opt.spacing {
				//angle := 0.0 //-12.0
				s := math.Sin(Opt.angle*math.Pi/180) * Opt.spacing
				drawDot(point{x + s, y}, 0.15, Opt.lineWidth, Opt.lightGray)
			}
		}
	}

	/*

		createGC()
		// set stroke color
		Opt.gc.SetStrokeColor(color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff})
		Opt.gc.SetLineWidth(Opt.lineWidth)

		Opt.border = 25 / 2.0
		xb := 25.0 - Opt.border
		yb := 25.0 - Opt.border

		for x := xb; x < Opt.pageWidth-Opt.border; x += Opt.spacing {
			for y := yb; y < Opt.pageHeight-Opt.border; y += Opt.spacing {

				Opt.gc.MoveTo(x, y)
				Opt.gc.ArcTo(x, y, 0.15, 0.15, 0, 2*math.Pi)
				Opt.gc.Close()
			}
		}

		Opt.gc.Close()
		Opt.gc.FillStroke()

		// center line if set
		if Opt.centermark {

			Opt.gc.SetStrokeColor(color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xff})
			Opt.gc.SetLineWidth(Opt.lineWidth)
			for x := xb + Opt.spacing; x < Opt.pageWidth-Opt.border; x += Opt.spacing {
				for y := yb; y < Opt.pageHeight-Opt.border; y += Opt.spacing {

					Opt.gc.MoveTo(x-(Opt.spacing/2), y-(Opt.spacing/2))
					Opt.gc.ArcTo(x-(Opt.spacing/2), y-(Opt.spacing/2), 0.15, 0.15, 0, 2*math.Pi)
					Opt.gc.Close()
				}
			}

			Opt.gc.Close()
			Opt.gc.FillStroke()
		}
	*/

	draw2dpdf.SaveToPdfFile(Opt.filename, Opt.dest)

}

func cursivegrid() {

	Opt.spacing = Opt.cursiveunits
	Opt.filename = fmt.Sprintf("pdf/cursive-%s-%s-%f-center.pdf", Opt.paperSize, Opt.paperOrientation, Opt.spacing)

	createPDFBase()
	createGC()

	Opt.gc.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	Opt.gc.SetLineWidth(Opt.lineWidth)

	Opt.gc.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	Opt.gc.SetLineWidth(Opt.lineWidth)

	down := 0.0
	pos := 0.0

	for {

		// ascender line
		Opt.gc.SetLineWidth(0.5)
		Opt.gc.SetStrokeColor(Opt.darkBlack)
		down = Opt.pageMarginTop + pos
		Opt.gc.MoveTo(Opt.pageMarginLeft, down)
		Opt.gc.LineTo(Opt.pageMarginRight, down)
		Opt.gc.Close()
		Opt.gc.FillStroke()

		// t-d line
		Opt.gc.SetLineWidth(0.2)
		Opt.gc.SetStrokeColor(Opt.lightGray)
		down = Opt.pageMarginTop + pos + Opt.spacing

		Opt.gc.MoveTo(Opt.pageMarginLeft, down)
		Opt.gc.LineTo(Opt.pageMarginRight, down)

		Opt.gc.Close()
		Opt.gc.FillStroke()

		// x-height line
		Opt.gc.SetLineWidth(0.2)
		down = Opt.pageMarginTop + pos + (Opt.spacing * 2)

		Opt.gc.MoveTo(Opt.pageMarginLeft, down)
		Opt.gc.LineTo(Opt.pageMarginRight, down)

		Opt.gc.Close()
		Opt.gc.FillStroke()

		// base line
		Opt.gc.SetLineWidth(0.5)
		Opt.gc.SetStrokeColor(Opt.darkBlack)
		down = Opt.pageMarginTop + pos + (Opt.spacing * 3)

		Opt.gc.MoveTo(Opt.pageMarginLeft, down)
		Opt.gc.LineTo(Opt.pageMarginRight, down)

		Opt.gc.Close()
		Opt.gc.FillStroke()

		// descender line
		Opt.gc.SetLineWidth(0.5)
		Opt.gc.SetStrokeColor(Opt.darkBlack)
		down = Opt.pageMarginTop + pos + (Opt.spacing * 5)

		Opt.gc.MoveTo(Opt.pageMarginLeft, down)
		Opt.gc.LineTo(Opt.pageMarginRight, down)

		Opt.gc.Close()
		Opt.gc.FillStroke()

		pos += (Opt.spacing * 6)

		//fmt.Println(pos, (Opt.pageHeight - Opt.pageMarginBottom))

		if pos > (Opt.pageHeight - Opt.margins) {
			break
		}

	}

	draw2dpdf.SaveToPdfFile(Opt.filename, Opt.dest)
}
