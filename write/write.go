package write

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func WriteImage(img image.Image, locationPath string, fileName string) {
	file := fmt.Sprintf("%s/%s.png", locationPath, fileName)
	outfile, err := os.Create(file)
	defer outfile.Close()
	if err != nil {
		panic(err.Error())
	}
	png.Encode(outfile, img)
}
