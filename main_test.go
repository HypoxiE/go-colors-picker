package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	var all_pixels = map[[3]uint8]int{
		{109, 91, 163}:  4,
		{107, 54, 9}:    10,
		{105, 90, 160}:  9,
		{107, 60, 9}:    100,
		{109, 101, 169}: 16,
		{113, 95, 167}:  4,
		{115, 98, 163}:  4,
	}

	a := Merge(all_pixels, 5)

	fmt.Printf("%v\n", a)
}
