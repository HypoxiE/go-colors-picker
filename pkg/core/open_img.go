package gocolorspicker

import (
	"image"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

func GetPixels(path string) ([]map[colorful.Color]int, error) {

	if strings.HasSuffix(path, ".gif") {
		rgba, bounds, err := DecodeDynamic(path)
		if err != nil {
			return []map[colorful.Color]int{}, err
		}

		var result []map[colorful.Color]int
		for i := range rgba {
			result = append(result, GetPixelsFromShot(rgba[i], bounds[i]))
		}
		return result, nil
	} else {
		rgba, bounds, err := DecodeStatic(path)
		if err != nil {
			return []map[colorful.Color]int{}, err
		}

		return []map[colorful.Color]int{GetPixelsFromShot(rgba, bounds)}, nil
	}
}

func DecodeDynamic(path string) ([]*image.RGBA, []image.Rectangle, error) {
	f, err := os.Open(path)
	if err != nil {
		return []*image.RGBA{}, []image.Rectangle{}, err
	}
	defer f.Close()

	gifImg, err := gif.DecodeAll(f)
	if err != nil {
		return []*image.RGBA{}, []image.Rectangle{}, err
	}

	var images []*image.RGBA
	var bounds []image.Rectangle
	for _, frame := range gifImg.Image {
		bound := frame.Bounds()
		rgba := image.NewRGBA(bound)
		draw.Draw(rgba, bound, frame, bound.Min, draw.Src)

		bounds = append(bounds, bound)
		images = append(images, rgba)
	}

	return images, bounds, nil
}

func DecodeStatic(path string) (*image.RGBA, image.Rectangle, error) {
	f, err := os.Open(path)
	if err != nil {
		return &image.RGBA{}, image.Rectangle{}, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return &image.RGBA{}, image.Rectangle{}, err
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	return rgba, bounds, nil
}

func GetPixelsFromShot(rgba *image.RGBA, bounds image.Rectangle) map[colorful.Color]int {
	pixelData := rgba.Pix
	var all_pixels = make(map[colorful.Color]int)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := (y-bounds.Min.Y)*rgba.Stride + (x-bounds.Min.X)*4
			r := pixelData[offset]
			g := pixelData[offset+1]
			b := pixelData[offset+2]
			rgb := colorful.Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255}

			all_pixels[rgb] += 1
		}
	}

	return all_pixels
}
