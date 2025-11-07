package main

import (
	"sort"

	"github.com/lucasb-eyer/go-colorful"
)

func ColorDifferenceIsNegligible(color1 [3]uint8, color2 [3]uint8, luft int) bool {
	result := true

	for i := range color1 {
		diff := int(color1[i]) - int(color2[i])
		if diff < 0 {
			result = result && (-diff) <= luft
		} else {
			result = result && diff <= luft
		}
	}
	return result
}

type Colors struct {
	SortedPixels    [][3]uint8
	Pixels_Matches  map[[3]uint8]int
	Pixels_WhiteLab map[[3]uint8]float64
}

func InitialColors(pixels_matches map[[3]uint8]int, ref_rgb *[3]uint8) Colors {
	result := Colors{SortedPixels: [][3]uint8{}, Pixels_Matches: pixels_matches, Pixels_WhiteLab: map[[3]uint8]float64{}}

	var ref colorful.Color
	if ref_rgb == nil {
		ref = colorful.Color{R: 1, G: 1, B: 1}
	} else {
		ref = colorful.Color{R: float64(ref_rgb[0]) / 255, G: float64(ref_rgb[1]) / 255, B: float64(ref_rgb[2]) / 255}
	}
	for pix := range pixels_matches {
		labc := colorful.Color{R: float64(pix[0]) / 255, G: float64(pix[1]) / 255, B: float64(pix[2]) / 255}
		result.Pixels_WhiteLab[pix] = labc.DistanceLab(ref)

		result.SortedPixels = append(result.SortedPixels, pix)
	}

	sort.Slice(result.SortedPixels, func(i, j int) bool {
		return result.Pixels_WhiteLab[result.SortedPixels[i]] < result.Pixels_WhiteLab[result.SortedPixels[j]]
	})

	return result
}

func Merge(all_pixels map[[3]uint8]int, luft int) map[[3]uint8]int {

	colors := InitialColors(all_pixels, nil)

	var choosed_colors [][3]uint8
	var choosed_counts []int
	for _, color := range colors.SortedPixels {
		if len(choosed_colors) == 0 {
			choosed_colors = append(choosed_colors, color)
			choosed_counts = append(choosed_counts, all_pixels[color])
			continue
		}

		prev_color, prev_count := &choosed_colors[len(choosed_colors)-1], &choosed_counts[len(choosed_counts)-1]

		if ColorDifferenceIsNegligible(*prev_color, color, luft) {
			for ind := range prev_color {
				prev_color[ind] = uint8((int(prev_color[ind])*(*prev_count) + int(color[ind])*all_pixels[color]) / ((*prev_count) + all_pixels[color]))
			}
			*prev_count += all_pixels[color]
		} else {
			choosed_colors = append(choosed_colors, color)
			choosed_counts = append(choosed_counts, all_pixels[color])
		}
	}

	result := make(map[[3]uint8]int)
	for i := range choosed_colors {
		result[choosed_colors[i]] = choosed_counts[i]
	}

	return result
}
