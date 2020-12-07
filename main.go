package main

import (
	"image/jpeg"
	"log"
	"os"

	"./imageprocessing"
)

func main() {
	curDir, _ := os.Getwd()
	file, err := os.Open(curDir + "/sample.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	source, err := jpeg.Decode(file)
	if err != nil {
		return
	}

	dest := imageprocessing.ColorReversal(source)

	outFile, err := os.Create("filtered_sample.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, dest, nil)
}
