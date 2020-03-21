package read

import (
	"image"
	_ "image/png"
	"os"

	"github.com/sqweek/dialog"
)

func GetPath() string {
	filename, err := dialog.File().Filter("PNG image file", "png").Load()
	if err != nil {
		panic(err)
	}
	return filename
}

func ReadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}
