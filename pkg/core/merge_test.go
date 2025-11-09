package gocolorspicker

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	all_pixels, _ := GetPixels("n_girl_overdose.jpg")

	fmt.Printf("%v\n", len(all_pixels))

	pixels := Merge(all_pixels[0], 0.1, 10)

	for pix, matches := range pixels {
		fmt.Printf("%v : %v\n", pix.Hex(), matches)
	}
}
