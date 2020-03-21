package transform

import (
	"image"
	"image/color"
	"math"
)

func convertToGray(r uint32, g uint32, b uint32) color.Gray {
	avg := 0.2125*float64(r) + 0.7154*float64(g) + 0.0721*float64(b)
	return color.Gray{uint8(math.Ceil(avg))}
}

func ConverToGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()
			grayColor := convertToGray(r, g, b)
			gray.Set(x, y, grayColor)
		}
	}
	return gray;
}