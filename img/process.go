package img

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

// Hist takes a image name and finds the 8bin histogram of colors, then saves the results in the bin folder
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
	var hist [8][4]int
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			hist[r>>13][0]++
			hist[g>>13][1]++
			hist[b>>13][2]++
			hist[a>>13][3]++
		}
	}

	for i := 0; i < 8; i++ {
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
