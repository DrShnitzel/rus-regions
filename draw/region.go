package draw

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawRegion(points *[][]float64) {
	scale := 50

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(10*scale, 10*scale, 100*scale, 100*scale))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x00, 0xff, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	gc.SetLineWidth(1)

	gc.MoveTo(100, 100) // should always be called first for a new path

	for _, v := range *points {
		gc.LineTo(v[1]*float64(scale), v[0]*float64(scale))
	}
	gc.FillStroke()
	gc.Close()
	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}
