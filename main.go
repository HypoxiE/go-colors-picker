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

	// Decode the image
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	// Create an RGBA image to ensure consistent pixel data format
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// Access pixel data as a byte slice (Pix)
	// rgba.Pix is a 1D slice of bytes representing R, G, B, A values consecutively
	// For a pixel at (x, y), the R value is at index (y*rgba.Stride + x*4)
	// G is at (y*rgba.Stride + x*4 + 1), B at +2, A at +3
	pixelData := rgba.Pix

	// Example: Print the RGBA values of the first pixel
	if len(pixelData) >= 4 {
		fmt.Printf("First pixel RGBA: R=%d, G=%d, B=%d, A=%d\n",
			pixelData[0], pixelData[1], pixelData[2], pixelData[3])
	}
	var all_pixels = make(map[[3]uint8]int)

	// Example: Iterate through all pixels and print their RGBA values
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

	//fmt.Printf("%v", len(all_pixels))

	for color, match := range all_pixels {
		fmt.Printf("%v, %v\n", color, match)
	}

	//fmt.Printf("%v", all_pixels)
}
