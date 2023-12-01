package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var spacing float64
	var centermark bool
	var center float64
	var dot bool

	flag.BoolVar(&dot, "dot", false, "draw dots or lines")
	flag.Float64Var(&spacing, "s", 7.0, "spacing between dots or lines in mm")
	flag.BoolVar(&centermark, "centermark", false, "draw center dot or line")

	flag.Parse()

	if centermark {
		center = spacing / 2.0
	}

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

	filename := fmt.Sprintf("svg/dots-%05.2fmm.svg", spacing)

	// file := os.Open("dot1.svg", os.O_CREATE|os.O_WRONLY, 0644)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(baseSVG)
	file.WriteString(c)

	border := 25 / 2.0

	xb := 25.0 - border
	yb := 25.0 - border

	for x := xb; x < 279-border; x += spacing {
		for y := yb; y < 215-border; y += spacing {

			d := "<ellipse style=\"fill:none;stroke:#AAAAAA;stroke-width:0.2\" id=\"path1\" "
			xx := fmt.Sprintf("%f", x)
			yy := fmt.Sprintf("%f", y)

			d += "cx=\"" + xx + "\" "
			d += "cy=\"" + yy + "\" "
			//cy="5.0"
			d += "rx=\"0.15\" ry=\"0.15\" />\n"

			if centermark {
				d += "<ellipse style=\"fill:none;stroke:#CCCCCC;stroke-width:0.2\" id=\"path1\" "
				xx = fmt.Sprintf("%f", x+center)
				yy = fmt.Sprintf("%f", y+center)

				d += "cx=\"" + xx + "\" "
				d += "cy=\"" + yy + "\" "
				//cy="5.0"
				d += "rx=\"0.11\" ry=\"0.11\" />\n"
			}
			file.WriteString(d)
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
	file.WriteString(e)

}
