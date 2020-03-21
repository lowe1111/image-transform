package write

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func WriteImage(img image.Image, locationPath string) {
	file := fmt.Sprintf("%s/result.png", locationPath)
	outfile, err := os.Create(file)
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
