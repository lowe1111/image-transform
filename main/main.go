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
	reversedImg := transform.ReverseImage(image)
	greyImg := transform.ConverToGray(image)
	write.WriteImage(image, "results", "result")
	write.WriteImage(reversedImg, "results", "resultReversed")
	write.WriteImage(greyImg, "results", "resultGrey")
}
