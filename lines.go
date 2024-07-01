package main

import "image/color"

func drawLine(a point, b point, width float64, linecolor color.RGBA) {
	createGC()
	Opt.gc.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	Opt.gc.SetLineWidth(width)
	Opt.gc.MoveTo(a.x, a.y)
	Opt.gc.LineTo(b.x, b.y)
	Opt.gc.Close()
	Opt.gc.FillStroke()
}
