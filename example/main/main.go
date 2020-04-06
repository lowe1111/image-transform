package main

import (
	"fmt"
	transform "image-go"
	"image-go/example/read"
	"image-go/example/write"
)

func main() {
	fmt.Print("Read your File...")
	image, err := read.GetImageFromPath()
	if err != nil {
		panic(err)
	}
	reversedImg := transform.ReverseImage(image)
	contrastImg := transform.GreyContrast(image, 2)
	greyImg := transform.ConverToGray(image)
	sepiaImg := transform.SepiaImage(image)
	write.WriteImage(image, "example/results", "result")
	write.WriteImage(reversedImg, "example/results", "resultReversed")
	write.WriteImage(contrastImg, "example/results", "resultContrast")
	write.WriteImage(sepiaImg, "example/results", "resultSepia")
	write.WriteImage(greyImg, "example/results", "resultGrey")
}
