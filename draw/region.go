package draw

import (
	"fmt"
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

// RegionWithPoint draws region and point for debuging purposes
func RegionWithPoint(points *[][]float64, x, y float64) {
	scale := 50

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(10*scale, 10*scale, 100*scale, 100*scale))

	// Draw polygon
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	gc.SetLineWidth(1)
	ppoints := *points
	gc.MoveTo(ppoints[0][1]*float64(scale), ppoints[0][0]*float64(scale))
	for _, v := range ppoints {
		gc.LineTo(v[1]*float64(scale), v[0]*float64(scale))
	}
	gc.LineTo(ppoints[0][1]*float64(scale), ppoints[0][0]*float64(scale))
	gc.FillStroke()
	gc.Close()

	// Draw point
	gc = draw2dimg.NewGraphicContext(dest)
	gc.MoveTo((x-1)*float64(scale), (y-1)*float64(scale))
	gc.LineTo((x+1)*float64(scale), (y+1)*float64(scale))
	gc.MoveTo((x-1)*float64(scale), (y+1)*float64(scale))
	gc.LineTo((x+1)*float64(scale), (y-1)*float64(scale))
	gc.FillStroke()
	gc.Close()

	// Save to file
	filePath := fmt.Sprintf("tmp/region_name_%f_%f.png", x, y)
	draw2dimg.SaveToPngFile(filePath, dest)
}
