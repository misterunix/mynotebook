package pdf

import (
	"image/color"
	"mynotebook/internal/common"
)

func DrawLine(a common.Point, b common.Point, width float64, linecolor color.RGBA) {
	CreateGC()
	common.Opt.GC.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	common.Opt.GC.SetLineWidth(width)
	common.Opt.GC.MoveTo(a.x, a.y)
	common.Opt.GC.LineTo(b.x, b.y)
	common.Opt.GC.Close()
	common.Opt.GC.FillStroke()
}
