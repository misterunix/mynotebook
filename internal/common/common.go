package common

import (
	"image/color"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/llgcode/draw2d/draw2dpdf"
)

type Point struct {
	X float64
	Y float64
}

// going to use bad programming practive and use global variables.
// All tied to this struct.
type options struct {
	Ladder           bool                      // for blackletter 2/4/2/4
	Spacing          float64                   // spacing between dots or lines in mm
	Centermark       bool                      // draw center dot or line
	CenterSpaceing   float64                   // center mark spacing
	Dot              bool                      // draw dots or lines
	Border           float64                   // border in mm
	Angle            float64                   // angle in degrees offset of center mark
	Cursiveunits     float64                   // units for cursive grid
	PaperOrientation string                    // paper orientation
	PageWidth        float64                   // page width in mm
	PageHeight       float64                   // page height in mm
	Margins          float64                   // page margins in mm
	LRmargin         float64                   // left and right margin
	PageMarginLeft   float64                   // page margin left in mm
	PageMarginRight  float64                   // page margin right in mm
	PageMarginTop    float64                   // page margin top in mm
	PageMarginBottom float64                   // page margin bottom in mm
	PaperSize        string                    // paper size
	LineWidth        float64                   // line width in mm
	Dest             *gofpdf.Fpdf              // PDF surface
	GC               *draw2dpdf.GraphicContext // graphic context
	Filename         string                    // filename
	File             *os.File
	DarkBlack        color.RGBA
	LightGray        color.RGBA
}

var Opt options
