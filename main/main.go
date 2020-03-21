package main

import (
	"fmt"
	"image-go/read"
	"image-go/transform"
	"image-go/write"
)

func main() {
	fmt.Print("Read your File...")
	filename := read.GetPath()
	image, _ := read.ReadImage(filename)
	greyImage := transform.ConverToGray(image)
	write.WriteImage(image, "results")
}
