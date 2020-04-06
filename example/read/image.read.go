package read

import (
	"image"
	_ "image/png"
	"os"

	"github.com/sqweek/dialog"
)

func GetImageFromPath() (image.Image, error) {
	filename := getPath()
	return readImage(filename)
}

func getPath() string {
	filename, err := dialog.File().Filter("PNG image file", "png").Load()
	if err != nil {
		panic(err)
	}
	return filename
}

func readImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}
