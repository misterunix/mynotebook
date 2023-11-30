package main

import (
	"fmt"
	"os"
)

func main() {

	b := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n"
	b += "<!-- Created with Inkscape (http://www.inkscape.org/) -->\n"
	b += "<svg width=\"279.39999mm\" height=\"215.89999mm\"\n"
	b += "viewBox=\"0 0 279.39999 215.89999\"\n"
	b += "version=\"1.1\" id=\"svg1\" inkscape:version=\"1.3.1 (6036e22fae, 2023-11-19, custom)\"\n"
	// b+= "sodipodi:docname=\"dot.svg\"\n"
	//
	//	xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
	//	xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
	b += "xmlns=\"http://www.w3.org/2000/svg\"\n"
	b += "xmlns:svg=\"http://www.w3.org/2000/svg\">\n"

	c := "<defs\nid=\"defs1\" />\n<g\ninkscape:label=\"Layer 1\"\ninkscape:groupmode=\"layer\"\nid=\"layer1\">\n"

	// file := os.Open("dot1.svg", os.O_CREATE|os.O_WRONLY, 0644)
	file, err := os.OpenFile("dot1.svg", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(b)
	file.WriteString(c)

	border := 25 / 2.0

	xb := 25.0 - border
	yb := 25.0 - border

	for x := xb; x < 279-border; x += 5 {
		for y := yb; y < 215-border; y += 5 {

			d := "<ellipse style=\"fill:none;stroke:#AAAAAA;stroke-width:0.2\" id=\"path1\" "
			xx := fmt.Sprintf("%f", x)
			yy := fmt.Sprintf("%f", y)

			d += "cx=\"" + xx + "\" "
			d += "cy=\"" + yy + "\" "
			//cy="5.0"
			d += "rx=\"0.15\" ry=\"0.15\" />\n"

			d += "<ellipse style=\"fill:none;stroke:#CCCCCC;stroke-width:0.2\" id=\"path1\" "
			xx = fmt.Sprintf("%f", x+2.5)
			yy = fmt.Sprintf("%f", y+2.5)

			d += "cx=\"" + xx + "\" "
			d += "cy=\"" + yy + "\" "
			//cy="5.0"
			d += "rx=\"0.11\" ry=\"0.11\" />\n"

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
