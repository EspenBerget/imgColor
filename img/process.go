package img

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func makeImage(x, y int, c color.Color) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
	return m
}

// Average takes a image name and finds the average color, then saves the result to
// result.jpg
func Average(name string) {
	reader, err := os.Open("./static/" + name)
	if err != nil {
		log.Panic("ERROR open: ", err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Panic("ERROR decoding: ", err)
	}

	b := img.Bounds()
	var ar, ag, ab, aa, count uint64 = 0, 0, 0, 0, 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			ar += uint64(math.Pow(float64(r), 2))
			ag += uint64(math.Pow(float64(g), 2))
			ab += uint64(math.Pow(float64(b), 2))
			aa += uint64(math.Pow(float64(a), 2))
			count++
		}
	}
	ar /= count
	ag /= count
	ab /= count
	aa /= count
	avg := color.RGBA{uint8(ar), uint8(ag), uint8(ab), uint8(aa)}

	res, err := os.Create("result.jpg")
	if err != nil {
		log.Panic(err)
	}
	m := makeImage(100, 100, avg)
	if err := jpeg.Encode(res, m, nil); err != nil {
		log.Panic(err)
	}
}

// Hist takes a image name and finds the 8bin histogram of colors, then saves the result to
// result.jpg
func Hist(name string) {
	reader, err := os.Open("./static/" + name)
	if err != nil {
		log.Panic("ERROR open: ", err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Panic("ERROR decoding: ", err)
	}

	b := img.Bounds()
	// Code copied from Go's example in the documentation for the image package
	var hist [4][4]int
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			hist[r>>14][0]++
			hist[g>>14][1]++
			hist[b>>14][2]++
			hist[a>>14][3]++
		}
	}

	for i := 0; i < 4; i++ {
		e := hist[i]
		bin := makeImage(50, 50, color.RGBA{uint8(e[0]), uint8(e[1]), uint8(e[2]), uint8(e[3])})
		res, err := os.Create(fmt.Sprintf("bin/%d.jpg", i))
		if err != nil {
			log.Panic(err)
		}
		if err := jpeg.Encode(res, bin, nil); err != nil {
			log.Panic(err)
		}
	}
}
