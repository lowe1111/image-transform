package imageTransform

import (
	"image"
	"image/color"
	"math"
)

const max = 65535

func higherContrast(num float64, contrast int8) float64 {
	contrastVal := 1 / float64(contrast)
	if num < max/2 {
		return num * contrastVal
	}
	if num*(1+contrastVal) > max {
		return max
	}
	return num * (1 + contrastVal)
}

func addContrast(r uint32, g uint32, b uint32, contrast int8) color.Gray16 {
	avg := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
	withContrast := higherContrast(avg, contrast)
	return color.Gray16{uint16(math.Round(withContrast))}
}

func GreyContrast(img image.Image, contrast int8) *image.Gray16 {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray16(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()
			grayColor := addContrast(r, g, b, contrast)
			gray.Set(x, y, grayColor)
		}
	}
	return gray
}
