package gocolorspicker

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	all_pixels, _ := GetPixels("n_girl_overdose.jpg")

	fmt.Printf("%v\n", len(all_pixels))

	all_pixels = Merge(all_pixels, 0.1, 10)

	for pix, matches := range all_pixels {
		fmt.Printf("%v : %v\n", pix.Hex(), matches)
	}
}
