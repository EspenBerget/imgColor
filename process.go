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

	b := img.Bounds()
	var ar, ag, ab, aa, count uint32 = 0, 0, 0, 0, 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			ar += r
			ag += g
			ab += b
			aa += a
			count++
		}
	}
	ar /= count
	ag /= count
	ab /= count
	aa /= count
	avg := color.RGBA{uint8(ar), uint8(ag), uint8(ab), uint8(aa)}

	fmt.Println("Average color is", avg, count)
	res, err := os.Create("result.jpg")
	if err != nil {
		log.Fatal(err)
	}
	m := makeImage(400, 400, avg)
	if err := jpeg.Encode(res, m, nil); err != nil {
		log.Fatal(err)
	}
}
