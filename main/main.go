package main

import (
	"fmt"
	"image-go/read"
	"image-go/transform"
	"image-go/write"
)

func main() {
	fmt.Print("Read your File...")
	image, err := read.GetImageFromPath()
	if err != nil {
		panic(err)
	}
	reversedImg := transform.ReverseImage(image)
	greyImg := transform.ConverToGray(image)
	sepiaImg := transform.SepiaImage(image)
	write.WriteImage(image, "results", "result")
	write.WriteImage(reversedImg, "results", "resultReversed")
	write.WriteImage(sepiaImg, "results", "resultSepia")
	write.WriteImage(greyImg, "results", "resultGrey")
}
