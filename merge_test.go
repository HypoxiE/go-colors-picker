package gocolorspicker

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestMerge(t *testing.T) {
	f, err := os.Open("n_girl_overdose.jpg")
	if err != nil {
		fmt.Printf("Error opening image: %v\n", err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	pixelData := rgba.Pix
	var all_pixels = make(map[colorful.Color]int)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := y*rgba.Stride + x*4
			r := pixelData[offset]
			g := pixelData[offset+1]
			b := pixelData[offset+2]
			rgb := colorful.Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255}

			all_pixels[rgb] += 1
		}
	}

	fmt.Printf("%v\n", len(all_pixels))

	all_pixels = Merge(all_pixels, 0.1, 10)

	for pix, matches := range all_pixels {
		fmt.Printf("%v : %v\n", pix.Hex(), matches)
	}
}
