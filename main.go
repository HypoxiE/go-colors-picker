package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg" // Import for JPEG support
	_ "image/png"  // Import for PNG support
	"os"
)

const NEED_COLORS = 10
const START_BACKLASH = 10
const END_BACKLASH = 50
const STEP_BACKLASH = 2

func main() {
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
	var all_pixels = make(map[[3]uint8]int)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := y*rgba.Stride + x*4
			r := pixelData[offset]
			g := pixelData[offset+1]
			b := pixelData[offset+2]
			rgb := [3]uint8{r, g, b}

			all_pixels[rgb] += 1
		}
	}

	fmt.Printf("%v\n", len(all_pixels))

	//i := START_BACKLASH
	//for {
	//	new_pixels := Merge(all_pixels, i)
	//	if len(all_pixels) == len(new_pixels) {
	//		i += STEP_BACKLASH
	//		continue
	//	}

	//	if i == END_BACKLASH || len(new_pixels) < NEED_COLORS {
	//		break
	//	}

	//	all_pixels = new_pixels
	//}

	all_pixels = Merge(all_pixels, 4)

	fmt.Printf("%v\n", len(all_pixels))
}
