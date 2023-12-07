package main

import (
	"flag"
	"fmt"
	"os"

	"image/color"

	"github.com/llgcode/draw2d/draw2dpdf"
)

type options struct {
	spacing        float64 // spacing between dots or lines in mm
	centermark     bool    // draw center dot or line
	centerSpaceing float64 // center mark spacing
	dot            bool    // draw dots or lines
	border         float64 // border in mm

	paperOrientation string // paper orientation

	pageWidth  float64 // page width in mm
	pageHeight float64 // page height in mm

	margins          float64 // page margins in mm
	pageMarginLeft   float64 // page margin left in mm
	pageMarginRight  float64 // page margin right in mm
	pageMarginTop    float64 // page margin top in mm
	pageMarginBottom float64 // page margin bottom in mm

	lineWidth float64 // line width in mm

	file *os.File
}

var Opt options

func main() {

	flag.BoolVar(&Opt.dot, "dot", false, "draw dots or lines")
	flag.Float64Var(&Opt.spacing, "s", 7.0, "spacing between dots or lines in mm")
	flag.BoolVar(&Opt.centermark, "c", false, "draw center dot or line")
	flag.StringVar(&Opt.paperOrientation, "o", "L", "paper orientation. L for landscape, P for portrait")

	flag.Parse()

	if Opt.centermark {
		Opt.centerSpaceing = Opt.spacing / 2.0
	}

	Opt.lineWidth = 0.2 // line width in mm

	Opt.pageWidth = 279.4
	Opt.pageHeight = 215.9

	Opt.margins = 25.4 // page margins in mm
	Opt.pageMarginLeft = Opt.margins
	Opt.pageMarginRight = Opt.pageWidth - Opt.margins
	Opt.pageMarginTop = Opt.margins
	Opt.pageMarginBottom = Opt.pageHeight - Opt.margins

	if Opt.dot {
		drawDots()
	} else {
		drawLines()
	}

}

func drawLines() {

	//margin := 25.4
	//marginHalf := margin / 2.0

	// create a new PDF surface with a given orientation and a given unit
	dest := draw2dpdf.NewPdf(Opt.paperOrientation, "mm", "Letter")

	// create a new Graphic context
	gc := draw2dpdf.NewGraphicContext(dest)

	//gc.SetFillColor(color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff})

	// set stroke color
	gc.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	gc.SetLineWidth(Opt.lineWidth)

	// base line
	for y := Opt.pageMarginTop; y < Opt.pageMarginBottom; y += Opt.spacing {
		gc.MoveTo(Opt.pageMarginLeft, y)
		gc.LineTo(Opt.pageMarginRight, y)
		//fmt.Println(Opt.pageMarginLeft, y, Opt.pageMarginRight, y)
	}

	// close Graphic context
	gc.Close()

	// fill and stroke
	gc.FillStroke()

	// center line

	// create a new Graphic context
	gc = draw2dpdf.NewGraphicContext(dest)
	gc.SetStrokeColor(color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xff})
	gc.SetLineWidth(Opt.lineWidth) // set line width

	for y := Opt.pageMarginTop; y < Opt.pageMarginBottom; y += Opt.spacing {
		y1 := y - Opt.centerSpaceing
		gc.MoveTo(Opt.pageMarginLeft, y1)
		gc.LineTo(Opt.pageMarginRight, y1)
	}

	// close Graphic context
	gc.Close()

	// fill and stroke
	gc.FillStroke()

	// save to file
	draw2dpdf.SaveToPdfFile("lines.pdf", dest)

	// var err error
	// baseSVG := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n"
	// baseSVG += "<!-- Created with Inkscape (http://www.inkscape.org/) -->\n"
	// baseSVG += "<svg width=\"279.39999mm\" height=\"215.89999mm\"\n"
	// baseSVG += "viewBox=\"0 0 279.39999 215.89999\"\n"
	// baseSVG += "version=\"1.1\" id=\"svg1\" inkscape:version=\"1.3.1 (6036e22fae, 2023-11-19, custom)\"\n"
	// // b+= "sodipodi:docname=\"dot.svg\"\n"
	// //
	// //	xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
	// //	xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
	// baseSVG += "xmlns=\"http://www.w3.org/2000/svg\"\n"
	// baseSVG += "xmlns:svg=\"http://www.w3.org/2000/svg\">\n"

	// c := "<defs\n"
	// c += "id=\"defs1\" />\n"
	// c += "<g\n"
	// c += "inkscape:label=\"Layer 1\"\n"
	// c += "inkscape:groupmode=\"layer\"\n"
	// c += "id=\"layer1\">\n"

	// filename := fmt.Sprintf("svg/lines-%05.2fmm.svg", Opt.spacing)

	// // file := os.Open("dot1.svg", os.O_CREATE|os.O_WRONLY, 0644)
	// Opt.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer Opt.file.Close()

	// Opt.file.WriteString(baseSVG)
	// Opt.file.WriteString(c)

	// Opt.border = 25 / 2.0

	// xbl := 25.0 - Opt.border
	// xbr := 279.4 - Opt.border
	// yb := 25.0 - Opt.border

	// for y := yb; y < 215-Opt.border; y += Opt.spacing {
	// 	d := fmt.Sprintf(" <line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" style=\"fill:none;stroke:#AAAAAA;stroke-width:0.2\" id=\"path1\" />\n", xbl, y, xbr, y)

	// 	if Opt.centermark {
	// 		d += fmt.Sprintf(" <line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" style=\"fill:none;stroke:#CCCCCC;stroke-width:0.2\" id=\"path1\" />\n", xbl, y+Opt.center, xbr, y+Opt.center)
	// 	}

	// 	Opt.file.WriteString(d)
	// }

}

func drawDots() {
	var err error
	baseSVG := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n"
	baseSVG += "<!-- Created with Inkscape (http://www.inkscape.org/) -->\n"
	baseSVG += "<svg width=\"279.39999mm\" height=\"215.89999mm\"\n"
	baseSVG += "viewBox=\"0 0 279.39999 215.89999\"\n"
	baseSVG += "version=\"1.1\" id=\"svg1\" inkscape:version=\"1.3.1 (6036e22fae, 2023-11-19, custom)\"\n"
	// b+= "sodipodi:docname=\"dot.svg\"\n"
	//
	//	xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
	//	xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
	baseSVG += "xmlns=\"http://www.w3.org/2000/svg\"\n"
	baseSVG += "xmlns:svg=\"http://www.w3.org/2000/svg\">\n"

	c := "<defs\n"
	c += "id=\"defs1\" />\n"
	c += "<g\n"
	c += "inkscape:label=\"Layer 1\"\n"
	c += "inkscape:groupmode=\"layer\"\n"
	c += "id=\"layer1\">\n"

	filename := fmt.Sprintf("svg/dots-%05.2fmm.svg", Opt.spacing)

	// file := os.Open("dot1.svg", os.O_CREATE|os.O_WRONLY, 0644)
	Opt.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer Opt.file.Close()

	Opt.file.WriteString(baseSVG)
	Opt.file.WriteString(c)

	Opt.border = 25 / 2.0

	xb := 25.0 - Opt.border
	yb := 25.0 - Opt.border

	for x := xb; x < 279-Opt.border; x += Opt.spacing {
		for y := yb; y < 215-Opt.border; y += Opt.spacing {

			d := "<ellipse style=\"fill:none;stroke:#AAAAAA;stroke-width:0.2\" id=\"path1\" "
			xx := fmt.Sprintf("%f", x)
			yy := fmt.Sprintf("%f", y)

			d += "cx=\"" + xx + "\" "
			d += "cy=\"" + yy + "\" "
			//cy="5.0"
			d += "rx=\"0.15\" ry=\"0.15\" />\n"

			if Opt.centermark {
				d += "<ellipse style=\"fill:none;stroke:#CCCCCC;stroke-width:0.2\" id=\"path1\" "
				xx = fmt.Sprintf("%f", x+Opt.centerSpaceing)
				yy = fmt.Sprintf("%f", y+Opt.centerSpaceing)

				d += "cx=\"" + xx + "\" "
				d += "cy=\"" + yy + "\" "
				//cy="5.0"
				d += "rx=\"0.11\" ry=\"0.11\" />\n"
			}
			Opt.file.WriteString(d)
		}
	}

	// xb = 25 + 25/2.0
	// yb = 25 + 25/2.0

	// for x := xb; x < 279-xb; x += 5 {
	// 	for y := yb; y < 215-yb; y += 5 {

	// 		d := "<ellipse style=\"fill:none;stroke:#BBBBBB;stroke-width:0.3\" id=\"path1\" "
	// 		xx := fmt.Sprintf("%f", x)
	// 		yy := fmt.Sprintf("%f", y)
	// 		d += "cx=\"" + xx + "\" "
	// 		d += "cy=\"" + yy + "\" "
	// 		//cy="5.0"
	// 		d += "rx=\"0.24392819\" ry=\"0.24348988\" />\n"
	// 		file.WriteString(d)
	// 	}
	// }

	e := "</g>\n</svg>\n"
	Opt.file.WriteString(e)

}
