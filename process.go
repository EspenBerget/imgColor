package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func makeImage(x, y int, c color.Color) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
	return m
}

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
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			avg.R += uint8(r)
			avg.G += uint8(g)
			avg.B += uint8(b)
			avg.A += uint8(a)
		}
	}
	fmt.Println("Average color is", avg)
	res, err := os.Create("result.jpg")
	if err != nil {
		log.Fatal(err)
	}
	m := makeImage(400, 400, avg)
	if err := jpeg.Encode(res, m, nil); err != nil {
		log.Fatal(err)
	}
}
