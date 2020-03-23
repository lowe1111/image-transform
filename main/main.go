package main

import (
	"fmt"
	"image-go/read"
	"image-go/transform"
	"image-go/write"
)

func main() {
	fmt.Print("Read your File...")
	image, _ := read.GetImageFromPath()
	greyImage := transform.ConverToGray(image)
	write.WriteImage(greyImage, "results")
}
