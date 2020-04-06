package imageTransform

import (
	"image"
	"image/color"
	"math"
)

func normalize(col float64) uint16 {
	if col > 65535 {
		return 65535
	}
	return uint16(math.Floor(col))
}

func sepianizeColor(r, g, b, a uint32) color.RGBA64 {
	tr := 0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b)
	tg := 0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b)
	tb := 0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b)
	return color.RGBA64{normalize(tr), normalize(tg), normalize(tb), uint16(a)}
}

func SepiaImage(img image.Image) *image.RGBA64 {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newImage := image.NewRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := img.At(x, y)
			r, g, b, a := oldColor.RGBA()
			reversedColors := sepianizeColor(r, g, b, a)
			newImage.Set(x, y, reversedColors)
		}
	}
	return newImage
}
