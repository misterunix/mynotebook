package pdf

import (
	"fmt"
	"image/color"
	"mynotebook/internal/common"
)

func DrawLine(a common.Point, b common.Point, width float64, linecolor color.RGBA) {
	CreateGC()
	fmt.Println(a, b)
	common.Opt.GC.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	common.Opt.GC.SetLineWidth(width)
	common.Opt.GC.MoveTo(a.X, a.Y)
	common.Opt.GC.LineTo(b.X, b.Y)
	common.Opt.GC.Close()
	common.Opt.GC.FillStroke()
}
