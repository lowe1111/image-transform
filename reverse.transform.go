package imageTransform

import (
	"image"
	"image/color"
)

func reverseColors(r, g, b, a uint32) color.RGBA64 {
	return color.RGBA64{uint16(0xffff - r), uint16(0xffff - g), uint16(0xffff - b), uint16(a)}
}

func ReverseImage(img image.Image) *image.RGBA64 {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newImage := image.NewRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := img.At(x, y)
			r, g, b, a := oldColor.RGBA()
			reversedColors := reverseColors(r, g, b, a)
			newImage.Set(x, y, reversedColors)
		}
	}
	return newImage
}
