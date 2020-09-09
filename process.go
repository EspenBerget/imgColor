package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // needed for its init function
	"log"
	"os"
)

// Process a image and return its colors average

func main() {
	reader, err := os.Open("sepia.jpg")
	if err != nil {
		log.Fatal("ERROR open: ", err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal("ERROR decoding: ", err)
	}
	avg := color.RGBA{0, 0, 0, 0}
	m := img.Bounds()
	for y := m.Min.Y; y < m.Max.Y; y++ {
		for x := m.Min.X; x < m.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			avg.R += uint8(r)
			avg.G += uint8(g)
			avg.B += uint8(b)
			avg.A += uint8(a)
		}
	}
	fmt.Println("Average color is", avg)
}
