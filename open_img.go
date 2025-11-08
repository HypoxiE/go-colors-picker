package gocolorspicker

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/lucasb-eyer/go-colorful"
)

func GetPixels(name string) (map[colorful.Color]int, error) {
	f, err := os.Open(name)
	if err != nil {
		return map[colorful.Color]int{}, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return map[colorful.Color]int{}, err
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

	return all_pixels, nil
}
